package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/jmervine/go-maxcdn"
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

	// Override cli's default help template
	cli.AppHelpTemplate = `Usage: {{.Name}} [arguments...]
Options:
   {{range .Flags}}{{.}}
   {{end}}
`

	app := cli.NewApp()

	app.Name = "maxpopular"
	app.Version = "0.0.1"

	cli.HelpPrinter = helpPrinter

	app.Flags = []cli.Flag{
		cli.StringFlag{"alias, a", "", "[required] consumer alias"},
		cli.StringFlag{"token, t", "", "[required] consumer token"},
		cli.StringFlag{"secret, s", "", "[required] consumer secret"},
		cli.IntFlag{"top, n", 0, "show top N results, zero shows all"},
	}

	app.Action = func(c *cli.Context) {
		alias = ensureArg(c.String("alias"), "ALIAS", c)
		token = ensureArg(c.String("token"), "TOKEN", c)
		secret = ensureArg(c.String("secret"), "SECRET", c)
		top = c.Int("top")
	}

	app.Run(os.Args)
}

func main() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	mapper := PopularFiles{}
	raw, err := max.Do("GET", "/reports/popularfiles.json", nil)
	check(err)

	err = mapper.Parse(raw)
	check(err)

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

func ensureArg(arg string, key string, c *cli.Context) string {
	if arg == "" {
		if env := os.Getenv(key); env != "" {
			return env
		}
		cli.ShowAppHelp(c)
	}
	return arg
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Replace cli's default help printer with cli's default help printer
// plus an exit at the end.
func helpPrinter(templ string, data interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	t := template.Must(template.New("help").Parse(templ))
	err := t.Execute(w, data)
	check(err)

	w.Flush()
	os.Exit(0)
}
