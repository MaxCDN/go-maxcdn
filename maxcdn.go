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

// Get does an OAuth signed http.Get
func (max *MaxCDN) Get(endpoint string, form url.Values) (mapper *GenericResponse, err error) {
	mapper = new(GenericResponse)
	raw, res, err := max.Do("GET", endpoint, form)
	mapper.Response = res
	if err != nil {
		return
	}

	err = mapper.Parse(raw)
	return
}

// Post does an OAuth signed http.Post
func (max *MaxCDN) Post(endpoint string, form url.Values) (mapper *GenericResponse, err error) {
	mapper = new(GenericResponse)
	raw, res, err := max.Do("POST", endpoint, form)
	mapper.Response = res
	if err != nil {
		return
	}

	err = mapper.Parse(raw)
	return
}

// Put does an OAuth signed http.Put
func (max *MaxCDN) Put(endpoint string, form url.Values) (mapper *GenericResponse, err error) {
	mapper = new(GenericResponse)
	raw, res, err := max.Do("PUT", endpoint, form)
	mapper.Response = res
	if err != nil {
		return
	}

	err = mapper.Parse(raw)
	return
}

// Delete does an OAuth signed http.Delete
func (max *MaxCDN) Delete(endpoint string) (mapper *GenericResponse, err error) {
	mapper = new(GenericResponse)
	raw, res, err := max.Do("DELETE", endpoint, nil)
	mapper.Response = res
	if err != nil {
		return
	}

	err = mapper.Parse(raw)
	return
}

// PurgeZone purges a specified zones cache.
func (max *MaxCDN) PurgeZone(zone int) (*GenericResponse, error) {
	return max.Delete(fmt.Sprintf("/zones/pull.json/%d/cache", zone))
}

// PurgeZones purges multiple zones caches.
func (max *MaxCDN) PurgeZones(zones []int) (resps []*GenericResponse, last error) {
	var resChannel = make(chan *GenericResponse)
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
func (max *MaxCDN) PurgeFile(zone int, file string) (mapper *GenericResponse, err error) {
	form := url.Values{}
	form.Set("file", file)

	mapper = new(GenericResponse)
	raw, res, err := max.Do("DELETE", fmt.Sprintf("/zones/pull.json/%d/cache", zone), form)
	mapper.Response = res
	if err != nil {
		return
	}

	err = mapper.Parse(raw)
	return
}

// PurgeFiles purges multiple files from a zone.
func (max *MaxCDN) PurgeFiles(zone int, files []string) (resps []*GenericResponse, last error) {
	var resChannel = make(chan *GenericResponse)
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

func (max *MaxCDN) url(endpoint string) string {
	endpoint = strings.TrimPrefix(endpoint, "/")
	return fmt.Sprintf("%s/%s/%s", APIHost, max.Alias, endpoint)
}

// Do is a generic method to interact with MaxCDN's RESTful API. It's
// used by all other methods.
//
// It's purpose though would be for you to generate your own struct more
// exactly mapping the json response to your purpose. More specific
// responses are planned for future versions, but there are too many make
// it worth implementing all of them, so this support should remain.
func (max *MaxCDN) Do(method, endpoint string, form url.Values) (raw []byte, res *http.Response, err error) {
	var req *http.Request

	req, err = http.NewRequest(method, max.url(endpoint), nil)
	if err != nil {
		return
	}

	if method == "GET" && req.URL.RawQuery != "" {
		return nil, nil, errors.New("oauth: url must not contain a query string")
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
			fmt.Printf("---\nRequest:\n%+v\n\n", j)
		}
	}

	res, err = max.HTTPClient.Do(req)
	defer res.Body.Close()

	if max.Verbose {
		if j, e := json.MarshalIndent(res, "", "  "); e == nil {
			fmt.Printf("---\nResponse:\n%+v\n\n", j)
		}
	}

	raw, err = ioutil.ReadAll(res.Body)

	// Note: returning the response along with the raw body and err seems a bit clunky,
	// but there are valid use-cases having the raw response is useful. For an example,
	// see tools/maxcurl/maxcurl.go and it's header flag implementation.
	return raw, res, err
}
