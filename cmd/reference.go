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

func init() {
	// scan directory for .pb files, this is how we will generate the name of the commands
	entries, err := os.ReadDir("./proto")
	// make sure the directory exists
	if err != nil {
		fmt.Println("Error reading directory", err)
		os.Exit(1)
	}

	// checking for .pb files
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".pb") {
			// get the name of the command
			referenceCmd := strings.TrimSuffix(entry.Name(), ".pb")

			// create a new command from the name we just got
			rootCmd.AddCommand(&cobra.Command{
				Use:   referenceCmd,
				Short: fmt.Sprintf("%s reference", referenceCmd),
				Run: func(cmd *cobra.Command, args []string) {
					fmt.Println("This is the reference command " + cmd.Use)
				},
			})
			// Here you will define your flags and configuration settings.

			// Cobra supports Persistent Flags which will work for this command
			// and all subcommands, e.g.:
			// referenceCmd.PersistentFlags().String("foo", "", "A help for foo")
			//                  varPersistancy( --word, -flag, defaultValue?, description )
			// referenceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

			// Flags that I would like to implement would be the following:
			// 1. --string, -s, "string to search for"
			// 2. --tag, -t, "tag to search for"
			// 3. --list, -l, "list of all topics stored in the current reference"
            // 4. --add, -a, "add a new topic to the current reference"
            // 5. --remove, -r, "remove a topic from the current reference"

		}
	}
}
