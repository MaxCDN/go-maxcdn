# go-maxcdn

MaxCDN Golang API.

## [API Documentation](http://godoc.org/github.com/jmervine/go-maxcdn)

```go
import "github.com/jmervine/go-maxcdn"
```
Package maxcdn is the golang bindings for MaxCDN's REST API.

This package should be considered beta. The final release will be moved to
`github.com/maxcdn/go-maxcdn`.

Developer Notes:

- Currently Pullzones does not support POST requests to
Endpoint.Zones.PullBy({zone_id}) as it returns mix types. Use Generic with type
assertions instead.
```go
    // Example:
	// Basic Get
	var data Account
	response, err := max.Get(&data, "/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Account.Name)

```
### Variables

```go
var APIHost = "https://rws.netdna.com"
```

APIHost is the hostname, including protocol, to MaxCDN's API.

```go
var Endpoint = endpoints{
    Account:        "/account.json",
    AccountAddress: "/account.json/address",
    Reports: &reports{
        PopularFiles: "/reports/popularfiles.json",
        Stats:        "/reports/stats.json",
    },
    Zones: &zones{
        Pull: "/zones/pull.json",
    },
}
```

Endpoint reflects all endpoints that are implemented as types and can be used as
data struct to be passed to request methods (e.g. Get, Put, etc.) for JSON
parsing. If the endpoint you are attempting to access isn't included in this
list, you'll need to use the Generic type, which uses an interface and type
assert the data values you wish to access.


### Types

#### Account
```go
type Account struct {
    Account struct {
        Alias                  string `json:"alias,omitempty"`
        DateCreated            string `json:"date_created,omitempty"`
        DateUpdated            string `json:"date_updated,omitempty"`
        DefaultPullZoneIpID    string `json:"default_pull_zone_ip_id,omitempty"`
        DefaultPushZoneIpID    string `json:"default_push_zone_ip_id,omitempty"`
        DefaultStorageIpID     string `json:"default_storage_ip_id,omitempty"`
        DefaultVodDirectIpID   string `json:"default_vod_direct_ip_id,omitempty"`
        DefaultVodPseudoIpID   string `json:"default_vod_pseudo_ip_id,omitempty"`
        DefaultVodRtmpIpID     string `json:"default_vod_rtmp_ip_id,omitempty"`
        DefaultVodStorageIpID  string `json:"default_vod_storage_ip_id,omitempty"`
        EdgerulesCredits       string `json:"edgerules_credits,omitempty"`
        FlexCredits            string `json:"flex_credits,omitempty"`
        ID                     string `json:"id,omitempty"`
        Name                   string `json:"name,omitempty"`
        SecureTokenPullCredits string `json:"secure_token_pull_credits,omitempty"`
        SslCredits             string `json:"ssl_credits,omitempty"`
        Status                 string `json:"status,omitempty"`
        StorageQuota           string `json:"storage_quota,omitempty"`
        ZoneCredits            string `json:"zone_credits,omitempty"`
    } `json:"account,omitempty"`
}
```


#### AccountAddress
```go
type AccountAddress struct {
    Address struct {
        City        string `json:"city,omitempty"`
        Country     string `json:"country,omitempty"`
        DateCreated string `json:"date_created,omitempty"`
        DateUpdated string `json:"date_updated,omitempty"`
        ID          string `json:"id,omitempty"`
        State       string `json:"state,omitempty"`
        Street1     string `json:"street1,omitempty"`
        Street2     string `json:"street2,omitempty"`
        Zip         string `json:"zip,omitempty"`
    } `json:"address,omitempty"`
}
```


#### Generic
```go
type Generic struct {
    Data map[string]interface{} `json:"data,omitempty"`
}
```


#### MaxCDN
```go
type MaxCDN struct {

    // MaxCDN Consumer Alias
    Alias string

    // Display raw http Request and Response for each http Transport
    Verbose bool

    HTTPClient *http.Client
    // contains filtered or unexported fields
}
```


#### NewMaxCDN
```go
func NewMaxCDN(alias, token, secret string) *MaxCDN
```
NewMaxCDN sets up a new MaxCDN instance.

```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	fmt.Printf("%#v\n", max)

```
#### Delete
```go
func (max *MaxCDN) Delete(endpoint string, form url.Values) (*Response, error)
```
Delete does an OAuth signed http.Delete

Delete does not take an endpointType because delete only returns a status code.


```go
    // Example:
	// This specific example shows how to purge a cache without using the Purge
	// methods, more as an example of using Delete, then anything, really.
	
	res, err := max.Delete(Endpoint.Zones.PullBy(123456), nil)
	if err != nil {
	    panic(err)
	}
	
	if res.Code == 200 {
	    fmt.Println("Purge suucceeded")
	}

```

#### Do
```go
func (max *MaxCDN) Do(method, endpoint string, form url.Values) (rsp *Response, err error)
```
Do is a low level method to interact with MaxCDN's RESTful API via Request and
return a parsed Response. It's used by all other methods.

This method closes the raw http.Response body.


```go
    // Example:
	// Run low level Do method.
	if rsp, err := max.Do("GET", "/account.json", nil); err == nil {
	    fmt.Printf("Response Code: %d\n", rsp.Code)
	
	    var data Account
	    if err = json.Unmarshal(rsp.Data, &data); err == nil {
	        fmt.Printf("%+v\n", data.Account)
	    }
	}

```

#### DoParse
```go
func (max *MaxCDN) DoParse(endpointType interface{}, method, endpoint string, form url.Values) (rsp *Response, err error)
```


```go
    // Example:
	// Run mid-level DoParse method.
	var data AccountAddress
	response, err := max.DoParse(&data, "GET", Endpoint.AccountAddress, nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Address.Street1)

```

#### Get
```go
func (max *MaxCDN) Get(endpointType interface{}, endpoint string, form url.Values) (*Response, error)
```
Get does an OAuth signed http.Get


```go
    // Example:
	var data AccountAddress
	response, err := max.Get(&data, Endpoint.AccountAddress, nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Address.Street1)

```

#### Post
```go
func (max *MaxCDN) Post(endpointType interface{}, endpoint string, form url.Values) (*Response, error)
```
Post does an OAuth signed http.Post


```go
    // Example:
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

```

#### PurgeFile
```go
func (max *MaxCDN) PurgeFile(zone int, file string) (*Response, error)
```
PurgeFile purges a specified file by zone from cache.


```go
    // Example:
	payload, err := max.PurgeFile(123456, "/master.css")
	if err != nil {
	    panic(err)
	}
	
	if payload.Code == 200 {
	    fmt.Println("Purge succeeded")
	}

```

#### PurgeFileString
```go
func (max *MaxCDN) PurgeFileString(zone string, file string) (*Response, error)
```
PurgeFile purges a specified file by zone from cache.



#### PurgeFiles
```go
func (max *MaxCDN) PurgeFiles(zone int, files []string) (resps []*Response, last error)
```
PurgeFiles purges multiple files from a zone.


```go
    // Example:
	files := []string{"/master.css", "/master.js"}
	payloads, err := max.PurgeFiles(123456, files)
	if err != nil {
	    panic(err)
	}
	
	if len(payloads) == len(files) {
	    fmt.Printf("Purges succeeded")
	}

```

#### PurgeZone
```go
func (max *MaxCDN) PurgeZone(zone int) (*Response, error)
```
PurgeZone purges a specified zones cache.


```go
    // Example:
	rsp, err := max.PurgeZone(123456)
	if err != nil {
	    panic(err)
	}
	
	if rsp.Code == 200 {
	    fmt.Println("Purge succeeded")
	}

```

#### PurgeZoneString
```go
func (max *MaxCDN) PurgeZoneString(zone string) (*Response, error)
```
PurgeZoneString purges a specified zones cache.



#### PurgeZones
```go
func (max *MaxCDN) PurgeZones(zones []int) (resps []*Response, last error)
```
PurgeZones purges multiple zones caches.


```go
    // Example:
	zones := []int{123456, 234567, 345678}
	rsps, err := max.PurgeZones(zones)
	if err != nil {
	    panic(err)
	}
	
	if len(rsps) == len(zones) {
	    fmt.Printf("Purges succeeded")
	}

```

#### PurgeZonesString
```go
func (max *MaxCDN) PurgeZonesString(zones []string) (resps []*Response, last error)
```
PurgeZonesString purges multiple zones caches.



#### Put
```go
func (max *MaxCDN) Put(endpointType interface{}, endpoint string, form url.Values) (*Response, error)
```
Put does an OAuth signed http.Put


```go
    // Example:
	form := url.Values{}
	form.Set("name", "example name")
	
	var data Account
	response, err := max.Put(&data, Endpoint.Account, form)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("code: %d\n", response.Code)
	fmt.Printf("name: %s\n", data.Account.Name)

```

#### Request
```go
func (max *MaxCDN) Request(method, endpoint string, form url.Values) (res *http.Response, err error)
```
Request is a low level method to interact with MaxCDN's RESTful API. It's used
by all other methods.

If using this method, you must manually close the res.Body or bad things may
happen.



#### PopularFiles
```go
type PopularFiles struct {
    CurrentPageSize int    `json:"current_page_size,omitempty"`
    Page            int    `json:"page,omitempty"`
    PageSize        string `json:"page_size,omitempty"`
    Pages           int    `json:"pages,omitempty"`
    PopularFiles    []struct {
        BucketID  string `json:"bucket_id,omitempty"`
        Hit       string `json:"hit,omitempty"`
        Size      string `json:"size,omitempty"`
        Timestamp string `json:"timestamp,omitempty"`
        Uri       string `json:"uri,omitempty"`
        Vhost     string `json:"vhost,omitempty"`
    } `json:"popularfiles,omitempty"`
    Summary struct {
        Hit  string `json:"hit,omitempty"`
        Size string `json:"size,omitempty"`
    } `json:"summary,omitempty"`
    Total string `json:"total,omitempty"`
}
```


#### Pullzone
```go
type Pullzone struct {
    Pullzone struct {
        BackendCompress       string `json:"backend_compress,omitempty"`
        CacheValid            string `json:"cache_valid,omitempty"`
        CanonicalLinkHeaders  string `json:"canonical_link_headers,omitempty"`
        CdnURL                string `json:"cdn_url,omitempty"`
        Compress              string `json:"compress,omitempty"`
        ContentDisposition    string `json:"content_disposition,omitempty"`
        CreationDate          string `json:"creation_date,omitempty"`
        DisallowRobots        string `json:"disallow_robots,omitempty"`
        DisallowRobotsTxt     string `json:"disallow_robots_txt,omitempty"`
        DnsCheck              string `json:"dns_check,omitempty"`
        Expires               string `json:"expires,omitempty"`
        HideSetcookieHeader   string `json:"hide_setcookie_header,omitempty"`
        ID                    string `json:"id,omitempty"`
        IgnoreCacheControl    string `json:"ignore_cache_control,omitempty"`
        IgnoreExpiresHeader   string `json:"ignore_expires_header,omitempty"`
        IgnoreSetcookieHeader string `json:"ignore_setcookie_header,omitempty"`
        Inactive              string `json:"inactive,omitempty"`
        Ip                    string `json:"ip,omitempty"`
        Label                 string `json:"label,omitempty"`
        Locked                string `json:"locked,omitempty"`
        Name                  string `json:"name,omitempty"`
        Port                  string `json:"port,omitempty"`
        ProxyCacheLock        string `json:"proxy_cache_lock,omitempty"`
        ProxyCacheLockTimeout string `json:"proxy_cache_lock_timeout,omitempty"`
        ProxyInactive         string `json:"proxy_inactive,omitempty"`
        PseudoStreaming       string `json:"pseudo_streaming,omitempty"`
        Queries               string `json:"queries,omitempty"`
        SetHostHeader         string `json:"set_host_header,omitempty"`
        Spdy                  string `json:"spdy,omitempty"`
        SpdyHeadersComp       string `json:"spdy_headers_comp,omitempty"`
        Sslshared             string `json:"sslshared,omitempty"`
        Suspend               string `json:"suspend,omitempty"`
        ThrottleFcc           string `json:"throttle_fcc,omitempty"`
        TmpURL                string `json:"tmp_url,omitempty"`
        Type                  string `json:"type,omitempty"`
        UpstreamEnabled       string `json:"upstream_enabled,omitempty"`
        URL                   string `json:"url,omitempty"`
        UseStale              string `json:"use_stale,omitempty"`
        ValidReferers         string `json:"valid_referers,omitempty"`
        WebpEnabled           string `json:"webp_enabled,omitempty"`
        XForwardFor           string `json:"x_forward_for,omitempty"`
    } `json:"pullzone,omitempty"`
}
```


#### Pullzones
```go
type Pullzones struct {
    CurrentPageSize int    `json:"current_page_size,omitempty"`
    Page            int    `json:"page,omitempty"`
    PageSize        string `json:"page_size,omitempty"`
    Pages           int    `json:"pages,omitempty"`
    Pullzones       []struct {
        BackendCompress       string `json:"backend_compress,omitempty"`
        CacheValid            string `json:"cache_valid,omitempty"`
        CanonicalLinkHeaders  string `json:"canonical_link_headers,omitempty"`
        CdnURL                string `json:"cdn_url,omitempty"`
        Compress              string `json:"compress,omitempty"`
        ContentDisposition    string `json:"content_disposition,omitempty"`
        CreationDate          string `json:"creation_date,omitempty"`
        DisallowRobots        string `json:"disallow_robots,omitempty"`
        DisallowRobotsTxt     string `json:"disallow_robots_txt,omitempty"`
        DnsCheck              string `json:"dns_check,omitempty"`
        Expires               string `json:"expires,omitempty"`
        HideSetcookieHeader   string `json:"hide_setcookie_header,omitempty"`
        ID                    string `json:"id,omitempty"`
        IgnoreCacheControl    string `json:"ignore_cache_control,omitempty"`
        IgnoreExpiresHeader   string `json:"ignore_expires_header,omitempty"`
        IgnoreSetcookieHeader string `json:"ignore_setcookie_header,omitempty"`
        Inactive              string `json:"inactive,omitempty"`
        Ip                    string `json:"ip,omitempty"`
        Label                 string `json:"label,omitempty"`
        Locked                string `json:"locked,omitempty"`
        Name                  string `json:"name,omitempty"`
        Port                  string `json:"port,omitempty"`
        ProxyCacheLock        string `json:"proxy_cache_lock,omitempty"`
        ProxyCacheLockTimeout string `json:"proxy_cache_lock_timeout,omitempty"`
        ProxyInactive         string `json:"proxy_inactive,omitempty"`
        PseudoStreaming       string `json:"pseudo_streaming,omitempty"`
        Queries               string `json:"queries,omitempty"`
        SetHostHeader         string `json:"set_host_header,omitempty"`
        Spdy                  string `json:"spdy,omitempty"`
        SpdyHeadersComp       string `json:"spdy_headers_comp,omitempty"`
        Sslshared             string `json:"sslshared,omitempty"`
        Suspend               string `json:"suspend,omitempty"`
        ThrottleFcc           string `json:"throttle_fcc,omitempty"`
        TmpURL                string `json:"tmp_url,omitempty"`
        Type                  string `json:"type,omitempty"`
        UpstreamEnabled       string `json:"upstream_enabled,omitempty"`
        URL                   string `json:"url,omitempty"`
        UseStale              string `json:"use_stale,omitempty"`
        ValidReferers         string `json:"valid_referers,omitempty"`
        WebpEnabled           string `json:"webp_enabled,omitempty"`
        XForwardFor           string `json:"x_forward_for,omitempty"`
    } `json:"pullzones,omitempty"`
    Total int `json:"total,omitempty"`
}
```


#### Response
```go
type Response struct {
    Code  int             `json:"code,omitempty"`
    Data  json.RawMessage `json:"data,omitempty"`
    Error struct {
        Message string `json:"message,omitempty"`
        Type    string `json:"type,omitempty"`
    } `json:"error,omitempty"`

    // Non-JSON data.
    Headers *http.Header
}
```


#### Stats
```go
type Stats struct {
    CurrentPageSize int    `json:"current_page_size,omitempty"`
    Page            int    `json:"page,omitempty"`
    PageSize        string `json:"page_size,omitempty"`
    Pages           int    `json:"pages,omitempty"`
    Stats           []struct {
        CacheHit    string `json:"cache_hit,omitempty"`
        Hit         string `json:"hit,omitempty"`
        NoncacheHit string `json:"noncache_hit,omitempty"`
        Size        string `json:"size,omitempty"`
        Timestamp   string `json:"timestamp,omitempty"`
    } `json:"stats,omitempty"`
    Summary struct {
        Stats struct {
            CacheHit    string `json:"cache_hit,omitempty"`
            Hit         string `json:"hit,omitempty"`
            NoncacheHit string `json:"noncache_hit,omitempty"`
            Size        string `json:"size,omitempty"`
            Timestamp   string `json:"timestamp,omitempty"`
        } `json:"stats,omitempty"`
        Total string `json:"total,omitempty"`
    } `json:"summary,omitempty"`
    Total string `json:"total,omitempty"`
}
```


#### StatsSummary
```go
type StatsSummary struct {
    Stats struct {
        CacheHit    string `json:"cache_hit,omitempty"`
        Hit         string `json:"hit,omitempty"`
        NoncacheHit string `json:"noncache_hit,omitempty"`
        Size        string `json:"size,omitempty"`
    } `json:"stats,omitempty"`
    Total string `json:"total,omitempty"`
}
```



