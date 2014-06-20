package maxcdn

import (
	"net/http"
	"testing"

	. "github.com/jmervine/GoT"
)

func TestAccount(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data Account
	_, err := max.DoParse(&data, "GET", "/account.json", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.Account.Alias, "aliasname")
}

func TestAccountAddress(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data AccountAddress
	_, err := max.DoParse(&data, "GET", "/account.json/address", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.Address.Zip, "91604")
}

func TestPopularFiles(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data PopularFiles
	_, err := max.DoParse(&data, "GET", "/reports/popularfiles.json", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.PopularFiles[0].Uri, "/master.css")
}

func TestStatsSummary(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data StatsSummary
	_, err := max.DoParse(&data, "GET", "/reports/stats.json", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.Stats.Hit, "18632")
}

func TestStats(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data Stats
	_, err := max.DoParse(&data, "GET", "/reports/stats.json/daily", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.Stats[0].Hit, "267")
}

func TestPullzone(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data Pullzone
	_, err := max.DoParse(&data, "GET", "/zones/pull.json/ZONE_ID", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.Pullzone.Name, "cdn-example-net")
}

func TestPullzones(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data Pullzones
	_, err := max.DoParse(&data, "GET", "/zones/pull.json", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.Pullzones[0].Name, "cdn-example-net")
}

func TestUsers(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data Users
	_, err := max.DoParse(&data, "GET", "/users.json", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.Users[0].Firstname, "Josh")
}

func TestUser(T *testing.T) {
	max := NewMaxCDN("alias", "token", "secret")

	var recorder http.Response
	max.HTTPClient = stubHTTPOkRecorded(&recorder)

	var data User
	_, err := max.DoParse(&data, "GET", "/users.json/USER_ID", nil)
	Go(T).AssertNil(err)
	Go(T).AssertEqual(data.User.Firstname, "joshua")
}

