package maxcdn

// This file contains Example and Functional Example methods, both for documentation
// and for functionally testing code.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

var (
	alias  = os.Getenv("ALIAS")
	token  = os.Getenv("TOKEN")
	secret = os.Getenv("SECRET")
	max    = NewMaxCDN(alias, secret, token)
)

/****
 * Examples
 *******************************************************************/

func Example() {
	// Basic Get
	var got Account
	res, err := max.Get(&got, Endpoint.Account, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", res.Code)
	fmt.Printf("name: %s\n", got.Account.Name)

	// Basic Put
	form := url.Values{}
	form.Set("name", "new name")

	var put Account
	if _, err = max.Put(&put, Endpoint.Account, form); err == nil && put.Account.Name == "new name" {
		fmt.Println("name successfully updated")
	}

	// Basic Delete
	if _, err = max.Delete(Endpoint.Zones.PullBy(123456), nil); err == nil {
		fmt.Println("zone successfully deleted")
	}

	// Generic data type
	var data Generic
	if _, err := max.Get(&data, Endpoint.Account, nil); err == nil {
		alias := data["alias"].(string)
		fmt.Printf("alias: %s\n", alias)
	}
}

/****
 * MaxCDN Examples
 *******************************************************************/

func ExampleNewMaxCDN() {
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	fmt.Printf("%#v\n", max)
}

func ExampleMaxCDN_DoParse() {
	// Run mid-level DoParse method.
	var data AccountAddress
	response, err := max.DoParse(&data, "GET", Endpoint.AccountAddress, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Address.Street1)
}

func ExampleMaxCDN_Do() {
	// Run low-level Do method.
	if rsp, err := max.Do("GET", "/account.json", nil); err == nil {
		fmt.Printf("Response Code: %d\n", rsp.Code)

		var data Account
		if err = json.Unmarshal(rsp.Data, &data); err == nil {
			fmt.Printf("%+v\n", data.Account)
		}
	}
}

func ExampleMaxCDN_Request() {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	// Run low-level Request method.
	res, err := max.Request("GET", Endpoint.Account, nil)
	defer res.Body.Close()
	check(err)

	body, err := ioutil.ReadAll(res.Body)
	check(err)

	fmt.Println(string(body))
}

func ExampleMaxCDN_Get() {
	var data AccountAddress
	response, err := max.Get(&data, Endpoint.AccountAddress, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Address.Street1)
}

func ExampleMaxCDN_Put() {
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
	rsp, err := max.PurgeZone(123456)
	if err != nil {
		panic(err)
	}

	if rsp.Code == 200 {
		fmt.Println("Purge succeeded")
	}
}

func ExampleMaxCDN_PurgeZones() {
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
	payload, err := max.PurgeFile(123456, "/master.css")
	if err != nil {
		panic(err)
	}

	if payload.Code == 200 {
		fmt.Println("Purge succeeded")
	}
}

func ExampleMaxCDN_PurgeFiles() {
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
 * Type Examples
 *******************************************************************/

func ExampleResponse() {
	var data Account
	response, _ := max.Get(&data, Endpoint.Account, nil)
	fmt.Printf("%+v\n", response)
}

func ExampleGeneric() {
	var data Generic
	if _, err := max.Get(&data, Endpoint.Account, nil); err == nil {
		alias := data["alias"].(string)
		name := data["name"].(string)
		fmt.Printf("alias: %s\n", alias)
		fmt.Printf("name:  %s\n", name)
	}
}

func ExampleAccount() {
	var data Account
	if _, err := max.Get(&data, Endpoint.Account, nil); err == nil {
		fmt.Printf("%+v\n", data.Account)
	}
}

func ExampleAccountAddress() {
	var data AccountAddress
	if _, err := max.Get(&data, Endpoint.AccountAddress, nil); err == nil {
		fmt.Printf("%+v\n", data.Address)
	}
}

func ExamplePopularFiles() {
	var data PopularFiles
	if _, err := max.Get(&data, Endpoint.Reports.PopularFiles, nil); err == nil {
		for i, file := range data.PopularFiles {
			fmt.Printf("%2d: %30s=%s, \n", i, file.Uri, file.Hit)
		}
	}
	fmt.Println("----")
	fmt.Printf("    %30s=%s, \n", "summary", data.Summary.Hit)
}

func ExampleStatsSummary() {
	var data StatsSummary
	if _, err := max.Get(&data, Endpoint.Reports.Stats, nil); err == nil {
		fmt.Printf("%+v\n", data.Stats)
	}
}

func ExampleStats() {
	var data Stats
	if _, err := max.Get(&data, Endpoint.Reports.StatsBy("hourly"), nil); err == nil {
		fmt.Printf("%+v\n", data.Stats)
		fmt.Printf("%+v\n", data.Summary)
	}
}


/****
 * These "Functional" examples are meant to be functional integration tests.
 * To run these as integration tests export your ALIAS, TOKEN and SECRET
 * to your envioronment before running 'go test', otherwise the http
 * request will be stubbed using the json files in './_fixtures/*.json'.
 *
 * Run like:
 *
 * $ ALIAS=your_alias TOKEN=your_token SECRET=your_secret make test
 *******************************************************************/
func Example_Functional() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	var form url.Values

	name := stringFromTimestamp()

	max := NewMaxCDN(alias, token, secret)
	if alias == "" || token == "" || secret == "" {
		max.HTTPClient = stubHTTPOk()
	}

	// GET Account
	var getAcct Account
	if res, err := max.Get(&getAcct, Endpoint.Account, nil); err != nil || res.Code != 200 || getAcct.Account.Name == "" {
		fmt.Println("GET account")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", getAcct)
	}

	// PUT Account
	var putAcct Account
	form = url.Values{}
	form.Set("name", name)
	if res, err := max.Put(&putAcct, Endpoint.Account, form); err != nil || res.Code != 200 || putAcct.Account.Name == "" {
		fmt.Println("PUT account")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", putAcct)
	}

	// GET AccountAddress
	var getAddr AccountAddress
	if res, err := max.Get(&getAddr, Endpoint.AccountAddress, nil); err != nil || res.Code != 200 || getAddr.Address.Street1 == "" {
		fmt.Println("GET address")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", getAddr)
	}

	// PUT AccountAddress
	var putAddr AccountAddress
	form = url.Values{}
	form.Set("street1", name)
	if res, err := max.Put(&putAddr, Endpoint.AccountAddress, form); err != nil || res.Code != 200 || putAddr.Address.Street1 == "" {
		fmt.Println("GET address")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", putAddr)
	}

	// Reports/PopularFiles
	var popular PopularFiles
	if res, err := max.Get(&popular, Endpoint.Reports.PopularFiles, nil); err != nil || res.Code != 200 || popular.PopularFiles[0].Uri == "" {
		fmt.Println("GET popularfiles")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", popular)
	}

	file := popular.PopularFiles[0].Uri

	// Reports/Stats
	var stats StatsSummary
	if res, err := max.Get(&stats, Endpoint.Reports.Stats, nil); err != nil || res.Code != 200 || stats.Total == "" {
		fmt.Println("GET stats")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", stats)
	}

	// Reports/Stats/Daily
	var daily Stats
	if res, err := max.Get(&daily, Endpoint.Reports.StatsBy("daily"), nil); err != nil || res.Code != 200 || daily.Total == "" {
		fmt.Println("GET stats/daily")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", daily)
	}

	// GET Zones/Pull
	var getPulls Pullzones
	if res, err := max.Get(&getPulls, Endpoint.Zones.Pull, nil); err != nil || res.Code != 200 || getPulls.Pullzones[0].ID == "" {
		fmt.Println("GET zones/getPull")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", daily)
	}

	// note: order matters for the rest

	// POST Zones/Pull
	var postPull Generic // until api is fixed
	form = url.Values{}
	form.Set("name", name)
	form.Set("url", "http://www.example.com")
	if res, err := max.Post(&postPull, Endpoint.Zones.Pull, form); err != nil || res.Code != 201 || postPull["pullzone"].(map[string]interface{})["name"].(string) == "" {
		fmt.Println("POST zones/pull")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", daily)
	}

	id := int(postPull["pullzone"].(map[string]interface{})["id"].(float64))

	// GET Zones/Pull/{{zone_id}}
	var getPull Pullzone
	if res, err := max.Get(&getPull, Endpoint.Zones.PullBy(id), nil); err != nil || res.Code != 200 || getPull.Pullzone.ID == "" {
		fmt.Println("GET zones/getPull/{{zone_id}}")
		fmt.Printf("error:\n%+v\n", err)
		fmt.Printf("code:\n%+v\n", res.Code)
		fmt.Printf("data:\n%+v\n", getPull)
	}

	// DELETE Zones/Pull
	if res, err := max.Delete(Endpoint.Zones.PullBy(id), nil); err != nil || res.Code != 200 {
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
