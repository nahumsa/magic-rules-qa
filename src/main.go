package main

import (
	"magic-rules-qa/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.Ingestion())
	rootCmd.AddCommand(cmd.Search())
	rootCmd.AddCommand(cmd.Validation())
	rootCmd.Execute()
}
