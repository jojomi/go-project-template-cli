package main

import (
	"github.com/iancoleman/strcase"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func getRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: strcase.ToKebab(ToolName),
		Run: handleRootCmd,
	}

	pf := cmd.PersistentFlags()
	pf.BoolP("verbose", "v", false, "activate verbose output")

	// configure command local flags
	// f := cmd.Flags()
	// f.BoolP("dry-run", "d", false, "Prevent destructive operations")

	// add subcommands
	// cmd.AddCommand(getOtherCmd())

	return cmd
}

func handleRootCmd(cmd *cobra.Command, args []string) {
	env := EnvRoot{}
	err := env.ParseFrom(cmd, args)
	if err != nil {
		log.Fatal().Err(err).Msg("could not parse command")
	}
	handleRoot(env)
}

func handleRoot(env EnvRoot) {
	setLoggerVerbosity(env.Verbose)

	// TODO implement logic here

	// read config
	// conf := mustGetConfig(configFilename)
	// fmt.Println(conf)
}
