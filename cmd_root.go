package main

import (
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

func getRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: strcase.ToKebab(ToolName),
	}
	return cmd
}

func handleRootCmd(cmd *cobra.Command, args []string) {

}
