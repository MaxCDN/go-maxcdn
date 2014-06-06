package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/jmervine/go-maxcdn"
	"gopkg.in/yaml.v1"
)

var config Config

func init() {

	// Override cli's default help template
	cli.AppHelpTemplate = `Usage: {{.Name}} [arguments...] PATH

Example:

    $ {{.Name}} -a ALIAS -t TOKEN -s SECRET /account.json

Options:

   {{range .Flags}}{{.}}
   {{end}}


'alias', 'token' and/or 'secret' can be set via exporting them to
your environment and ALIAS, TOKEN and/or SECRET.

Additionally, they can be set in a YAML configuration via the
config option. 'pretty' and 'host' can also be set via
configuration, but not environment.

Precedence is argument > environment > configuration.

WARNING:
    Default configuration path works for *nix systems only and
    replies on the 'HOME' environment variable. For Windows, please
    supply a full path.

Sample configuration:

    ---
    alias: YOUR_ALIAS
    token: YOUR_TOKEN
    secret: YOUR_SECRET
    pretty: true

`

	app := cli.NewApp()
	app.Name = "maxcurl"
	app.Version = "0.0.1"

	cli.HelpPrinter = helpPrinter

	app.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "~/.maxcdn.yml", "yaml file containing all required args"},
		cli.StringFlag{"alias, a", "", "[required] consumer alias"},
		cli.StringFlag{"token, t", "", "[required] consumer token"},
		cli.StringFlag{"secret, s", "", "[required] consumer secret"},
		cli.StringFlag{"method, X", "GET", "request method"},
		cli.StringFlag{"host, H", "", "override default API host"},
		cli.BoolFlag{"headers, i", "show headers with body"},
		cli.BoolFlag{"pretty, pp", "pretty print json output"},
	}

	app.Action = func(c *cli.Context) {
		// Precedence
		// 1. CLI Argument
		// 2. Environment (when applicable)
		// 3. Configuration

		config, _ = LoadConfig(c.String("config"))

		if v := c.String("alias"); v != "" {
			config.Alias = v
		} else if v := os.Getenv("ALIAS"); v != "" {
			config.Alias = v
		}

		if v := c.String("token"); v != "" {
			config.Token = v
		} else if v := os.Getenv("TOKEN"); v != "" {
			config.Token = v
		}

		if v := c.String("secret"); v != "" {
			config.Secret = v
		} else if v := os.Getenv("SECRET"); v != "" {
			config.Secret = v
		}

		config.Method = c.String("method")
		config.Headers = c.Bool("headers")
		config.Path = c.Args().First()

		if !config.Validate() {
			cli.ShowAppHelp(c)
		}

		if v := c.String("host"); v != "" {
			config.Host = v
		}
		// handle host override
		if config.Host != "" {
			maxcdn.APIHost = config.Host
		}

		if v := c.Bool("pretty"); v {
			config.Pretty = v
		}
	}

	app.Run(os.Args)
}

func main() {
	max := maxcdn.NewMaxCDN(config.Alias, config.Token, config.Secret)

	// seperate path and query
	u, err := url.Parse(config.Path)
	check(err)

	config.Path = u.Path
	form := u.Query()

	// request raw data from maxcdn
	raw, res, err := max.Do(config.Method, config.Path, form)
	check(err)

	if config.Pretty {
		// format pretty
		var j interface{}
		err = json.Unmarshal(raw, &j)
		check(err)

		raw, err = json.MarshalIndent(j, "", "  ")
		check(err)
	}

	// print
	if config.Headers {
		fmt.Println(fmtHeaders(res.Header))
	}
	fmt.Printf("%v\n", string(raw))
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

func fmtHeaders(headers map[string][]string) (out string) {
	for k, v := range headers {
		out += fmt.Sprintf("%s => %s\n", k, strings.Join(v, ", "))
	}
	return
}

/*
 * Config file handlers
 */

type Config struct {
	Host   string `yaml: host,omitempty`
	Alias  string `yaml: alias,omitempty`
	Token  string `yaml: token,omitempty`
	Secret string `yaml: secret,omitempty`
	Pretty bool   `yaml: pretty,omitempty`

	// configs not from yaml
	Method, Path string
	Headers      bool
}

func LoadConfig(file string) (c Config, e error) {
	// TODO: this is unix only, look at fixing for windows
	file = strings.Replace(file, "~", os.Getenv("HOME"), 1)

	c = Config{}
	if data, err := ioutil.ReadFile(file); err == nil {
		e = yaml.Unmarshal(data, &c)
	}

	return
}

func (c *Config) Validate() bool {
	return (c.Alias != "" && c.Token != "" && c.Secret != "" && c.Path != "")
}
