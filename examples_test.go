package maxcdn_test

// This file contains Example and Functional Example methods, both for documentation
// and for functionally testing code.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/MaxCDN/go-maxcdn"
)

var (
	alias  = os.Getenv("ALIAS")
	token  = os.Getenv("TOKEN")
	secret = os.Getenv("SECRET")
)

func Example() {
	// Basic Get
	max := maxcdn.NewMaxCDN(alias, token, secret)
	var got maxcdn.Generic
	res, err := max.Get(&got, "/account.json", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", res.Code)
	fmt.Printf("name: %s\n", got["name"].(string))

	// Basic Put
	form := url.Values{}
	form.Set("name", "new name")

	var put maxcdn.Generic
	if _, err = max.Put(&put, "/account.json", form); err == nil &&
		put["name"].(string) == "new name" {
		fmt.Println("name successfully updated")
	}

	// Basic Delete
	if _, err = max.Delete("/zones/pull.json/123456", nil); err == nil {
		fmt.Println("zone successfully deleted")
	}

	// Logs
	if logs, err := max.GetLogs(nil); err == nil {
		for _, line := range logs.Records {
			fmt.Printf("%+v\n", line)
		}
	}
}

func Example_newMaxCDN() {
	max := maxcdn.NewMaxCDN(alias, token, secret)
	fmt.Printf("%#v\n", max)
}

func ExampleMaxCDN_doParse() {
	// Run mid-level DoParse method.
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	response, err := max.DoParse(&data, "GET", "/account.json/address", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data["street1"].(string))
}

func ExampleMaxCDN_do() {
	// Run low-level Do method.
	max := maxcdn.NewMaxCDN(alias, token, secret)

	if rsp, err := max.Do("GET", "/account.json", nil); err == nil {
		fmt.Printf("Response Code: %d\n", rsp.Code)

		var data maxcdn.Generic
		if err = json.Unmarshal(rsp.Data, &data); err == nil {
			fmt.Printf("%+v\n", data["account"])
		}
	}
}

func ExampleMaxCDN_request() {
	// Run low-level Request method.
	max := maxcdn.NewMaxCDN(alias, token, secret)

	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	res, err := max.Request("GET", "/account.json", nil)
	defer res.Body.Close()
	check(err)

	body, err := ioutil.ReadAll(res.Body)
	check(err)

	fmt.Println(string(body))
}

func ExampleMaxCDN_get() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	response, err := max.Get(&data, "/account.json/address", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data["street1"].(string))
}

func ExampleMaxCDN_getLogs() {
	max := maxcdn.NewMaxCDN(alias, token, secret)
	if logs, err := max.GetLogs(nil); err == nil {
		for _, line := range logs.Records {
			fmt.Printf("%+v\n", line)
		}
	}
}

func ExampleMaxCDN_put() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	form := url.Values{}
	form.Set("name", "example name")

	var data maxcdn.Generic
	response, err := max.Put(&data, "/account.json", form)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data["name"].(string))
}

func ExampleMaxCDN_delete() {
	// This specific example shows how to purge a cache without using the Purge
	// methods, more as an example of using Delete, then anything, really.
	max := maxcdn.NewMaxCDN(alias, token, secret)

	res, err := max.Delete("/zones/pull.json/123456/cache", nil)
	if err != nil {
		panic(err)
	}

	if res.Code == 200 {
		fmt.Println("Purge suucceeded")
	}
}

func ExampleMaxCDN_purgeZone() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	rsp, err := max.PurgeZone(123456)
	if err != nil {
		panic(err)
	}

	if rsp.Code == 200 {
		fmt.Println("Purge succeeded")
	}
}

func ExampleMaxCDN_purgeZones() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	zones := []int{123456, 234567, 345678}
	rsps, err := max.PurgeZones(zones)
	if err != nil {
		panic(err)
	}

	if len(rsps) == len(zones) {
		fmt.Printf("Purges succeeded")
	}
}

func ExampleMaxCDN_purgeFile() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	payload, err := max.PurgeFile(123456, "/master.css")
	if err != nil {
		panic(err)
	}

	if payload.Code == 200 {
		fmt.Println("Purge succeeded")
	}
}

func ExampleMaxCDN_purgeFiles() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	files := []string{"/master.css", "/master.js"}
	payloads, err := max.PurgeFiles(123456, files)
	if err != nil {
		panic(err)
	}

	if len(payloads) == len(files) {
		fmt.Printf("Purges succeeded")
	}
}

func ExampleMaxCDN_post() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	form := url.Values{}
	form.Set("name", "newzone")

	// When creating a new zone, the url must be real and resolve.
	form.Set("url", "http://www.example.com")

	var data maxcdn.Generic
	_, err := max.Post(&data, "/zones/pull.json", form)
	if err != nil {
		panic(err)
	}

	if data["name"].(string) == "newzone" {
		fmt.Println("Successfully created new Pull Zone.")
	}
}

func Example_response() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	response, _ := max.Get(&data, "/account.json", nil)
	fmt.Printf("%+v\n", response)
}

func Example_generic() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	if _, err := max.Get(&data, "/account.json", nil); err == nil {
		alias := data["alias"].(string)
		name := data["name"].(string)
		fmt.Printf("alias: %s\n", alias)
		fmt.Printf("name:  %s\n", name)
	}
}

func Example_account() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	if _, err := max.Get(&data, "/account.json", nil); err == nil {
		fmt.Printf("%+v\n", data)
	}
}

func Example_accountAddress() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	if _, err := max.Get(&data, "/account.json/address", nil); err == nil {
		fmt.Printf("%+v\n", data)
	}
}

func Example_popularFiles() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	if _, err := max.Get(&data, "/reports/popularfiles.json", nil); err == nil {
		for i, file := range data {
			uri := file.(map[string]interface{})["uri"].(string)
			hit := file.(map[string]interface{})["hit"].(string)
			fmt.Printf("%2s: %30s=%s, \n", i, uri, hit)
		}
	}
	fmt.Println("----")
	summaryHit := data["summary"].(map[string]interface{})["uri"].(string)
	fmt.Printf("    %30s=%s, \n", "summary", summaryHit)
}

func Example_statsSummary() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	if _, err := max.Get(&data, "/reports/stats.json", nil); err == nil {
		fmt.Printf("%+v\n", data)
	}
}

func Example_stats() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	var data maxcdn.Generic
	if _, err := max.Get(&data, "/reports/stats.json/hourly", nil); err == nil {
		fmt.Printf("%+v\n", data)
	}
}

// These "Functional" examples are meant to be functional integration tests.
// To run these as integration tests export your ALIAS, TOKEN and SECRET
// to your envioronment before running 'go test', otherwise the http
// request will be stubbed using the json files in './_fixtures/*.json'.

// Run like:

// $ ALIAS=your_alias TOKEN=your_token SECRET=your_secret make test

func Example_functional() {
	var form url.Values

	name := stringFromTimestamp()
	max := maxcdn.NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HTTPClient = stubHTTPOk()
	}

	// GET /account.json
	var getAcct maxcdn.Generic
	res, err := max.Get(&getAcct, "/account.json", nil)
	if err != nil || res.Code != 200 ||
		getAcct["account"].(map[string]interface{})["name"].(string) == "" {
		fmt.Println("GET account")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", getAcct)
	}

	// PUT /account.json
	var putAcct maxcdn.Generic
	form = url.Values{}
	form.Set("name", name)
	res, err = max.Put(&putAcct, "/account.json", form)
	if err != nil || res.Code != 200 ||
		putAcct["account"].(map[string]interface{})["name"].(string) == "" {
		fmt.Println("PUT account")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", putAcct)
	}

	// GET /account.json/address
	var getAddr maxcdn.Generic
	res, err = max.Get(&getAddr, "/account.json/address", nil)
	if err != nil || res.Code != 200 ||
		getAddr["address"].(map[string]interface{})["street1"].(string) == "" {
		fmt.Println("GET address")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", getAddr)
	}

	// PUT /account.json/address
	var putAddr maxcdn.Generic
	form = url.Values{}
	form.Set("street1", name)
	res, err = max.Put(&putAddr, "/account.json/address", form)
	if err != nil || res.Code != 200 ||
		putAddr["address"].(map[string]interface{})["street1"].(string) == "" {
		fmt.Println("GET address")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", putAddr)
	}

	// GET /reports/popularfiles.json
	var popular maxcdn.Generic
	res, err = max.Get(&popular, "/reports/popularfiles.json", nil)
	if err != nil || res.Code != 200 ||
		popular["popularfiles"].([]interface{})[0].(map[string]interface{})["uri"].(string) == "" {
		fmt.Println("GET popularfiles")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", popular)
	}

	file := popular["popularfiles"].([]interface{})[0].(map[string]interface{})["uri"].(string)

	// GET /reports/stats.json
	var stats maxcdn.Generic
	res, err = max.Get(&stats, "/reports/stats.json", nil)
	if err != nil || res.Code != 200 || stats["total"] == "" {
		fmt.Println("GET stats")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", stats)
	}

	// GET /reports/stats.json/daily
	var daily maxcdn.Generic
	res, err = max.Get(&daily, "/reports/stats.json/daily", nil)
	if err != nil || res.Code != 200 || daily["total"] == "" {
		fmt.Println("GET stats/daily")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", daily)
	}

	// GET /zones/pull.json
	var getPulls maxcdn.Generic
	res, err = max.Get(&getPulls, "/zones/pull.json", nil)
	if err != nil || res.Code != 200 ||
		getPulls["pullzones"].([]interface{})[0].(map[string]interface{})["id"].(string) == "" {
		fmt.Println("GET zones/getPull")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", daily)
	}

	// note: order matters for the rest

	// POST /zones/pull.json
	var postPull maxcdn.Generic
	form = url.Values{}
	form.Set("name", name)
	form.Set("url", "http://www.example.com")
	res, err = max.Post(&postPull, "/zones/pull.json", form)
	if err != nil || res.Code != 201 ||
		postPull["pullzone"].(map[string]interface{})["name"].(string) == "" {
		fmt.Println("POST zones/pull")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", daily)
	}

	id := int(postPull["pullzone"].(map[string]interface{})["id"].(float64))

	// GET /zones/pull.json/{{zone_id}}
	var getPull maxcdn.Generic
	endpoint := fmt.Sprintf("/zones/pull.json/%d", id)
	res, err = max.Get(&getPull, endpoint, nil)
	if err != nil || res.Code != 200 ||
		getPull["pullzone"].(map[string]interface{})["id"].(string) == "" {
		fmt.Println("GET zones/getPull/{{zone_id}}")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", getPull)
	}

	// DELETE /zones/pull.json/{{zone_id}}
	res, err = max.Delete(endpoint, nil)
	if err != nil || res.Code != 200 {
		fmt.Println("DELETE zones/pull/{{zone_id}}")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
	}

	// PurgeZone
	if res, err := max.PurgeZone(id); err != nil || res.Code != 200 {
		fmt.Println("PurgeZone")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
	}

	// PurgeFile
	if res, err := max.PurgeFile(id, file); err != nil || res.Code != 200 {
		fmt.Println("PurgeFile")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
	}

	// Output:
}
