// Package maxcdn is the golang bindings for MaxCDN's REST API.
//
// This package should be considered beta. The final release will be moved to
// `github.com/maxcdn/go-maxcdn`.
package maxcdn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/garyburd/go-oauth/oauth"
)

const (
	userAgent   = "Go MaxCDN API Client"
	contentType = "application/x-www-form-urlencoded"
)

// APIHost is the hostname, including protocol, to MaxCDN's API.
var APIHost = "https://rws.netdna.com"

// MaxCDN is the core struct for interacting with MaxCDN.
//
// HTTPClient can be overridden as needed, but will be set to
// http.DefaultClient by default.
type MaxCDN struct {

	// MaxCDN Consumer Alias
	Alias string

	// Display raw http Request and Response for each http Transport
	Verbose    bool
	client     oauth.Client
	HTTPClient *http.Client
}

// NewMaxCDN sets up a new MaxCDN instance.
func NewMaxCDN(alias, token, secret string) *MaxCDN {
	return &MaxCDN{
		HTTPClient: http.DefaultClient,
		Alias:      alias,
		client: oauth.Client{
			Credentials: oauth.Credentials{
				Token:  token,
				Secret: secret,
			},
			TemporaryCredentialRequestURI: APIHost + "oauth/request_token",
			TokenRequestURI:               APIHost + "oauth/access_token",
		},
	}
}

type Response struct {
	Raw  *http.Response
	Body []byte

	// Including error in Response for those times when a []Response
	// is being worked with.
	Error error
}

// NewResponse generate a new response directly from Do's return
// values.
func NewResponse(res *http.Response, err error) Response {
	defer res.Body.Close()
	r := Response{Raw: res}

	if err != nil {
		r.Error = err
		return r
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		r.Error = err
		return r
	}

	r.Body = body
	return r
}

// Get does an OAuth signed http.Get
//
// Get and other request methods return error even though it's also included in
// the Response. This is to conform to the "net/http".Get response format to make
// this more familar to consumers.
func (max *MaxCDN) Get(endpoint string, form url.Values) (*Response, error) {
	res := NewResponse(max.Do("GET", endpoint, form))
	return &res, res.Error
}

// Post does an OAuth signed http.Post
func (max *MaxCDN) Post(endpoint string, form url.Values) (*Response, error) {
	res := NewResponse(max.Do("POST", endpoint, form))
	return &res, res.Error
}

// Put does an OAuth signed http.Put
func (max *MaxCDN) Put(endpoint string, form url.Values) (*Response, error) {
	res := NewResponse(max.Do("PUT", endpoint, form))
	return &res, res.Error
}

// Delete does an OAuth signed http.Delete
func (max *MaxCDN) Delete(endpoint string, form url.Values) (*Response, error) {
	res := NewResponse(max.Do("DELETE", endpoint, form))
	return &res, res.Error
}

// PurgeZone purges a specified zones cache.
func (max *MaxCDN) PurgeZone(zone int) (*Response, error) {
	return max.Delete(fmt.Sprintf("/zones/pull.json/%d/cache", zone), nil)
}

// PurgeZones purges multiple zones caches.
func (max *MaxCDN) PurgeZones(zones []int) (resps []*Response, last error) {
	var resChannel = make(chan *Response)
	var errChannel = make(chan error)

	mutex := sync.Mutex{}
	for _, zone := range zones {
		go func(zone int) {
			res, err := max.PurgeZone(zone)
			resChannel <- res
			errChannel <- err
		}(zone)
	}

	// Wait for all responses to come back.
	// TODO: Consider adding some method of timing out.
	for _ = range zones {
		res := <-resChannel
		err := <-errChannel

		// I think the mutex might be overkill here, but I'm being
		// safe.
		mutex.Lock()
		resps = append(resps, res)
		last = err
		mutex.Unlock()
	}
	return
}

// PurgeFile purges a specified file by zone from cache.
func (max *MaxCDN) PurgeFile(zone int, file string) (*Response, error) {
	form := url.Values{}
	form.Set("file", file)

	return max.Delete(fmt.Sprintf("/zones/pull.json/%d/cache", zone), form)
}

// PurgeFiles purges multiple files from a zone.
func (max *MaxCDN) PurgeFiles(zone int, files []string) (resps []*Response, last error) {
	var resChannel = make(chan *Response)
	var errChannel = make(chan error)

	mutex := sync.Mutex{}
	for _, file := range files {
		go func(file string) {
			res, err := max.PurgeFile(zone, file)

			resChannel <- res
			errChannel <- err
		}(file)
	}

	// Wait for all responses to come back.
	// TODO: Consider adding some method of timing out.
	for _ = range files {
		res := <-resChannel
		err := <-errChannel

		// I think the mutex might be overkill here, but I'm being
		// safe.
		mutex.Lock()
		resps = append(resps, res)
		last = err
		mutex.Unlock()
	}
	return
}

// Do is a low level method to interact with MaxCDN's RESTful API. It's
// used by all other methods.
//
// If using this method, you must manually close the res.Body or bad things
// may happen.
func (max *MaxCDN) Do(method, endpoint string, form url.Values) (res *http.Response, err error) {
	var req *http.Request

	req, err = http.NewRequest(method, max.url(endpoint), nil)
	if err != nil {
		return
	}

	if method == "GET" && req.URL.RawQuery != "" {
		return nil, errors.New("oauth: url must not contain a query string")
	}

	if form != nil {
		if method == "GET" {
			req.URL.RawQuery = form.Encode()
		} else {
			req.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
		}

		// Only post needs a signed form.
		if method != "POST" {
			form = nil
		}
	}

	req.Header.Set("Authorization", max.client.AuthorizationHeader(nil, method, req.URL, form))
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", userAgent)

	if max.Verbose {
		if j, e := json.MarshalIndent(req, "", "  "); e == nil {
			fmt.Printf("Request: %s\n---\n\n", j)
		}
	}

	res, err = max.HTTPClient.Do(req)
	if max.Verbose {
		if j, e := json.MarshalIndent(res, "", "  "); e == nil {
			fmt.Printf("Response: %s\n---\n\n", j)
		}
	}
	return
}

func (max *MaxCDN) url(endpoint string) string {
	endpoint = strings.TrimPrefix(endpoint, "/")
	return fmt.Sprintf("%s/%s/%s", APIHost, max.Alias, endpoint)
}
