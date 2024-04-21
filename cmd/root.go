package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stl3mf",
	Short: "STL to 3MF conversion tooling",
	Long:  `...`,
}

// Execute a subcommand in the cmd directory
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
