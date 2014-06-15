// This file contains Example and Functional Example methods, both for documentation
// and for functionally testing code.

package maxcdn

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/url"
	"os"
	//"strconv"
)

var (
	alias  = os.Getenv("ALIAS")
	token  = os.Getenv("TOKEN")
	secret = os.Getenv("SECRET")
)

/****
 * Documentation Examples
 *******************************************************************/

func Example() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	// Basic Get
	var data Account
	response, err := max.Get(&data, "/account.json", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Account.Name)
}

func ExampleNewMaxCDN() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	fmt.Printf("%#v\n", max)
}

/*

func ExampleMaxCDN_Do() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	// Run low level Do method.
	res, err := max.Do("GET", "/account.json", nil)

	// Raw http.Response requires that you close the Body.
	defer res.Body.Close()

	if err != nil {
		panic(fmt.Errorf("[%s] %v", res.Status, err))
	}

	// Read the Body.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", body)
}

*/

func ExampleMaxCDN_Get() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	var data AccountAddress
	response, err := max.Get(&data, Endpoint.AccountAddress, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Address.Street1)
}

func ExampleMaxCDN_Put() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	form := url.Values{}
	form.Set("name", "example name")

	var data Account
	response, err := max.Put(&data, Endpoint.Account, form)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Account.Name)
}

func ExampleMaxCDN_Delete() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	// This specific example shows how to purge a cache without using the Purge
	// methods, more as an example of using Delete, then anything, really.

	res, err := max.Delete(Endpoint.Zones.PullBy(123456), nil)
	if err != nil {
		panic(err)
	}

	if res.Code == 200 {
		fmt.Println("Purge suucceeded")
	}
}

func ExampleMaxCDN_PurgeZone() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	rsp, err := max.PurgeZone(123456)
	if err != nil {
		panic(err)
	}

	if rsp.Code == 200 {
		fmt.Println("Purge succeeded")
	}
}

func ExampleMaxCDN_PurgeZones() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	zones := []int{123456, 234567, 345678}
	rsps, err := max.PurgeZones(zones)
	if err != nil {
		panic(err)
	}

	if len(rsps) == len(zones) {
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

func ExampleMaxCDN_Post() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))

	form := url.Values{}
	form.Set("name", "newzone")

	// When creating a new zone, the url must be real and resolve.
	form.Set("url", "http://www.example.com")

	var data Pullzone
	_, err := max.Post(&data, Endpoint.Zones.Pull, form)
	if err != nil {
		panic(err)
	}

	if data.Pullzone.Name == "newzone" {
		fmt.Println("Successfully created new Pull Zone.")
	}
}

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

func Example_Functional_MaxCDN_Get() {
	max := NewMaxCDN(alias, token, secret)

	if alias == "" || token == "" || secret == "" {
		max.HTTPClient = stubHTTPOk()
	}

	var data Account
	_, err := max.Get(&data, Endpoint.Account, nil)
	if err != nil {
		panic(err)
	}

	if data.Account.Name != "" {
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

	var data Account
	_, err := max.Put(&data, Endpoint.Account, form)

	if err != nil {
		panic(err)
	}

	if data.Account.Name == name {
		fmt.Println("PUT /account.json succeeded")
	}

	// Output: PUT /account.json succeeded
}

func Example_Functional_MaxCDN_Post() {
	// I'm using a timestamp as unique name for testing, you shouldn't
	// do this.
	name := stringFromTimestamp()

	max := NewMaxCDN(alias, token, secret)

	// This won't work until MaxCDN fixes the API response to match the
	// response that GET /zones/pull.json/{zone_id}
	if alias != "" || token != "" || secret != "" {
		max.HTTPClient = stubHTTPOk()
		name = "cdn-example-net"

		form := url.Values{}
		form.Set("name", name)

		// When creating a new zone, the url must be real and resolve.
		form.Set("url", "http://www.example.com")

		var data Pullzone
		_, err := max.Post(&data, Endpoint.Zones.Pull, form)
		if err != nil {
			panic(err)
		}

		if data.Pullzone.Name == name {
			fmt.Println("POST /zones/pull.json succeeded")

			id := data.Pullzone.ID

			rsp, err := max.Delete(Endpoint.Zones.PullByString(id), nil)
			if err != nil {
				panic(err)
			}

			if rsp.Code == 200 {
				fmt.Println("DELETE /zones/pull.json succeeded")
			}
		}
	} else {
		fmt.Println("POST /zones/pull.json succeeded\nDELETE /zones/pull.json succeeded")
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
	var data Pullzones
	rsp, err := max.Get(&data, Endpoint.Zones.Pull, nil)
	if err != nil {
		panic(err)
	}

	zone_id := data.Pullzones[0].ID

	// Now purge that zone's cache.
	rsp, err = max.PurgeZoneString(zone_id)
    if err != nil {
        panic(err)
    }

	if rsp.Code == 200 {
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
	var data Pullzones
	rsp, err := max.Get(&data, Endpoint.Zones.Pull, nil)
	if err != nil {
		panic(err)
	}

	zone_id := data.Pullzones[0].ID

	// Next fetch file name.
	var files PopularFiles
	rsp, err = max.Get(&files, Endpoint.Reports.PopularFiles, nil)
	if err != nil {
		panic(err)
	}

	file := files.Popularfiles[0].Uri

	// Now purge that zone's cache.
	rsp, err = max.PurgeFileString(zone_id, file)
	if err != nil {
		panic(err)
	}

	if rsp.Code == 200 {
		fmt.Println("Purge succeeded")
	}

	// Output: Purge succeeded
}
