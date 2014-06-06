maxcurl
=======

"curl" (sort of) MaxCDN endpoints, return raw json output.

Installing:
-----------

This can also be installed for system wide use if your `GOBIN` is set via the following:

```bash
$ go get github.com/jmervine/go-maxcdn
$ go install github.com/jmervine/go-maxcdn/tools/maxcurl
$ maxcurl -h
Usage: maxcurl [arguments...]
# ...
```

Building:
---------

```
$ cd tools/maxcurl
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

