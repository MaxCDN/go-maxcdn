custom mapper example
=====================

This example shows two things, first and primarily how to set up a custom mapper, should you need one. Second, a basic `Get`.


Custom Mapping:
---------------

Custom mapping requires that you parse the JSON response yourself. The recommended method of doing this is to build a struct to match the JSON response. In addition to the JSON fields, it's a good idea to also add an `Error` and `Raw` mapping for error handling and debugging. See `type PopularFiles struct { ... }` in the example.

> JSON response examples can be found in [MaxCDN's Documentation](http://docs.maxcdn.com/) and [mervine.net/json2struct](http://mervine.net/json2struct) can be used (and was in this example) as a starting point for your struct, but modification may be required.


Trying this example:
--------------------

```
$ cd _example/custom_mapper
$ go build -o popular

$ ./purge -h
Usage of ./popular:
  -alias="": MaxCDN consumer alias.
  -secret="": MaxCDN consumer secret.
  -token="": MaxCDN consumer token.
  -top=0: Only show top N results, zero shows all.

$ ./popular -alias YOUR_ALIAS -token YOUR_TOKEN -secret YOUR_SECRET -top 5
      hits | file
   -----------------
      6295 | /bootstrap/favicon.ico
      5624 | /bootstrap/master.css
       199 | /
       133 | /bootstrap/images/apple-touch-icon-114x114.png
       107 | /bootstrap/images/apple-touch-icon.png

```

