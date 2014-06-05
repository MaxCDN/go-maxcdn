maxpurge example
================

This example shows usage of the `PurgeZone` and `PurgeFile` methods.

Trying this example:
--------------------

```
$ cd examples/maxpurge
$ go build -o maxpurge

$ ./maxpurge -h
Usage: maxpurge [arguments...]
Options:
   --alias, -a          [required] consumer alias
   --token, -t          [required] consumer token
   --secret, -s         [required] consumer secret
   --zone, -z           [required] zone to be purged
   --file, -f           cached file to be purged
   --version, -v        print the version
   --help, -h           show help

$ ./maxpurge -alias YOUR_ALIAS -token YOUR_TOKEN -secret YOUR_SECRET \
    -zone 123456
Purge successful after: 2.078010673s.

$ ./maxpurge -alias YOUR_ALIAS -token YOUR_TOKEN -secret YOUR_SECRET \
    -zone 123456 -file "/master.css"
Purge successful after: 1.078010673s.
```

Installing this example:
------------------------

This can also be installed for system wide use if your `GOBIN` is set via the following:

```bash
$ go get github.com/jmervine/go-maxcdn
$ go install github.com/jmervine/go-maxcdn/examples/maxpurge
$ maxpurge -h
Usage: maxpurge [arguments...]
Options:
   --alias, -a          [required] consumer alias
   --token, -t          [required] consumer token
   --secret, -s         [required] consumer secret
   --zone, -z           [required] zone to be purged
   --file, -f           cached file to be purged
   --version, -v        print the version
   --help, -h           show help
```

