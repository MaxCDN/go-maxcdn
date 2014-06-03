// Package maxcdn is the golang bindings for MaxCDN's REST API.
//
// At this time it should be considered very alpha.
package maxcdn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
)

const (
	ApiPath     = "https://rws.netdna.com"
	UserAgent   = "Go MaxCDN API Client"
	ContentType = "application/x-www-form-urlencoded"
)

// Response is the core data type for JSON responses from API calls.
type Response struct {
	Code  float64                `json:"code"`
	Data  map[string]interface{} `json:"data"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

// MaxCDN is the core struct for interacting with MaxCDN.
//
// HttpClient can be overridden as needed, but will be set to
// http.DefaultClient by default.
type MaxCDN struct {
	Alias      string
	client     oauth.Client
	HttpClient *http.Client
}

// NewMaxCDN sets up a new MaxCDN instance.
func NewMaxCDN(alias, token, secret string) *MaxCDN {
	return &MaxCDN{
		HttpClient: http.DefaultClient,
		Alias:      alias,
		client: oauth.Client{
			Credentials: oauth.Credentials{
				Token:  token,
				Secret: secret,
			},
			TemporaryCredentialRequestURI: ApiPath + "oauth/request_token",
			TokenRequestURI:               ApiPath + "oauth/access_token",
		},
	}
}

func (max *MaxCDN) Get(endpoint string, form url.Values) (*Response, error) {
	return max.do("GET", endpoint, form)
}

func (max *MaxCDN) Post(endpoint string, form url.Values) (*Response, error) {
	return max.do("POST", endpoint, form)
}

func (max *MaxCDN) Put(endpoint string, form url.Values) (*Response, error) {
	return max.do("PUT", endpoint, form)
}

func (max *MaxCDN) Delete(endpoint string) (*Response, error) {
	return max.do("DELETE", endpoint, nil)
}

func (max *MaxCDN) url(endpoint string) string {
	endpoint = strings.TrimPrefix(endpoint, "/")
	return fmt.Sprintf("%s/%s/%s", ApiPath, max.Alias, endpoint)
}

func (max *MaxCDN) do(method, endpoint string, form url.Values) (response *Response, err error) {
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
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)

	resp, err := max.HttpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	return max.parse(resp)
}

func (max *MaxCDN) parse(resp *http.Response) (*Response, error) {
	// process the response

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload Response
	err = json.Unmarshal(data, &payload)

	return &payload, err
}
