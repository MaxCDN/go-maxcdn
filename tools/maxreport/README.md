maxreport
=========

Runs reports against [MaxCDN's Reports API](http://docs.maxcdn.com/#reports-api).

Installing:
-----------

This can also be installed for system wide use if your `GOBIN` is set via the following:

```bash
$ go get github.com/jmervine/go-maxcdn
$ go install github.com/jmervine/go-maxcdn/tools/maxreport
$ maxreport -h
Usage: maxreport [arguments...]
# ...

```

Building:
---------

```
$ cd tools/maxreport
$ go build

$ ./maxreport -h
Usage: maxreport [arguments...]
# ...

$ ./maxreport -a ALIAS -t TOKEN -s SECRET popular -t 5
Running popular files report.

      hits | file
   -----------------
      6295 | /bootstrap/favicon.ico
      5624 | /bootstrap/master.css
       199 | /
       133 | /bootstrap/images/apple-touch-icon-114x114.png
       107 | /bootstrap/images/apple-touch-icon.png

```
