/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Removes the current blockchain",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		verify := promptui.Prompt{
			Label:     "Are you sure you want to clear the current blockchain: ",
			IsConfirm: true,
		}

		_, err := verify.Run()
		if err != nil {
			fmt.Println("Error clearing blockchain")
			return
		}

		fmt.Println("Clearing blockchain")
		os.RemoveAll("./tmp")
		fmt.Println("Cleared blockchain")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
