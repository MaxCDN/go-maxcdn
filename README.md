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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

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
    Code  int                    `json:"code"`
    Data  map[string]interface{} `json:"data"`
    Raw   []byte                 // include raw json in GenericResponse
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
    } `json:"error"`
}
```


``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
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
    Alias string

    HTTPClient *http.Client
    // contains filtered or unexported fields
}
```


``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```
#### NewMaxCDN

```go
func NewMaxCDN(alias, token, secret string) *MaxCDN
```
> NewMaxCDN sets up a new MaxCDN instance.

``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```
#### Delete

```go
func (max *MaxCDN) Delete(endpoint string) (mapper *GenericResponse, err error)
```
> Delete does an OAuth signed http.Delete



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```

#### Do

```go
func (max *MaxCDN) Do(method, endpoint string, form url.Values) (raw []byte, err error)
```



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
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



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```

#### Post

```go
func (max *MaxCDN) Post(endpoint string, form url.Values) (mapper *GenericResponse, err error)
```
> Post does an OAuth signed http.Post



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```

#### PurgeFile

```go
func (max *MaxCDN) PurgeFile(zone int, file string) (mapper *GenericResponse, err error)
```
> PurgeFile purges a specified file by zone from cache.



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```

#### PurgeFiles

```go
func (max *MaxCDN) PurgeFiles(zone int, files []string) (responses []GenericResponse, last error)
```
> PurgeFiles purges multiple files from a zone.



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```

#### PurgeZone

```go
func (max *MaxCDN) PurgeZone(zone int) (*GenericResponse, error)
```
> PurgeZone purges a specified zones cache.



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```

#### PurgeZones

```go
func (max *MaxCDN) PurgeZones(zones []int) (responses []GenericResponse, last error)
```
> PurgeZones purges multiple zones caches.



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```

#### Put

```go
func (max *MaxCDN) Put(endpoint string, form url.Values) (mapper *GenericResponse, err error)
```
> Put does an OAuth signed http.Put



``` go
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
	raw, err := max.Do("GET", "/account.json", nil)
	
	if err != nil {
	    return
	}
	
	mapper := new(GenericResponse)
	mapper.Raw = raw // include raw json in GenericResponse
	
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
	    panic(err)
	}
	
	if mapper.Error.Message != "" || mapper.Error.Type != "" {
	    err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

```


