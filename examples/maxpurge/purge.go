package main

import (
    "os"
    "fmt"
    "flag"
    "time"
    "strconv"

    "github.com/jmervine/go-maxcdn"
)

var alias, token, secret, zone, file string
var start time.Time

func init() {
    start = time.Now()

    // Parse Arguments
    flag.StringVar(&file, "file", "", "MaxCDN cached file to be purged (empty purges all).")
    flag.StringVar(&zone, "zone", "", "MaxCDN zone to be purged.")
    flag.StringVar(&alias, "alias", "", "MaxCDN consumer alias.")
    flag.StringVar(&token, "token", "", "MaxCDN consumer token.")
    flag.StringVar(&secret, "secret", "", "MaxCDN consumer secret.")
    flag.Parse()

    // Esure Arguments
    zone   = ensureArg(zone, "ZONE")
    alias  = ensureArg(alias, "ALIAS")
    token  = ensureArg(token, "TOKEN")
    secret = ensureArg(secret, "SECRET")
}

func main() {
    max := maxcdn.NewMaxCDN(alias, token, secret)

    i, err := strconv.ParseInt(zone, 0, 64)
    check(err)

    zoneid := int(i)

    var response *maxcdn.GenericResponse
    if file != "" {
        response, err = max.PurgeFile(zoneid, file)
    } else {
        response, err = max.PurgeZone(zoneid)
    }
    check(err)

    if response.Code == 200 {
        fmt.Printf("Purge successful after: %v.\n", time.Since(start))
    }
}

func ensureArg(arg string, key string) string {
    if arg == "" {
        if env := os.Getenv(key); env != "" {
            return env
        }
        flag.Usage()
        os.Exit(0)
    }
    return arg
}

func check(err error) {
    if err != nil {
        fmt.Printf("%v.\n\nPurge failed after %v.\n", err, time.Since(start))
        os.Exit(2)
    }
}
