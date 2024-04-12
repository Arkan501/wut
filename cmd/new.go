/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "os"
    "strings"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new ref [referenceName]",
	Long: `create a new reference document`,
	Run: func(cmd *cobra.Command, args []string) {
        // since the command name is "new ref" then we need to parse the second
        // argument to get the reference name.
        doesReferenceExist(args[1])
        referenceName := args[1]
		fmt.Println("new reference", referenceName, "created")
        //doesReferenceExist(args[0])

	},
}

func doesReferenceExist(filename string) {
    entries, err := os.ReadDir("proto")
    if err != nil {
        fmt.Println("Error reading directory", err)
        os.Exit(1)
    }

    for _, entry := range entries {
        if strings.Contains(entry.Name(), filename) {
            fmt.Println("Reference already exists")
            os.Exit(1)
        }
    }
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
