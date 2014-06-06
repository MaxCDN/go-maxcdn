maxpopular
==========

Return a sorted list of top requested files, based on the MaxCDN "/reports/popularfiles.json" endpoint.

Custom Mapping:
---------------

This tool provides an example of custom mapping, which requires that you parse the JSON response yourself. The recommended method of doing this is to build a struct to match the JSON response. In addition to the JSON fields, it's a good idea to also add an `Error` and `Raw` mapping for error handling and debugging. See `type PopularFiles struct { ... }` in the example.

> JSON response examples can be found in [MaxCDN's Documentation](http://docs.maxcdn.com/) and [mervine.net/json2struct](http://mervine.net/json2struct) can be used (and was in this) as a starting point for your struct, but modification may be required.


Installing:
-----------

This can also be installed for system wide use if your `GOBIN` is set via the following:

```bash
$ go get github.com/jmervine/go-maxcdn
$ go install github.com/jmervine/go-maxcdn/tools/maxpopular
$ maxpopular -h
Usage: maxpopular [arguments...]
# ...

```

Building:
---------

```
$ cd tools/maxpopular
$ go build

$ ./maxpopular -h
Usage: maxpopular [arguments...]
# ...

$ ./maxpopular -alias ALIAS -token TOKEN -secret SECRET -top 5
      hits | file
   -----------------
      6295 | /bootstrap/favicon.ico
      5624 | /bootstrap/master.css
       199 | /
       133 | /bootstrap/images/apple-touch-icon-114x114.png
       107 | /bootstrap/images/apple-touch-icon.png

```
