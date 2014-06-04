purge example
=============

This example shows usage of 'maxcdn.PurgeZone' and 'maxcdn.PurgeFile'.

Trying this example:
--------------------

```
$ cd _example/purge
$ go build -o purge

$ ./purge -h
Usage of ./purge:
  -alias="": MaxCDN consumer alias.
  -file="": MaxCDN cached file to be purged (empty purges all).
  -secret="": MaxCDN consumer secret.
  -token="": MaxCDN consumer token.
  -zone="": MaxCDN zone to be purged.

$ ./purge -alias YOUR_ALIAS -token YOUR_TOKEN -secret YOUR_SECRET \
    -zone 123456
Purge successful after: 2.078010673s.

$ ./purge -alias YOUR_ALIAS -token YOUR_TOKEN -secret YOUR_SECRET \
    -zone 123456 -file "/master.css"
Purge successful after: 1.078010673s.
```

