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
$ go get github.com/jmervine/go-maxcdn
$ go install github.com/jmervine/go-maxcdn/tools/maxpurge
$ maxpurge -h
Usage: maxpurge [arguments...]
# ...
```

Building:
---------

```
$ cd tools/maxpurge
$ go build -o maxpurge

$ ./maxpurge -h
Usage: maxpurge [arguments...]
# ...

$ ./maxpurge -alias ALIAS -token TOKEN -secret SECRET \
    -zone 123456 -zone 234567
Purge successful after: 2.078010673s.

$ ./maxpurge -alias ALIAS -token TOKEN -secret SECRET \
    -zone 123456 -file "/master.css" -file "/master.js"
Purge successful after: 1.078010673s.
```
