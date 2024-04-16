/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the current version",
	Long: `It's not that deep bro`,
	Run: func(cmd *cobra.Command, args []string) {      // This is the function that will run when the command is called
        // hardcoded for now
		fmt.Println("wut alpha")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
