maxcurl example
===============

This example, beyond being a very useful tool in itself, shows the most basic `Do` usage.

> All examples may seem to a be a bit overly complex, but that's because (a) the documentation
> has examples for pretty much everything and (b) these are intended to be fully functional
> and useful tools for interacting with MaxCDN's API.

Trying this example:
--------------------

```
$ cd examples/maxcurl
$ go build

$ ./maxcurl -h
Usage: maxcurl [arguments...]
# ...

$ ./maxcurl -p /zones/pull.json/count -pp
{
  "code": 200,
  "data": {
    "count": "214"
  }
}
```

Installing this example:
------------------------

This can also be installed for system wide use if your `GOBIN` is set via the following:

```bash
$ go get github.com/jmervine/go-maxcdn
$ go install github.com/jmervine/go-maxcdn/examples/maxcurl
$ maxcurl -h
Usage: maxcurl [arguments...]
# ...
```

