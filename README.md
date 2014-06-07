# go-maxcdn

MaxCDN Golang API.

## [API Documentation](http://godoc.org/github.com/jmervine/go-maxcdn)

```go
import "github.com/jmervine/go-maxcdn"
```
Package maxcdn is the golang bindings for MaxCDN's REST API.

This package should be considered beta. The final release will be moved to
`github.com/maxcdn/go-maxcdn`.
```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	// Basic Get
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)
	
	// Below is pretty much exactly what 'maxcdn.Get' is doing.
	// The purpose though would be for you to generate your
	// own struct more exactly mapping the json response to
	// your purpose. More specific responses are planned for
	// future versions, but there are too many make it worth
	// implementing all of them, so this support should remain.
	raw, res, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    panic(fmt.Errorf("[%s] %v", res.Status, err))
	}
	
	mapper := GenericResponse{}
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```
### Constants

```go
const (
    PopularFilesEndpoint = "/reports/popularfiles.json"
    StatsEndpoint        = "/reports/stats.json"
)
```

> Here lies known endpoints.


### Variables

```go
var APIHost = "https://rws.netdna.com"
```

> APIHost is the hostname, including protocol, to MaxCDN's API.


### Types

#### GenericResponse
```go
type GenericResponse struct {
    Code  int                    `json:"code"`
    Data  map[string]interface{} `json:"data"`
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
    } `json:"error"`
    Raw      []byte         // include raw json in GenericResponse
    Response *http.Response // include response in GenericResponse
}
```


#### Parse
```go
func (mapper *GenericResponse) Parse(raw []byte) (err error)
```
> Parse turns an http response in to a GenericResponse



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
> NewMaxCDN sets up a new MaxCDN instance.

```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	fmt.Printf("%#v\n", max)

```
#### Delete
```go
func (max *MaxCDN) Delete(endpoint string) (mapper *GenericResponse, err error)
```
> Delete does an OAuth signed http.Delete


```go
    // Example:
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

```

#### Do
```go
func (max *MaxCDN) Do(method, endpoint string, form url.Values) (raw []byte, res *http.Response, err error)
```
> Do is a generic method to interact with MaxCDN's RESTful API. It's used by
> all other methods.

> It's purpose though would be for you to generate your own struct more
> exactly mapping the json response to your purpose. More specific responses
> are planned for future versions, but there are too many make it worth
> implementing all of them, so this support should remain.


```go
    // Example:
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

```

#### Get
```go
func (max *MaxCDN) Get(endpoint string, form url.Values) (mapper *GenericResponse, err error)
```
> Get does an OAuth signed http.Get


```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### GetDailyStats
```go
func (max *MaxCDN) GetDailyStats(form url.Values) (mapper *MultiStats, err error)
```
> GetDailyStats does an OAuth signed http.Get for "/reports/stats.json/hourly"



#### GetHourlyStats
```go
func (max *MaxCDN) GetHourlyStats(form url.Values) (mapper *MultiStats, err error)
```
> GetHourlyStats does an OAuth signed http.Get for
> "/reports/stats.json/hourly"



#### GetMonthlyStats
```go
func (max *MaxCDN) GetMonthlyStats(form url.Values) (mapper *MultiStats, err error)
```
> GetMonthlyStats does an OAuth signed http.Get for
> "/reports/stats.json/hourly"



#### GetPopularFiles
```go
func (max *MaxCDN) GetPopularFiles(form url.Values) (mapper *PopularFiles, err error)
```
> Get does an OAuth signed http.Get for "/reports/popularfiles.json"



#### GetStatsByType
```go
func (max *MaxCDN) GetStatsByType(report string, form url.Values) (mapper *MultiStats, err error)
```
> GetHourlyStats does an OAuth signed http.Get for
> "/reports/stats.json/{report_type}".

> Valid report types are; 'hourly', 'daily' and 'monthly'



#### GetStatsSummary
```go
func (max *MaxCDN) GetStatsSummary(form url.Values) (mapper *SummaryStats, err error)
```
> Get does an OAuth signed http.Get for "/reports/stats.json"



#### Post
```go
func (max *MaxCDN) Post(endpoint string, form url.Values) (mapper *GenericResponse, err error)
```
> Post does an OAuth signed http.Post


```go
    // Example:
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

```

#### PurgeFile
```go
func (max *MaxCDN) PurgeFile(zone int, file string) (mapper *GenericResponse, err error)
```
> PurgeFile purges a specified file by zone from cache.


```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.PurgeFile(123456, "/master.css")
	if err != nil {
	    panic(err)
	}
	
	if payload.Code == 200 {
	    fmt.Println("Purge succeeded")
	}

```

#### PurgeFiles
```go
func (max *MaxCDN) PurgeFiles(zone int, files []string) (resps []*GenericResponse, last error)
```
> PurgeFiles purges multiple files from a zone.


```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
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
func (max *MaxCDN) PurgeZone(zone int) (*GenericResponse, error)
```
> PurgeZone purges a specified zones cache.


```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.PurgeZone(123456)
	if err != nil {
	    panic(err)
	}
	
	if payload.Code == 200 {
	    fmt.Println("Purge succeeded")
	}

```

#### PurgeZones
```go
func (max *MaxCDN) PurgeZones(zones []int) (resps []*GenericResponse, last error)
```
> PurgeZones purges multiple zones caches.


```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	zones := []int{123456, 234567, 345678}
	payloads, err := max.PurgeZones(zones)
	if err != nil {
	    panic(err)
	}
	
	if len(payloads) == len(zones) {
	    fmt.Printf("Purges succeeded")
	}

```

#### Put
```go
func (max *MaxCDN) Put(endpoint string, form url.Values) (mapper *GenericResponse, err error)
```
> Put does an OAuth signed http.Put


```go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	form := url.Values{}
	form.Set("name", "example_name")
	payload, err := max.Put("/account.json", form)
	
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### MultiStats
```go
type MultiStats struct {
    Code int `json:"code"`
    Data struct {
        CurrentPageSize int     `json:"current_page_size"`
        Page            int     `json:"page"`
        PageSize        string  `json:"page_size"`
        Pages           int     `json:"pages"`
        Stats           []Stats `json:"stats"`
        Summary         Stats   `json:"summary"`
        Total           string  `json:"total"`
    } `json:"data"`
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
    } `json:"error"`
    Raw      []byte
    Response *http.Response
}
```


#### Parse
```go
func (mapper *MultiStats) Parse(raw []byte) (err error)
```
> Parse turns an http response in to a Stats



#### PopularFiles
```go
type PopularFiles struct {
    Code int `json:"code"`
    Data struct {
        CurrentPageSize int    `json:"current_page_size"`
        Page            int    `json:"page"`
        PageSize        string `json:"page_size"`
        Pages           int    `json:"pages"`
        Popularfiles    []struct {
            BucketID  string `json:"bucket_id"`
            Hit       string `json:"hit"`
            Size      string `json:"size"`
            Timestamp string `json:"timestamp"`
            Uri       string `json:"uri"`
            Vhost     string `json:"vhost"`
        } `json:"popularfiles"`
        Summary struct {
            Hit  string `json:"hit"`
            Size string `json:"size"`
        } `json:"summary"`
        Total string `json:"total"`
    } `json:"data"`

    // Added for extra support, see maxcdn.GenericResponse
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
    } `json:"error"`
    Raw      []byte
    Response *http.Response
}
```


#### Parse
```go
func (mapper *PopularFiles) Parse(raw []byte) (err error)
```
> Parse turns an http response in to a PopularFiles



#### Stats
```go
type Stats struct {
    CacheHit    string `json:"cache_hit"`
    Hit         string `json:"hit"`
    NoncacheHit string `json:"noncache_hit"`
    Size        string `json:"size"`
    Timestamp   string `json:"timestamp"`
}
```


#### SummaryStats
```go
type SummaryStats struct {
    Code int `json:"code"`
    Data struct {
        Stats Stats `json:"stats"`
        //Stats struct {
        //CacheHit    string `json:"cache_hit"`
        //Hit         string `json:"hit"`
        //NoncacheHit string `json:"noncache_hit"`
        //Size        string `json:"size"`
        //Timestamp   string `json:"timestamp"`
        //} `json:"stats"`
        Total string `json:"total"`
    } `json:"data"`
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
    } `json:"error"`
    Raw      []byte
    Response *http.Response
}
```


#### Parse
```go
func (mapper *SummaryStats) Parse(raw []byte) (err error)
```
> Parse turns an http response in to a StatsSummary




