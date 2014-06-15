// This file contains helper methods for testing.

package maxcdn

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
	//fmt.Println("stub")
	urlParts := strings.Split(r.URL.Path, "/")
	endpoint := urlParts[len(urlParts)-1]
	code := 200

	var filename string

	if crt.Error {
		filename = "error.json"
		code = 500
	} else if r.Method == "DELETE" {
		filename = "delete.json"
	} else if endpoint == "pull.json" && r.Method != "GET" {
		filename = "pullzone.json"
	} else if endpoint == "pull.json" {
		filename = "pullzones.json"
	} else {
		filename = endpoint
	}

	read, err := ioutil.ReadFile("_fixtures/" + filename)
	if err != nil {
		panic(err)
	}

	crt.ResponseRecord.Body = ioutil.NopCloser(bytes.NewBuffer(read))
	crt.ResponseRecord.StatusCode = code
	crt.ResponseRecord.Request = r

	return crt.ResponseRecord, nil
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
