package maxcdn

import "fmt"

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
