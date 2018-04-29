package maxcdn

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	assert.Equal(t, "alias", max.Alias)
}

func TestMaxCDN_Get(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data Generic
	rsp, err := max.Get(&data, "/account.json", nil)

	// check error
	assert.Nil(err)

	// check response
	assert.NotNil(rsp)
	assert.NotNil(rsp.Data)

	// check account
	name := data["account"].(map[string]interface{})["name"].(string)
	assert.Equal("MaxCDN sampleCode", name)

	// check record of http request from stub
	assert.Equal("GET", recorder.Request.Method)
	assert.Equal("/alias/account.json", recorder.Request.URL.Path)
	assert.Equal("", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}

func TestMaxCDN_GetLogs(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	rsp, err := max.GetLogs(nil)

	// check error
	assert.Nil(err)

	// check response
	assert.NotNil(rsp)
	assert.NotNil(rsp.Page)

	// check account
	assert.Equal(1, rsp.Page)
	assert.Equal("1404229642374", rsp.NextPageKey)

	// check record of http request from stub
	assert.Equal("GET", recorder.Request.Method)
	assert.Equal("/alias/v3/reporting/logs.json", recorder.Request.URL.Path)
	assert.Equal("", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}

func TestMaxCDN_Put(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	form := url.Values{}
	form.Add("name", "MaxCDN sampleCode")

	var data Generic
	rsp, err := max.Put(&data, "/account.json", form)

	// check error
	assert.Nil(err)

	// check response
	assert.NotNil(rsp)
	assert.NotNil(rsp.Data)

	// check account
	name := data["account"].(map[string]interface{})["name"].(string)
	assert.Equal("MaxCDN sampleCode", name)

	// check record of http request from stub
	assert.Equal("PUT", recorder.Request.Method)
	assert.Equal("/alias/account.json", recorder.Request.URL.Path)
	assert.Equal("", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	body, err := ioutil.ReadAll(recorder.Request.Body)
	assert.Nil(err)
	assert.Equal("name=MaxCDN+sampleCode", string(body))
}

func TestMaxCDN_Post(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	form := url.Values{}
	form.Add("name", "foo")

	var data Generic
	rsp, err := max.Post(&data, "/zones/pull.json", form)

	// check error
	assert.Nil(err)

	// check response
	assert.NotNil(rsp)
	assert.NotNil(rsp.Data)

	assert.Equal("POST", recorder.Request.Method)
	assert.Equal("/alias/zones/pull.json", recorder.Request.URL.Path)
	assert.Equal("", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	body, err := ioutil.ReadAll(recorder.Request.Body)
	assert.Nil(err)
	assert.Equal("name=foo", string(body))
}

func TestMaxCDN_Delete(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	rsp, err := max.Delete("/zones/pull.json/123456/cache", nil)
	assert.Nil(err)
	assert.NotNil(rsp)
	assert.NotNil(rsp.Code)

	assert.Equal("DELETE", recorder.Request.Method)
	assert.Equal("/alias/zones/pull.json/123456/cache", recorder.Request.URL.Path)
	assert.Equal("", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}

func TestMaxCDN_PurgeZone(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	rsp, err := max.PurgeZone(123456)
	assert.Nil(err)
	assert.NotNil(rsp)
	assert.NotNil(rsp.Code)

	assert.Equal("DELETE", recorder.Request.Method)
	assert.Equal("/alias/zones/pull.json/123456/cache", recorder.Request.URL.Path)
	assert.Equal("", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}

func TestMaxCDN_PurgeZones(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	rsps, err := max.PurgeZones([]int{12345, 23456, 34567})
	assert.Nil(err)
	assert.NotNil(rsps)

	assert.Equal("DELETE", recorder.Request.Method)
	assert.Equal("", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}

func TestMaxCDN_PurgeZonesString(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	rsps, err := max.PurgeZonesString([]string{"12345", "23456", "34567"})
	assert.Nil(err)
	assert.NotNil(rsps)

	assert.Equal("DELETE", recorder.Request.Method)
	assert.Equal("", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}

func TestMaxCDN_PurgeFile(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	rsp, err := max.PurgeFile(123456, "/master.css")
	assert.Nil(err)
	assert.NotNil(rsp)
	assert.NotNil(rsp.Code)

	assert.Equal("DELETE", recorder.Request.Method)
	assert.Equal("/alias/zones/pull.json/123456/cache", recorder.Request.URL.Path)
	assert.Equal("files=%2Fmaster.css", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}

func TestMaxCDN_PurgeFileString(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	rsp, err := max.PurgeFileString("123456", "/master.css")
	assert.Nil(err)
	assert.NotNil(rsp)
	assert.NotNil(rsp.Code)

	assert.Equal("DELETE", recorder.Request.Method)
	assert.Equal("/alias/zones/pull.json/123456/cache", recorder.Request.URL.Path)
	assert.Equal("files=%2Fmaster.css", recorder.Request.URL.Query().Encode())
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}

func TestMaxCDN_PurgeFiles(t *testing.T) {
	assert := assert.New(t)
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	files := []string{"/master.css", "/master.js", "/index.html"}
	rsp, err := max.PurgeFiles(123456, files)
	assert.Nil(err)
	assert.NotNil(rsp)

	assert.Equal("DELETE", recorder.Request.Method)
	assert.Contains(recorder.Request.URL.Query().Encode(), "files=")
	assert.Equal(contentType, recorder.Request.Header.Get("Content-Type"))
	assert.NotEqual("", recorder.Request.Header.Get("Authorization"))

	// check body
	assert.Nil(recorder.Request.Body)
}
