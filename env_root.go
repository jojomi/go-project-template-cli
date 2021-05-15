package main

import "github.com/spf13/cobra"

// EnvRoot encapsulates the environment for the CLI root handler.
type EnvRoot struct {
	Verbose bool
}

// ParseFrom reads the state from a given cobra command and its args.
func (e *EnvRoot) ParseFrom(command *cobra.Command, args []string) error {
	var err error

	e.Verbose, err = command.Flags().GetBool("verbose")
	if err != nil {
		return err
	}
	return nil
}
