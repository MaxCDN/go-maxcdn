// This file contains Example and Functional Example methods, both for documentation
// and for functionally testing code.

package maxcdn

import (
//"encoding/json"
"fmt"
//"net/url"
"os"
//"strconv"
)

/****
 * Documentation Examples
 *******************************************************************/

func Example() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	// Basic Get
	payload, err := max.Get("/account.json", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", payload)
}

/*
func ExampleNewMaxCDN() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	fmt.Printf("%#v\n", max)
}

func ExampleMaxCDN_Do() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	raw, res, err := max.Do("GET", "/account.json", nil)

	if err != nil {
		panic(fmt.Errorf("[%s] %v", res.Status, err))
	}

	mapper := GenericResponse{}
	mapper.Response = res
	mapper.Raw = raw

	err = json.Unmarshal(raw, &mapper)
	if err != nil {
		panic(err)
	}

	if mapper.Error.Message != "" || mapper.Error.Type != "" {
		err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}
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

func ExampleMaxCDN_PurgeFile() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	payload, err := max.PurgeFile(123456, "/master.css")
	if err != nil {
		panic(err)
	}

	if payload.Code == 200 {
		fmt.Println("Purge succeeded")
	}
}

func ExampleMaxCDN_PurgeFiles() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	files := []string{"/master.css", "/master.js"}
	payloads, err := max.PurgeFiles(123456, files)
	if err != nil {
		panic(err)
	}

	if len(payloads) == len(files) {
		fmt.Printf("Purges succeeded")
	}
}
*/
/****
 * Functional Examples
 *
 * All "Functional" examples are meant to be functional integration tests.
 * To run these as integration tests export your ALIAS, TOKEN and SECRET
 * to your envioronment before running 'go test', otherwise the http
 * request will be stubbed using the json files in './_fixtures/*.json'.
 *
 * Run like:
 *
 * $ ALIAS=your_alias TOKEN=your_token SECRET=your_secret go test
 *******************************************************************/
/*
func Example_Functional_MaxCDN_Get() {
	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HTTPClient = stubHTTPOk()
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

func Example_Functional_MaxCDN_Put() {
	// I'm using a timestamp as unique name for testing, you shouldn't
	// do this.
	name := stringFromTimestamp()

	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HTTPClient = stubHTTPOk()
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

func Example_Functional_MaxCDN_Post() {
	// I'm using a timestamp as unique name for testing, you shouldn't
	// do this.
	name := stringFromTimestamp()

	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HTTPClient = stubHTTPOk()
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

func Example_Functional_MaxCDN_PurgeZone() {
	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HTTPClient = stubHTTPOk()
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

func Example_Functional_MaxCDN_PurgeFile() {

	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HTTPClient = stubHTTPOk()
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

	// Next fetch file name.
	payload, err = max.Get("/reports/pull/popularfiles.json", nil)
	if err != nil {
		panic(err)
	}

	data = payload.Data["popularfiles"].([]interface{})

	file := data[0].(map[string]interface{})["uri"].(string)

	// Now purge that zone's cache.
	payload, err = max.PurgeFile(int(id), file)
	if err != nil {
		panic(err)
	}

	if payload.Code == 200 {
		fmt.Println("Purge succeeded")
	}

	// Output: Purge succeeded
}
*/
