package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

var (
	configFilename = "config.yml"
)

func main() {
	setupLogger()

	conf := mustGetConfig(configFilename)
	fmt.Println(conf)
	os.Exit(1)

	c := mustGetConfig(configFilename)
	fmt.Println(c)

	// build root command
	rootCmd := getRootCmd()

	// add version command
	rootCmd.AddCommand(getVersionCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
