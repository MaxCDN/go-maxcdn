package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/jmervine/go-maxcdn"
)

var alias, token, secret, method, path string
var pretty, help bool

func init() {

	// Override cli's default help template
	cli.AppHelpTemplate = `Usage: {{.Name}} [arguments...]
Options:
   {{range .Flags}}{{.}}
   {{end}}
`

	app := cli.NewApp()

	app.Name = "maxcurl"
	app.Version = "0.0.1"

	cli.HelpPrinter = helpPrinter

	app.Flags = []cli.Flag{
		cli.StringFlag{"alias, a", "", "[required] consumer alias"},
		cli.StringFlag{"token, t", "", "[required] consumer token"},
		cli.StringFlag{"secret, s", "", "[required] consumer secret"},
		cli.StringFlag{"path, p", "", "[required] request path, e.g. /account.json"},

		cli.StringFlag{"method, X", "GET", "request method"},
		cli.BoolFlag{"pretty, pp", "pretty print json output"},
	}

	app.Action = func(c *cli.Context) {
		alias = ensureArg(c.String("alias"), "ALIAS", c)
		token = ensureArg(c.String("token"), "TOKEN", c)
		secret = ensureArg(c.String("secret"), "SECRET", c)
		method = c.String("method")
		path = c.String("path")
		pretty = c.Bool("pretty")

		if path == "" {
			cli.ShowAppHelp(c)
		}
	}

	app.Run(os.Args)
}

func main() {
	max := maxcdn.NewMaxCDN(alias, token, secret)

	u, err := url.Parse(path)
	check(err)

	path = u.Path
	form := u.Query()

	raw, err := max.Do(method, path, form)
	check(err)

	if pretty {
		var j interface{}
		err = json.Unmarshal(raw, &j)
		check(err)

		raw, err = json.MarshalIndent(j, "", "  ")
		check(err)
	}

	fmt.Printf("%v\n", string(raw))
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
