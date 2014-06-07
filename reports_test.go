package maxcdn

import (
	//"fmt"
	//"io/ioutil"
	"net/http"
	//"net/url"
	//"os"

	"testing"

	. "github.com/jmervine/GoT"
)

func TestMaxCDN_GetPopularFiles(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	payload, err := max.GetPopularFiles(nil)
	Go(T).AssertNil(err)
	Go(T).RefuteNil(payload)

	Go(T).AssertEqual(recorder.Request.Method, "GET")
	Go(T).AssertEqual(recorder.Request.URL.Path, "/alias/reports/popularfiles.json")
	Go(T).AssertEqual(recorder.Request.URL.Query().Encode(), "")
	Go(T).AssertEqual(recorder.Request.Header.Get("Content-Type"), contentType)
	Go(T).RefuteEqual(recorder.Request.Header.Get("Authorization"), "")

	// check body
	Go(T).AssertNil(recorder.Request.Body)
}
