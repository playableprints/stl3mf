package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VERSION string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run:   version,
}

func version(cmd *cobra.Command, args []string) {
	fmt.Println("stl3mf by Playable Prints. See more at https://www.playableprints.com")
	fmt.Println(VERSION)
}
