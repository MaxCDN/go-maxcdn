maxcurl example
===============

This example, beyond being a very useful tool in itself, shows the most basic `Do` usage.

Trying this example:
--------------------

```
$ cd examples/maxcurl
$ go build

$ ./maxcurl -h
Usage: maxcurl [arguments...]
Options:
   --alias, -a          [required] consumer alias
   --token, -t          [required] consumer token
   --secret, -s         [required] consumer secret
   --path, -p           [required] request path, e.g. /account.json
   --method, -X 'GET'   request method
   --pretty, --pp       pretty print json output
   --version, -v        print the version
   --help, -h           show help

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
Options:
   --alias, -a          [required] consumer alias
   --token, -t          [required] consumer token
   --secret, -s         [required] consumer secret
   --path, -p           [required] request path, e.g. /account.json
   --method, -X 'GET'   request method
   --pretty, --pp       pretty print json output
   --version, -v        print the version
   --help, -h           show help
```

