maxcdn tools
============

All tools use a configuration file as it's last means of getting `alias`, `secret` and
`token`. See individal tool `help` for addtional options available in your configuration.

```yaml
---
alias: YOUR_ALIAS
token: YOUR_TOKEN
secret: YOUR_SECRET
```

See [sample.maxcdn.yml](sample.maxcdn.yml) for a more complete example.


Installing:
-----------

This can also be installed for system wide use if your `GOBIN` is set via the following:

```bash
# via 'go get' && 'go install'
##

go get github.com/jmervine/go-maxcdn
go install github.com/jmervine/go-maxcdn/tools/{{tool}}

# manually
##

git clone https://github.com/jmervine/go-maxcdn
cd go-maxcdn/tools

# build and install all tools
make build install

# or a single tool
make build/{{tool}} install/{{tool}}

# windows users
cd {{tool}}
go build
move {{tool}} c:\path\to\GOBIN
```

See individal tool README for additional instructions.

Prebuilt Binaries:
------------------

A set of binaries for all tools have been prebuilt using golang's cross compiler on `Linux 3.8.0-36-generic #52~precise1-Ubuntu SMP x86_64`.

**Building All Binaries:**

```bash
make setup # only once, this takes a while
make build/all

# or
make build/all/{{tool}}
```

Here's what's available for each tool:

- maxreport
    - [linux-386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/linux/386/maxpurge)
    - [linux-amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/linux/amd64/maxpurge)
    - [linux-arm](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/linux/arm/maxpurge)
    - [darwin-386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/darwin/386/maxpurge)
    - [darwin-amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/darwin/amd64/maxpurge)
    - [freebsd-386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/freebsd/386/maxpurge)
    - [freebsd-amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/freebsd/amd64/maxpurge)
    - [freebsd-arm](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/freebsd/arm/maxpurge)
    - [windows-386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/windows/386/maxpurge.exe)
    - [windows-amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxreport/builds/windows/amd64/maxpurge.exe)
- maxpurge
    - [linux/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/linux/386/maxpurge)
    - [linux/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/linux/amd64/maxpurge)
    - [linux/arm](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/linux/arm/maxpurge)
    - [darwin/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/darwin/386/maxpurge)
    - [darwin/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/darwin/amd64/maxpurge)
    - [freebsd/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/freebsd/386/maxpurge)
    - [freebsd/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/freebsd/amd64/maxpurge)
    - [freebsd/arm](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/freebsd/arm/maxpurge)
    - [windows/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/windows/386/maxpurge.exe)
    - [windows/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxpurge/builds/windows/amd64/maxpurge.exe)
- maxcurl
    - [linux/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/linux/386/maxcurl)
    - [linux/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/linux/amd64/maxcurl)
    - [linux/arm](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/linux/arm/maxcurl)
    - [darwin/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/darwin/386/maxcurl)
    - [darwin/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/darwin/amd64/maxcurl)
    - [freebsd/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/freebsd/386/maxcurl)
    - [freebsd/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/freebsd/amd64/maxcurl)
    - [freebsd/arm](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/freebsd/arm/maxcurl)
    - [windows/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/windows/386/maxcurl.exe)
    - [windows/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxcurl/builds/windows/amd64/maxcurl.exe)
- maxtail
    - [linux/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/linux/386/maxtail)
    - [linux/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/linux/amd64/maxtail)
    - [linux/arm](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/linux/arm/maxtail)
    - [darwin/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/darwin/386/maxtail)
    - [darwin/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/darwin/amd64/maxtail)
    - [freebsd/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/freebsd/386/maxtail)
    - [freebsd/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/freebsd/amd64/maxtail)
    - [freebsd/arm](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/freebsd/arm/maxtail)
    - [windows/386](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/windows/386/maxtail.exe)
    - [windows/amd64](https://github.com/jmervine/go-maxcdn/raw/master/tools/maxtail/builds/windows/amd64/maxtail.exe)

> Note: As of yet, these binaries have not been tested on all OS/ARCH combinations.

