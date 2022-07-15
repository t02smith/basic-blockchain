package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new blockchain",
	Long:  `Create a new blockchain and send the reward for mining the genesis block the given address.`,
}

func init() {
	rootCmd.AddCommand(createCmd)

}
