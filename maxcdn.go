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
	logsPath    = "/v3/reporting/logs.json"
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
func (max *MaxCDN) Get(endpointType interface{}, endpoint string, form url.Values) (*Response, error) {
	return max.DoParse(endpointType, "GET", endpoint, form)
}

// GetLogs is a seperate getter for MaxCDN's logs.json endpoint, as it currently doesn't follow
// the json format of other endpoints.
func (max *MaxCDN) GetLogs(form url.Values) (Logs, error) {
	var logs Logs
	rsp, err := max.Request("GET", logsPath, form)
	defer rsp.Body.Close()
	if err != nil {
		return logs, err
	}

	var raw []byte
	raw, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return logs, err
	}

	err = json.Unmarshal(raw, &logs)
	return logs, err
}

// Post does an OAuth signed http.Post
func (max *MaxCDN) Post(endpointType interface{}, endpoint string, form url.Values) (*Response, error) {
	return max.DoParse(endpointType, "POST", endpoint, form)
}

// Put does an OAuth signed http.Put
func (max *MaxCDN) Put(endpointType interface{}, endpoint string, form url.Values) (*Response, error) {
	return max.DoParse(endpointType, "PUT", endpoint, form)
}

// Delete does an OAuth signed http.Delete
//
// Delete does not take an endpointType because delete only returns a status code.
func (max *MaxCDN) Delete(endpoint string, form url.Values) (*Response, error) {
	return max.Do("DELETE", endpoint, form)
}

// PurgeZone purges a specified zones cache.
func (max *MaxCDN) PurgeZone(zone int) (*Response, error) {
	return max.Delete(fmt.Sprintf("/zones/pull.json/%d/cache", zone), nil)
}

// PurgeZoneString purges a specified zones cache.
func (max *MaxCDN) PurgeZoneString(zone string) (*Response, error) {
	return max.Delete(fmt.Sprintf("/zones/pull.json/%s/cache", zone), nil)
}

// PurgeZonesString purges multiple zones caches.
func (max *MaxCDN) PurgeZonesString(zones []string) (resps []*Response, last error) {
	var resChannel = make(chan *Response)
	var errChannel = make(chan error)

	mutex := sync.Mutex{}
	for _, zone := range zones {
		go func(zone string) {
			res, err := max.PurgeZoneString(zone)
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

// PurgeZones purges multiple zones caches.
func (max *MaxCDN) PurgeZones(zones []int) (resps []*Response, last error) {
	zoneStrings := make([]string, len(zones))

	for i, zone := range zones {
		zoneStrings[i] = fmt.Sprintf("%d", zone)
	}

	return max.PurgeZonesString(zoneStrings)
}

// PurgeFile purges a specified file by zone from cache.
func (max *MaxCDN) PurgeFile(zone int, file string) (*Response, error) {
	return max.PurgeFileString(fmt.Sprintf("%d", zone), file)
}

// PurgeFile purges a specified file by zone from cache.
func (max *MaxCDN) PurgeFileString(zone string, file string) (*Response, error) {
	form := url.Values{}
	form.Set("file", file)

	return max.Delete(fmt.Sprintf("/zones/pull.json/%s/cache", zone), form)
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

func (max *MaxCDN) DoParse(endpointType interface{}, method, endpoint string, form url.Values) (rsp *Response, err error) {
	rsp, err = max.Do(method, endpoint, form)
	if err != nil {
		return
	}
	err = json.Unmarshal(rsp.Data, &endpointType)
	return
}

// Do is a low level method to interact with MaxCDN's RESTful API via Request
// and return a parsed Response. It's used by all other methods.
//
// This method closes the raw http.Response body.
func (max *MaxCDN) Do(method, endpoint string, form url.Values) (rsp *Response, err error) {
	rsp = new(Response)
	res, err := max.Request(method, endpoint, form)
	defer res.Body.Close()

	if err != nil {
		return
	}

	headers := res.Header
	rsp.Headers = &headers

	var raw []byte
	raw, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(raw, &rsp)
	if err != nil {
		return
	}

	if rsp.Code > 299 {
		return rsp, fmt.Errorf("%s: %s", rsp.Error.Type, rsp.Error.Message)
	}

	return
}

// Request is a low level method to interact with MaxCDN's RESTful API. It's
// used by all other methods.
//
// If using this method, you must manually close the res.Body or bad things
// may happen.
func (max *MaxCDN) Request(method, endpoint string, form url.Values) (res *http.Response, err error) {
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
