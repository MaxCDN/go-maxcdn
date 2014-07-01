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
Usage: maxcurl [arguments...] PATH
# ...
```

Building:
---------

```
$ cd tools/maxcurl
$ go build

$ ./maxcurl -h
Usage: maxcurl [arguments...] PATH
# ...

$ ./maxcurl -a ALIAS -t TOKEN -s SECRET -pp /zones/pull.json/count
{
  "code": 200,
  "data": {
    "count": "214"
  }
}
```

