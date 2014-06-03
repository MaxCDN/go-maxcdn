# go-maxcdn

MaxCDN Golang API.

> Note: This is **very** alpha. Final release will be at github.com/maxcdn/go-maxcdn.

## [API Documentation](http://godoc.org/github.com/jmervine/go-maxcdn)

```go
import "github.com/jmervine/go-maxcdn"
```
Package maxcdn is the golang bindings for MaxCDN's REST API.

At this time it should be considered very alpha.

##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

### Constants

```go
const APIHost = "https://rws.netdna.com"
```

> APIHost is the hostname, including protocol, to MaxCDN's API.


### Types

#### GenericResponse

```go
type GenericResponse struct {
    Code  float64                `json:"code"`
    Data  map[string]interface{} `json:"data"`
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
    } `json:"error"`
}
```



#### MaxCDN

```go
type MaxCDN struct {
    Alias string

    HTTPClient *http.Client
    // contains filtered or unexported fields
}
```



#### NewMaxCDN

```go
func NewMaxCDN(alias, token, secret string) *MaxCDN
```
> NewMaxCDN sets up a new MaxCDN instance.

##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	fmt.Printf("%#v\n", max)

#### Delete

```go
func (max *MaxCDN) Delete(endpoint string) (*GenericResponse, error)
```
> Delete does an OAuth signed http.Delete



##### Example:
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


#### Get

```go
func (max *MaxCDN) Get(endpoint string, form url.Values) (*GenericResponse, error)
```
> Get does an OAuth signed http.Get



##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)


#### Post

```go
func (max *MaxCDN) Post(endpoint string, form url.Values) (*GenericResponse, error)
```
> Post does an OAuth signed http.Post



##### Example:
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


#### PurgeZone

```go
func (max *MaxCDN) PurgeZone(zone int) (*GenericResponse, error)
```
> PurgeZone purges a specified zones cache.



##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.PurgeZone(123456)
	if err != nil {
	    panic(err)
	}
	
	if payload.Code == 200 {
	    fmt.Println("Purge succeeded")
	}


#### PurgeZones

```go
func (max *MaxCDN) PurgeZones(zones []int) (responses []GenericResponse, last error)
```
> PurgeZones purges a multiple zones caches.



##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	zones := []int{123456, 234567, 345678}
	payloads, err := max.PurgeZones(zones)
	if err != nil {
	    panic(err)
	}
	
	if len(payloads) == len(zones) {
	    fmt.Printf("Purges succeeded")
	}


#### Put

```go
func (max *MaxCDN) Put(endpoint string, form url.Values) (*GenericResponse, error)
```
> Put does an OAuth signed http.Put



##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	form := url.Values{}
	form.Set("name", "example_name")
	payload, err := max.Put("/account.json", form)
	
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)



