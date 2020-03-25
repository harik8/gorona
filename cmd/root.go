package cmd

import (
	"os"

	cli "github.com/harik8/gorona/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gorona",
	Short: "gorona cli",
	Long:  "gorona | gorona <country>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cli.GetCountries()
		} else if len(args) == 1 {
			cli.GetCountry(args[0])
		} else {
			if err := cmd.Help(); err != nil {
				os.Exit(1)
			}
		}
	},
}

// Execute : cmd Execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
