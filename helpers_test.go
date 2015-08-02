// This file contains helper methods for testing.

package maxcdn_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// utils

// Generate a unique string for testing from current timestamp.
func stringFromTimestamp() (name string) {
	t := time.Now()
	return fmt.Sprintf("go-%04d%02d%02d%02d%02d%02d%03d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/int(time.Millisecond))
}

// stubRoundTripper is an http client intercept that grabs
// the request and returns json depending on it's path.
//
// if Error is true, it will return an error response from
// _fixures/error.json

type stubRoundTripper struct {
	ResponseRecord *http.Response
	Error          bool
}

func (crt *stubRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	var (
		urlParts = strings.Split(r.URL.Path, "/")
		endpoint = urlParts[len(urlParts)-1]
		code     = http.StatusOK
		filename string
	)

	switch {
	case crt.Error:
		filename = "error.json"
		code = 500
	case r.Method == "DELETE":
		filename = "delete.json"
	case endpoint == "pull.json" && r.Method == "PUT":
		filename = "pullzone.json"
	case endpoint == "pull.json" && r.Method == "POST":
		filename = "post.pull.json"
	case endpoint == "pull.json":
		filename = "pullzones.json"
	case endpoint == "address":
		filename = "address.json"
	case endpoint == "daily":
		filename = "stats.daily.json"
	case strings.Contains(r.URL.Path, "pull.json"):
		filename = "pullzone.json"
	case endpoint == "users.json", strings.Contains(r.URL.Path, "users.json"):
		filename = "users.json"
	default:
		filename = endpoint
	}

	read := fetchJSON(filename)

	crt.ResponseRecord.Body = ioutil.NopCloser(bytes.NewBuffer(read))
	crt.ResponseRecord.StatusCode = code
	crt.ResponseRecord.Request = r

	return crt.ResponseRecord, nil
}

func fetchJSON(p string) []byte {
	read, err := ioutil.ReadFile("_fixtures/" + p)
	if err != nil {
		panic(err)
	}
	return read
}

func stubHTTPOkRecorded(recorder *http.Response) *http.Client {
	return &http.Client{
		Transport: &stubRoundTripper{
			ResponseRecord: recorder,
		},
	}
}

func stubHTTPOk() *http.Client {
	return stubHTTPOkRecorded(new(http.Response))
}

func stubHTTPErrRecorded(recorder *http.Response) *http.Client {
	return &http.Client{
		Transport: &stubRoundTripper{
			ResponseRecord: recorder,
			Error:          true,
		},
	}
}
