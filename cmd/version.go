package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func version() {
	fmt.Println("0.1.0")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Bands CLI version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
