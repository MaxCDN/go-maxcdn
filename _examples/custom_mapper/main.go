package main

import (
    "os"
    "fmt"
    "flag"
    "encoding/json"

    "../.."
)

var alias, token, secret string
var top int

// Generated using http://mervine.net/json2struct
// - changed float64 values to int
type PopularFiles struct {
	Code int `json:"code"`
	Data struct {
		CurrentPageSize int    `json:"current_page_size"`
		Page            int    `json:"page"`
		PageSize        string `json:"page_size"`
		Pages           int    `json:"pages"`
		Popularfiles    []struct {
			BucketID  string `json:"bucket_id"`
			Hit       string `json:"hit"`
			Size      string `json:"size"`
			Timestamp string `json:"timestamp"`
			Uri       string `json:"uri"`
			Vhost     string `json:"vhost"`
		} `json:"popularfiles"`
		Summary struct {
			Hit  string `json:"hit"`
			Size string `json:"size"`
		} `json:"summary"`
		Total string `json:"total"`
	} `json:"data"`

    // Added for extra support, see maxcdn.GenericResponse
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
    } `json:"error"`
    Raw []byte
}

// Parse turns an http response in to a GenericResponse
func (mapper *PopularFiles) Parse(raw []byte) (err error) {
	mapper.Raw = raw

	err = json.Unmarshal(raw, &mapper)
	if err != nil {
		return err
	}

	if mapper.Error.Message != "" || mapper.Error.Type != "" {
		err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

	return err
}

func init() {
    // Parse Arguments
    flag.StringVar(&alias, "alias", "", "MaxCDN consumer alias.")
    flag.StringVar(&token, "token", "", "MaxCDN consumer token.")
    flag.StringVar(&secret, "secret", "", "MaxCDN consumer secret.")
    flag.IntVar(&top, "top", 0, "Only show top N results, zero shows all.")
    flag.Parse()

    // Esure Arguments
    alias  = ensureArg(alias, "ALIAS")
    token  = ensureArg(token, "TOKEN")
    secret = ensureArg(secret, "SECRET")
}

func main() {
    max := maxcdn.NewMaxCDN(alias, token, secret)

    mapper := PopularFiles{}
    raw, err := max.Do("GET", "/reports/popularfiles.json", nil)
    if err != nil {
        panic(err)
    }

    err = mapper.Parse(raw)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%10s | %s\n", "hits", "file")
    fmt.Println("   -----------------")

    for i, file := range mapper.Data.Popularfiles {
        if top != 0 && i == top {
            break
        }
        fmt.Printf("%10s | %s\n", file.Hit, file.Uri)
    }
    fmt.Println()
}

func ensureArg(arg string, key string) string {
    if arg == "" {
        if env := os.Getenv(key); env != "" {
            return env
        }
        flag.Usage()
        os.Exit(1)
    }
    return arg
}

