package cmd

import (
	"fmt"
	"stl3mf/models"
	"strings"

	"github.com/spf13/cobra"
)

var (
	checkFilepath string
)

func init() {

	checkCmd.Flags().StringVarP(&checkFilepath, "file", "f", "", "STL filepath")

	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check an STL for errors",
	RunE:  check,
}

func check(cmd *cobra.Command, args []string) error {
	checkFilepath = strings.ReplaceAll(checkFilepath, "\\", "/")

	if !strings.HasSuffix(strings.ToLower(checkFilepath), ".stl") {
		return fmt.Errorf("need to point to an STL file")
	}

	return models.CheckSTL(checkFilepath)
}
