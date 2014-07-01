maxpurge
========

This provides a simple interface for purging pull zones and their files.

> TODO:
>
> - support a list of zones
> - support no zone, which purges all zones
> - support a list of files

Installing:
-----------

This can also be installed for system wide use if your `GOBIN` is set via the following:

```bash
# via 'go get' && 'go install'
##

$ go get github.com/jmervine/go-maxcdn
$ go install github.com/jmervine/go-maxcdn/tools/maxpurge
$ maxpurge -h
Usage: maxpurge [arguments...] PATH
# ...

# manually
##

git clone https://github.com/jmervine/go-maxcdn
cd go-maxcdn/tools
make build/maxpurge install/maxpurge
```
