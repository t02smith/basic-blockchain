package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new object related to the blockchain",
	Long:  `See "basic-blockchain --help" for sub-commands`,
}

func init() {
	rootCmd.AddCommand(createCmd)

}
