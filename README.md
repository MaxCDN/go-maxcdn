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
	
	if payload.Error.Message != "" {
	    panic(errors.New(fmt.Sprintf("%s %s", payload.Error.Type, payload.Error.Message)))
	}
	
	fmt.Printf("%#v\n", payload.Data)

### Constants

```go
const (
    ApiPath     = "https://rws.netdna.com"
    UserAgent   = "Go MaxCDN API Client"
    ContentType = "application/x-www-form-urlencoded"
)
```



### Types

#### MaxCDN

```go
type MaxCDN struct {
    Alias string

    HttpClient *http.Client
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
func (max *MaxCDN) Delete(endpoint string) (*Response, error)
```



##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	// This specific example shows how to purge a cache without using the Purge
	// methods, more as an example of using Delete, then anything, really.
	
	payload, err := max.Delete(fmt.Sprintf("/zones/pull.json/%d/cache", 123456))
	if err != nil {
	    panic(err)
	}
	
	if payload.Error.Message != "" {
	    panic(errors.New(fmt.Sprintf("%s %s", payload.Error.Type, payload.Error.Message)))
	}
	
	if payload.Code == 200 {
	    fmt.Println("Purge succeeded")
	}


#### Get

```go
func (max *MaxCDN) Get(endpoint string, form url.Values) (*Response, error)
```



##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	if payload.Error.Message != "" {
	    panic(errors.New(fmt.Sprintf("%s %s", payload.Error.Type, payload.Error.Message)))
	}
	
	fmt.Printf("%#v\n", payload.Data)


#### Post

```go
func (max *MaxCDN) Post(endpoint string, form url.Values) (*Response, error)
```



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
	
	if payload.Error.Message != "" {
	    panic(errors.New(fmt.Sprintf("%s %s", payload.Error.Type, payload.Error.Message)))
	}
	
	data := payload.Data["pullzone"].(map[string]interface{})
	if data["name"] == "newzone" {
	    fmt.Println("Successfully created new Pull Zone.")
	}


#### Put

```go
func (max *MaxCDN) Put(endpoint string, form url.Values) (*Response, error)
```



##### Example:
	max := NewMaxCDN(os.Getenv("ALIAS"), os.Getenv("TOKEN"), os.Getenv("SECRET"))
	
	form := url.Values{}
	form.Set("name", "example_name")
	payload, err := max.Put("/account.json", form)
	
	if err != nil {
	    panic(err)
	}
	
	if payload.Error.Message != "" {
	    panic(errors.New(fmt.Sprintf("%s %s", payload.Error.Type, payload.Error.Message)))
	}
	
	fmt.Printf("%#v\n", payload.Data)


#### Response

```go
type Response struct {
    Code  float64                `json:"code"`
    Data  map[string]interface{} `json:"data"`
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
    } `json:"error"`
}
```




