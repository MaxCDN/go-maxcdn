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
	// This, like all examples are meant to be functional integration tests
	// as well as examples for documentation. To run these as integration tests
	// export your ALIAS, TOKEN and SECRET to your envioronment before running
	// 'go test', otherwise the http request will be stubbed using the json files
	// in './_fixtures/*.json'.
	
	handleErrors := func(r *Response, e error) {
	    if e != nil {
	        panic(e)
	    }
	
	    if r.Error.Message != "" {
	        panic(errors.New(fmt.Sprintf("%s %s", r.Error.Type, r.Error.Message)))
	    }
	}
	
	max := NewMaxCDN(alias, token, secret)
	
	if alias == "" || token == "" || secret == "" {
	    max.HttpClient = stubHttpOk()
	}
	
	payload, err := max.Get("/account.json", nil)
	handleErrors(payload, err)
	
	data := payload.Data["account"].(map[string]interface{})
	if data["name"] != "" {
	    fmt.Println("GET /account.json succeeded")
	}
	
	// Output: GET /account.json succeeded

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
	// This, like all examples are meant to be functional integration tests
	// as well as examples for documentation. To run these as integration tests
	// export your ALIAS, TOKEN and SECRET to your envioronment before running
	// 'go test', otherwise the http request will be stubbed using the json files
	// in './_fixtures/*.json'.
	
	// This specific example shows how to purge a cache without using the Purge
	// methods, more as an example of using Delete, then anything, really.
	
	handleErrors := func(r *Response, e error) {
	    if e != nil {
	        panic(e)
	    }
	
	    if r.Error.Message != "" {
	        panic(errors.New(fmt.Sprintf("%s %s", r.Error.Type, r.Error.Message)))
	    }
	}
	
	max := NewMaxCDN(alias, token, secret)
	
	if alias == "" || token == "" || secret == "" {
	    max.HttpClient = stubHttpOk()
	}
	
	// Start by fetching the first zone, as that's the one we'll be purging.
	payload, err := max.Get("/zones/pull.json", nil)
	handleErrors(payload, err)
	
	data := payload.Data["pullzones"].([]interface{})
	zone := data[0].(map[string]interface{})
	id := zone["id"].(string)
	
	// Now purge that zone's cache.
	payload, err = max.Delete(fmt.Sprintf("/zones/pull.json/%s/cache", id))
	handleErrors(payload, err)
	
	if payload.Code == 200 {
	    fmt.Println("Purge succeeded")
	}
	
	// Output: Purge succeeded


#### Get

```go
func (max *MaxCDN) Get(endpoint string, form url.Values) (*Response, error)
```



##### Example:
	max := NewMaxCDN(alias, token, secret)
	payload, err := max.Get("/account.json", nil)
	if err != nil {
	    panic(err)
	}
	
	data := payload.Data["account"].(map[string]interface{})
	fmt.Println("%#v\n", data)


#### Post

```go
func (max *MaxCDN) Post(endpoint string, form url.Values) (*Response, error)
```



##### Example:
	// This, like all examples are meant to be functional integration tests
	// as well as examples for documentation. To run these as integration tests
	// export your ALIAS, TOKEN and SECRET to your envioronment before running
	// 'go test', otherwise the http request will be stubbed using the json files
	// in './_fixtures/*.json'.
	
	handleErrors := func(r *Response, e error) {
	    if e != nil {
	        panic(e)
	    }
	
	    if r.Error.Message != "" {
	        panic(errors.New(fmt.Sprintf("%s %s", r.Error.Type, r.Error.Message)))
	    }
	}
	
	// I'm using a timestamp as unique name for testing, you shouldn't
	// do this.
	name := stringFromTimestamp()
	
	max := NewMaxCDN(alias, token, secret)
	
	if alias == "" || token == "" || secret == "" {
	    max.HttpClient = stubHttpOk()
	    name = "newpullzone3"
	}
	
	form := url.Values{}
	form.Set("name", name)
	
	// When creating a new zone, the url must be real and resolve.
	form.Set("url", "http://www.example.com")
	
	payload, err := max.Post("/zones/pull.json", form)
	handleErrors(payload, err)
	
	data := payload.Data["pullzone"].(map[string]interface{})
	if data["name"] == name {
	    fmt.Println("POST /zones/pull.json succeeded")
	
	    id := int(data["id"].(float64))
	
	    payload, err = max.Delete(fmt.Sprintf("/zones/pull.json/%d", id))
	    handleErrors(payload, err)
	
	    if payload.Code == 200 {
	        fmt.Println("DELETE /zones/pull.json succeeded")
	    }
	}
	
	// Output:
	// POST /zones/pull.json succeeded
	// DELETE /zones/pull.json succeeded


#### Put

```go
func (max *MaxCDN) Put(endpoint string, form url.Values) (*Response, error)
```



##### Example:
	// This, like all examples are meant to be functional integration tests
	// as well as examples for documentation. To run these as integration tests
	// export your ALIAS, TOKEN and SECRET to your envioronment before running
	// 'go test', otherwise the http request will be stubbed using the json files
	// in './_fixtures/*.json'.
	
	// I'm using a timestamp as unique name for testing, you shouldn't
	// do this.
	name := stringFromTimestamp()
	
	max := NewMaxCDN(alias, token, secret)
	
	if alias == "" || token == "" || secret == "" {
	    max.HttpClient = stubHttpOk()
	    name = "MaxCDN sampleCode"
	}
	
	form := url.Values{}
	form.Set("name", name)
	payload, err := max.Put("/account.json", form)
	
	if err != nil {
	    panic(err)
	}
	
	data := payload.Data["account"].(map[string]interface{})
	if data["name"] == name {
	    fmt.Println("PUT /account.json succeeded")
	}
	
	// Output: PUT /account.json succeeded


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




