package maxcdn

import (
	"fmt"
	"testing"

	. "github.com/jmervine/GoT"
)

func TestEndpoint(T *testing.T) {
	//Account
	Go(T).AssertEqual(Endpoint.Account, "/account.json")
	Go(T).AssertEqual(Endpoint.AccountAddress, "/account.json/address")

	//Reports
	Go(T).AssertEqual(Endpoint.Reports.PopularFiles, "/reports/popularfiles.json")
	Go(T).AssertEqual(Endpoint.Reports.Stats, "/reports/stats.json")
	Go(T).AssertEqual(Endpoint.Reports.StatsBy("hourly"), "/reports/stats.json/hourly")

	//Zones
	Go(T).AssertEqual(Endpoint.Zones.Pull, "/zones/pull.json")
	Go(T).AssertEqual(Endpoint.Zones.PullBy(12345), "/zones/pull.json/12345")
	Go(T).AssertEqual(Endpoint.Zones.PullByString("12345"), "/zones/pull.json/12345")
	Go(T).AssertEqual(Endpoint.Zones.PullCacheBy(12345), "/zones/pull.json/12345/cache")
	Go(T).AssertEqual(Endpoint.Zones.PullCacheByString("12345"), "/zones/pull.json/12345/cache")
}

func ExampleEndpoint() {
	// for pull zone with id of '123456'
	fmt.Printf(" => %s \n", Endpoint.Zones.PullBy(123456))
	// => /zones/pull.json/123456

	// for popular files report
	fmt.Printf(" => %s \n", Endpoint.Reports.PopularFiles)
	// => /reports/popularfiles.json

	// for hourly stats report
	fmt.Printf(" => %s \n", Endpoint.Reports.StatsBy("hourly"))
	//=> /reports/stats.json/hourly
}
