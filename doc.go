/*
Package maxcdn is the golang bindings for MaxCDN's REST API.

This package should be considered beta. The final release will be moved to
`github.com/maxcdn/go-maxcdn`.

Developer Notes:

- Currently Pullzones does not support POST requests to Endpoint.Zones.PullBy({zone_id}) as it returns mix types. Use Generic with type assertions instead.

*/
package maxcdn
