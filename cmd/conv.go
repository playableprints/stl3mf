package cmd

import (
	"fmt"
	"stl3mf/models"
	"strings"

	"github.com/spf13/cobra"
)

var (
	convFilepath string
	convCheck    bool
)

func init() {

	convCmd.Flags().StringVarP(&convFilepath, "file", "f", "", "STL filepath")
	convCmd.Flags().BoolVarP(&convCheck, "check", "c", false, "Check file before convert")

	rootCmd.AddCommand(convCmd)
}

var convCmd = &cobra.Command{
	Use:   "conv",
	Short: "Convert an STL",
	RunE:  convert,
}

func convert(cmd *cobra.Command, args []string) error {
	convFilepath = strings.ReplaceAll(convFilepath, "\\", "/")

	if !strings.HasSuffix(strings.ToLower(convFilepath), ".stl") {
		return fmt.Errorf("need to point to an STL file")
	}

	var err error

	if convCheck {
		err = models.CheckSTL(convFilepath)
	}

	if err == nil {
		err = models.ConvertToSTL(convFilepath)
	}

	return err
}
