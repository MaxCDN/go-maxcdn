package maxcdn

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"testing"

	. "github.com/jmervine/GoT"
)

var (
	alias  = os.Getenv("ALIAS")
	token  = os.Getenv("TOKEN")
	secret = os.Getenv("SECRET")
)

func Test(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")
	Go(T).AssertEqual(max.Alias, "alias")
	Go(T).AssertEqual(max.client.Credentials.Token, "token")
	Go(T).AssertEqual(max.client.Credentials.Secret, "secret")
}

func TestMaxCDN_Get(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var record http.Response
	max.HttpClient = stubHttpOkRecorded(&record)

	payload, err := max.Get("/account.json", nil)
	Go(T).AssertNil(err)
	Go(T).RefuteNil(payload)

	Go(T).AssertEqual(record.Request.Method, "GET")
	Go(T).AssertEqual(record.Request.URL.Path, "/alias/account.json")
	Go(T).AssertEqual(record.Request.URL.Query().Encode(), "")
	Go(T).AssertEqual(record.Request.Header.Get("Content-Type"), ContentType)
	Go(T).RefuteEqual(record.Request.Header.Get("Authorization"), "")
}

func TestMaxCDN_Put(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var record http.Response
	max.HttpClient = stubHttpOkRecorded(&record)

	form := url.Values{}
	form.Add("name", "foo")

	payload, err := max.Put("/account.json", form)
	Go(T).AssertNil(err)
	Go(T).RefuteNil(payload)

	Go(T).AssertEqual(record.Request.Method, "PUT")
	Go(T).AssertEqual(record.Request.URL.Path, "/alias/account.json")
	Go(T).AssertEqual(record.Request.URL.Query().Encode(), "")
	Go(T).AssertEqual(record.Request.Header.Get("Content-Type"), ContentType)
	Go(T).RefuteEqual(record.Request.Header.Get("Authorization"), "")

	// check body
	body, err := ioutil.ReadAll(record.Request.Body)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(string(body), "name=foo")
}

func TestMaxCDN_Post(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var record http.Response
	max.HttpClient = stubHttpOkRecorded(&record)

	form := url.Values{}
	form.Add("name", "foo")

	payload, err := max.Post("/zones/pull.json", form)
	Go(T).AssertNil(err)
	Go(T).RefuteNil(payload)

	Go(T).AssertEqual(record.Request.Method, "POST")
	Go(T).AssertEqual(record.Request.URL.Path, "/alias/zones/pull.json")
	Go(T).AssertEqual(record.Request.URL.Query().Encode(), "")
	Go(T).AssertEqual(record.Request.Header.Get("Content-Type"), ContentType)
	Go(T).RefuteEqual(record.Request.Header.Get("Authorization"), "")

	// check body
	body, err := ioutil.ReadAll(record.Request.Body)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(string(body), "name=foo")
}

func TestMaxCDN_Delete(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var record http.Response
	max.HttpClient = stubHttpOkRecorded(&record)

	payload, err := max.Delete("/zones/pull.json/123456")
	Go(T).AssertNil(err)
	Go(T).RefuteNil(payload)

	Go(T).AssertEqual(record.Request.Method, "DELETE")
	Go(T).AssertEqual(record.Request.URL.Path, "/alias/zones/pull.json/123456")
	Go(T).AssertEqual(record.Request.URL.Query().Encode(), "")
	Go(T).AssertEqual(record.Request.Header.Get("Content-Type"), ContentType)
	Go(T).RefuteEqual(record.Request.Header.Get("Authorization"), "")
}

func TestMaxCDN_PurgeZone(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var record http.Response
	max.HttpClient = stubHttpOkRecorded(&record)

	payload, err := max.PurgeZone(123456)
	Go(T).AssertNil(err)
	Go(T).RefuteNil(payload)

	Go(T).AssertEqual(record.Request.Method, "DELETE")
	Go(T).AssertEqual(record.Request.URL.Path, "/alias/zones/pull.json/123456/cache")
	Go(T).AssertEqual(record.Request.URL.Query().Encode(), "")
	Go(T).AssertEqual(record.Request.Header.Get("Content-Type"), ContentType)
	Go(T).RefuteEqual(record.Request.Header.Get("Authorization"), "")
}

// Overly elaborte examples go...
//
// Why? These are used as integration tests as well.
//
// Run like:
//
// $ ALIAS=your_alias TOKEN=your_token SECRET=your_secret go test

func Example() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	payload, err := max.Get("/account.json", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", payload.Data)
}

func Example_Functional() {
	// This, like all examples are meant to be functional integration tests.
	// To run these as integration tests export your ALIAS, TOKEN and SECRET
	//to your envioronment before running 'go test', otherwise the http
	// request will be stubbed using the json files in './_fixtures/*.json'.

	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HttpClient = stubHttpOk()
	}

	payload, err := max.Get("/account.json", nil)
	if err != nil {
		panic(err)
	}

	data := payload.Data["account"].(map[string]interface{})
	if data["name"] != "" {
		fmt.Println("GET /account.json succeeded")
	}

	// Output: GET /account.json succeeded
}

func ExampleNewMaxCDN() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	fmt.Printf("%#v\n", max)
}

func ExampleMaxCDN_Get() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	payload, err := max.Get("/account.json", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", payload.Data)
}

func ExampleMaxCDN_Put() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	form := url.Values{}
	form.Set("name", "example_name")
	payload, err := max.Put("/account.json", form)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", payload.Data)
}

func Example_Functional_MaxCDN_Put() {
	// This, like all examples are meant to be functional integration tests.
	// To run these as integration tests export your ALIAS, TOKEN and SECRET
	//to your envioronment before running 'go test', otherwise the http
	// request will be stubbed using the json files in './_fixtures/*.json'.

	// I'm using a timestamp as unique name for testing, you shouldn't
	// do this.
	name := stringFromTimestamp()

	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HttpClient = stubHttpOk()
		name = "MaxCDN sampleCode"
	}

	form := url.Values{}
	form.Set("name", name)
	payload, err := max.Put("/account.json", form)

	if err != nil {
		panic(err)
	}

	data := payload.Data["account"].(map[string]interface{})
	if data["name"] == name {
		fmt.Println("PUT /account.json succeeded")
	}

	// Output: PUT /account.json succeeded
}

func ExampleMaxCDN_Delete() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	// This specific example shows how to purge a cache without using the Purge
	// methods, more as an example of using Delete, then anything, really.

	payload, err := max.Delete(fmt.Sprintf("/zones/pull.json/%d/cache", 123456))
	if err != nil {
		panic(err)
	}

	if payload.Code == 200 {
		fmt.Println("Purge succeeded")
	}
}

func ExampleMaxCDN_Post() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	form := url.Values{}
	form.Set("name", "newzone")

	// When creating a new zone, the url must be real and resolve.
	form.Set("url", "http://www.example.com")

	payload, err := max.Post("/zones/pull.json", form)
	if err != nil {
		panic(err)
	}

	data := payload.Data["pullzone"].(map[string]interface{})
	if data["name"] == "newzone" {
		fmt.Println("Successfully created new Pull Zone.")
	}

}

func Example_Functional_MaxCDN_Post() {
	// This, like all examples are meant to be functional integration tests.
	// To run these as integration tests export your ALIAS, TOKEN and SECRET
	//to your envioronment before running 'go test', otherwise the http
	// request will be stubbed using the json files in './_fixtures/*.json'.

	// I'm using a timestamp as unique name for testing, you shouldn't
	// do this.
	name := stringFromTimestamp()

	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HttpClient = stubHttpOk()
		name = "newpullzone3"
	}

	form := url.Values{}
	form.Set("name", name)

	// When creating a new zone, the url must be real and resolve.
	form.Set("url", "http://www.example.com")

	payload, err := max.Post("/zones/pull.json", form)
	if err != nil {
		panic(err)
	}

	data := payload.Data["pullzone"].(map[string]interface{})
	if data["name"] == name {
		fmt.Println("POST /zones/pull.json succeeded")

		id := int(data["id"].(float64))

		payload, err = max.Delete(fmt.Sprintf("/zones/pull.json/%d", id))
		if err != nil {
			panic(err)
		}

		if payload.Code == 200 {
			fmt.Println("DELETE /zones/pull.json succeeded")
		}
	}

	// Output:
	// POST /zones/pull.json succeeded
	// DELETE /zones/pull.json succeeded
}

func ExampleMaxCDN_PurgeZone() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	payload, err := max.PurgeZone(123456)
	if err != nil {
		panic(err)
	}

	if payload.Code == 200 {
		fmt.Println("Purge succeeded")
	}
}

func Example_Functional_MaxCDN_PurgeZone() {
	// This, like all examples are meant to be functional integration tests.
	// To run these as integration tests export your ALIAS, TOKEN and SECRET
	//to your envioronment before running 'go test', otherwise the http
	// request will be stubbed using the json files in './_fixtures/*.json'.

	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HttpClient = stubHttpOk()
	}

	// Start by fetching the first zone, as that's the one we'll be purging.
	payload, err := max.Get("/zones/pull.json", nil)
	if err != nil {
		panic(err)
	}

	data := payload.Data["pullzones"].([]interface{})
	zone := data[0].(map[string]interface{})

	// complexity due to bug in response where string is returned instead of an int
	// for "id", see:
	//
	// https://github.com/MaxCDN/api-docs/issues/20
	id, e := strconv.ParseInt(zone["id"].(string), 0, 64)
	if e != nil {
		panic(e)
	}

	// Now purge that zone's cache.
	payload, err = max.PurgeZone(int(id))
	if err != nil {
		panic(err)
	}

	if payload.Code == 200 {
		fmt.Println("Purge succeeded")
	}

	// Output: Purge succeeded
}

func ExampleMaxCDN_PurgeZones() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	zones := []int{123456, 234567, 345678}
	payloads, err := max.PurgeZones(zones)
	if err != nil {
		panic(err)
	}

	if len(payloads) == len(zones) {
		fmt.Printf("Purges succeeded")
	}
}

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
	} else if endpoint == "pull.json" && r.Method == "GET" {
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

func stubHttpOkRecorded(record *http.Response) *http.Client {
	return &http.Client{
		Transport: &stubRoundTripper{
			ResponseRecord: record,
		},
	}
}

func stubHttpOk() *http.Client {
	return stubHttpOkRecorded(new(http.Response))
}

func stubHttpErr() *http.Client {
	return &http.Client{Transport: &stubRoundTripper{Error: true}}
}
