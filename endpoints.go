package maxcdn

// This file contains the struct holding all implemented endpoint data.

import "fmt"

type endpoints struct {
	Account        string
	AccountAddress string
	Reports        *reports
	Zones          *zones
	Users          string
}

// User: Endpoint.User("johndoe")
// => /users.json/johndoe
func (e *endpoints) User(userId string) string {
	return fmt.Sprintf("%s/%s", e.Users, userId)
}

type reports struct {
	PopularFiles string
	Stats        string
}

type zones struct {
	Pull string
	//Push string
}

// StatsBy:
//
// Endpoint.Reports.StatsBy("hourly")
// => /reports/stats.json/hourly
//
// Endpoint.Reports.StatsBy("daily")
// => /reports/stats.json/daily
//
// Endpoint.Reports.StatsBy("monthly")
// => /reports/stats.json/monthly
func (r *reports) StatsBy(t string) string {
	return fmt.Sprintf("%s/%s", r.Stats, t)
}

// PullBy: Endpoint.Zones.PullBy("123456")
// => /zones/pull.json/123456
func (z *zones) PullBy(t int) string {
	return fmt.Sprintf("%s/%d", z.Pull, t)
}

// PullByString: Endpoint.Zones.PullByString("123456")
// => /zones/pull.json/123456
func (z *zones) PullByString(t string) string {
	return fmt.Sprintf("%s/%s", z.Pull, t)
}

// PullCacheByString: Endpoint.Zones.PullByStringCache(123456)
// => /zones/pull.json/123456/cache
func (z *zones) PullCacheBy(t int) string {
	return fmt.Sprintf("%s/%d/cache", z.Pull, t)
}

// PushByString: Endpoint.Zones.PullByStringCache("123456")
// => /zones/pull.json/123456/cache
func (z *zones) PullCacheByString(t string) string {
	return fmt.Sprintf("%s/%s/cache", z.Pull, t)
}

// PushByString: Endpoint.Zones.PushBy(123456)
// => /zones/push.json/123456
//func (z *zones) PushBy(t int) string {
//return fmt.Sprintf("%s/%d", z.Push, t)
//}

// PushByString: Endpoint.Zones.PushByString("123456")
// => /zones/push.json/123456
//func (z *zones) PushByString(t string) string {
//return fmt.Sprintf("%s/%s", z.Push, t)
//}

// Endpoint
//
// This reflects all endpoints that are implemented as types and can be used
// as data struct to be passed to request methods (e.g. Get, Put, etc.) for
// JSON parsing. If the endpoint you are attempting to access isn't included
// in this list, you'll need to use the Generic type, which uses an interface
// and type assert the data values you wish to access.
//
// Endpoint examples:
//
//  // for pull zone with id of '123456'
//  e := Endpoint.Zones.PullBy('123456')
//  => /zones/pull.json/123456
//
//  // for popular files report
//  e := Endpoint.Reports.PopularFiles
//  => /reports/popularfiles.json
//
//  // for hourly stats report
//  e := Endpoint.Reports.StatsBy('hourly')
//  => /reports/stats.json/hourly
var Endpoint = endpoints{
	Account:        "/account.json",
	AccountAddress: "/account.json/address",
	Reports: &reports{
		PopularFiles: "/reports/popularfiles.json",
		Stats:        "/reports/stats.json",
	},
	Zones: &zones{
		Pull: "/zones/pull.json",
		//Push: "/zones/push.json",
	},
}
