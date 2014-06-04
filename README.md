# go-maxcdn

MaxCDN Golang API.

## [API Documentation](http://godoc.org/github.com/jmervine/go-maxcdn)

```go
import "github.com/jmervine/go-maxcdn"
```
Package maxcdn is the golang bindings for MaxCDN's REST API.

This package should be considered beta. The final release will be moved to
`github.com/maxcdn/go-maxcdn`.

``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```
### Variables

```go
var APIHost = "https://rws.netdna.com"
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


``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```
#### MaxCDN

```go
type MaxCDN struct {
    Alias string

    HTTPClient *http.Client
    // contains filtered or unexported fields
}
```


``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```
#### NewMaxCDN

```go
func NewMaxCDN(alias, token, secret string) *MaxCDN
```
> NewMaxCDN sets up a new MaxCDN instance.

``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```
#### Delete

```go
func (max *MaxCDN) Delete(endpoint string) (*GenericResponse, error)
```
> Delete does an OAuth signed http.Delete



``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### Get

```go
func (max *MaxCDN) Get(endpoint string, form url.Values) (*GenericResponse, error)
```
> Get does an OAuth signed http.Get



``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### Post

```go
func (max *MaxCDN) Post(endpoint string, form url.Values) (*GenericResponse, error)
```
> Post does an OAuth signed http.Post



``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### PurgeFile

```go
func (max *MaxCDN) PurgeFile(zone int, file string) (*GenericResponse, error)
```
> PurgeFile purges a specified file by zone from cache.



``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### PurgeFiles

```go
func (max *MaxCDN) PurgeFiles(zone int, files []string) (responses []GenericResponse, last error)
```
> PurgeFiles purges multiple files from a zone.



``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### PurgeZone

```go
func (max *MaxCDN) PurgeZone(zone int) (*GenericResponse, error)
```
> PurgeZone purges a specified zones cache.



``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### PurgeZones

```go
func (max *MaxCDN) PurgeZones(zones []int) (responses []GenericResponse, last error)
```
> PurgeZones purges multiple zones caches.



``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```

#### Put

```go
func (max *MaxCDN) Put(endpoint string, form url.Values) (*GenericResponse, error)
```
> Put does an OAuth signed http.Put



``` go
    // Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	fmt.Printf("%#v\n", payload.Data)

```


