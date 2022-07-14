package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var address string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "basic-blockchain",
	Short: "A basic proof of work blockchain",
	Long:  `A proof of work blockchain based upon the guide by Noah Hein.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
