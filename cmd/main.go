package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/shanedell/gorevproxy/pkg/gorevproxy"
)

var help = "gorevproxy - CLI to spin up Golang Reverse proxy. Configured via JSON or YAML file."

var (
	configFile string
	readJSON   bool
	readYAML   bool
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gorevproxy",
		Short: help,
		Long:  help,
		RunE:  main,
	}

	cmd.PersistentFlags().StringVarP(
		&configFile,
		"config",
		"c",
		"config.yml",
		"File path to the JSON or YAML proxy config",
	)

	cmd.PersistentFlags().BoolVar(
		&readJSON,
		"json",
		false,
		"Read config file as JSON",
	)

	cmd.PersistentFlags().BoolVar(
		&readYAML,
		"yaml",
		false,
		"Read config file as YAML",
	)

	return cmd
}

func main(_ *cobra.Command, _ []string) error {
	if !readJSON && !readYAML {
		log.Println("File format to read config file as not specified, file type will be used, if that doesn't work YAML is defaulted to")

		if strings.HasSuffix(configFile, ".json") {
			readJSON = true
		} else {
			readYAML = true
		}
	}

	return gorevproxy.Run(&gorevproxy.ProxyArgs{
		ConfigFile: configFile,
		ReadJSON:   readJSON,
		ReadYAML:   readYAML,
	})
}
