package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/jmervine/go-maxcdn"
	"gopkg.in/yaml.v1"
)

var config Config

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

'alias', 'token' and/or 'secret' can be set via exporting them to
your environment and ALIAS, TOKEN and/or SECRET.

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

`

	app := cli.NewApp()

	app.Name = "maxpopular"
	app.Version = "0.0.1"

	cli.HelpPrinter = helpPrinter

	app.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "~/.maxcdn.yml", "yaml file containing all required args"},
		cli.StringFlag{"alias, a", "", "[required] consumer alias"},
		cli.StringFlag{"token, t", "", "[required] consumer token"},
		cli.StringFlag{"secret, s", "", "[required] consumer secret"},
		cli.StringFlag{"host, H", "", "override default API host"},
		cli.IntFlag{"top, n", 0, "show top N results, zero shows all"},
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

		config.Top = c.Int("top")

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
	}

	app.Run(os.Args)
}

func main() {
	max := maxcdn.NewMaxCDN(config.Alias, config.Token, config.Secret)

	mapper := PopularFiles{}
	raw, err := max.Do("GET", "/reports/popularfiles.json", nil)
	check(err)

	err = mapper.Parse(raw)
	check(err)

	fmt.Printf("%10s | %s\n", "hits", "file")
	fmt.Println("   -----------------")

	for i, file := range mapper.Data.Popularfiles {
		if config.Top != 0 && i == config.Top {
			break
		}
		fmt.Printf("%10s | %s\n", file.Hit, file.Uri)
	}
	fmt.Println()
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

/*
 * Config file handlers
 */

type Config struct {
	Host   string `yaml: host,omitempty`
	Alias  string `yaml: alias,omitempty`
	Token  string `yaml: token,omitempty`
	Secret string `yaml: secret,omitempty`

	// configs not from yaml
	Top int
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
	return (c.Alias != "" && c.Token != "" && c.Secret != "")
}
