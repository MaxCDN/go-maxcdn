package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"text/template"
	"time"

	"github.com/codegangsta/cli"
	"github.com/jmervine/go-maxcdn"
	"gopkg.in/yaml.v1"
)

var start time.Time
var config Config

func init() {

	// Override cli's default help template
	cli.AppHelpTemplate = `Usage: {{.Name}} [arguments...]
Options:
   {{range .Flags}}{{.}}
   {{end}}

'alias', 'token' 'secret' and/or 'zone' can be set via exporting them
to your environment and ALIAS, TOKEN, SECRET and/or ZONE.

Additionally, they can be set in a YAML configuration via the
config option. 'host' can also be set via configuration, but not
environment.

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
    zone: YOUR_ZONE_ID

`

	app := cli.NewApp()

	app.Name = "maxpurge"
	app.Version = "0.0.1"

	cli.HelpPrinter = helpPrinter

	app.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "~/.maxcdn.yml", "yaml file containing all required args"},
		cli.StringFlag{"alias, a", "", "[required] consumer alias"},
		cli.StringFlag{"token, t", "", "[required] consumer token"},
		cli.StringFlag{"secret, s", "", "[required] consumer secret"},
		cli.StringFlag{"zone, z", "", "[required] zone to be purged"},
		cli.StringFlag{"file, f", "", "cached file to be purged"},
		cli.StringFlag{"host, H", "", "override default API host"},
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

		if v := c.String("zone"); v != "" {
			config.Zone = v
		} else if v := os.Getenv("ZONE"); v != "" {
			config.Zone = v
		}

		config.File = c.String("file")

		if !config.Validate() {
			fmt.Printf("%+v\n", config)
			cli.ShowAppHelp(c)
		}

		if v := c.String("host"); v != "" {
			config.Host = v
		}
		// handle host override
		if config.Host != "" {
			maxcdn.APIHost = config.Host
		}
	}

	app.Run(os.Args)

	start = time.Now()
}

func main() {
	max := maxcdn.NewMaxCDN(config.Alias, config.Token, config.Secret)

	i, err := strconv.ParseInt(config.Zone, 0, 64)
	check(err)

	zoneid := int(i)

	var response *maxcdn.GenericResponse
	if config.File != "" {
		response, err = max.PurgeFile(zoneid, config.File)
	} else {
		response, err = max.PurgeZone(zoneid)
	}
	check(err)

	if response.Code == 200 {
		fmt.Printf("Purge successful after: %v.\n", time.Since(start))
	}
}

func check(err error) {
	if err != nil {
		fmt.Printf("%v.\n\nPurge failed after %v.\n", err, time.Since(start))
		os.Exit(2)
	}
}

// Replace cli's default help printer with cli's default help printer
// plus an exit at the end.
func helpPrinter(templ string, data interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	t := template.Must(template.New("help").Parse(templ))
	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}
	w.Flush()
	os.Exit(0)
}

/*
 * Config file handlers
 */

type Config struct {
	Host   string `yaml: host,omitempty`
	Alias  string `yaml: alias,omitempty`
	Token  string `yaml: token,omitempty`
	Secret string `yaml: secret,omitempty`
	Zone   string `yaml: secret,omitempty`

	// configs not from yaml
	File string
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
	return (c.Alias != "" && c.Token != "" && c.Secret != "" && c.Zone != "")
}
