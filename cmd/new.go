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
	Use:  "new ref [referenceName]",
	Long: `create a new reference document`,
	Run: func(cmd *cobra.Command, args []string) {
		// since the command name is "new ref" then we need to parse the second
		// argument to get the reference name.
		referenceName := args[1]
		makeNewReference(referenceName)
	},
}

func makeNewReference(refName string) {
    directory := "./proto/"
	entries, err := os.ReadDir(directory)
	// making sure we can read the directory
	if err != nil {
		fmt.Println("Error reading directory", err)
		os.Exit(1)
	}

	fileName := refName + ".pb"

	// checking if the reference already exists
	for _, entry := range entries {
		if strings.Contains(entry.Name(), fileName) {
			fmt.Println("The reference", refName, "already exists")
			// exit out of the loop and the function if the reference already exists
			return
		}
	}

	// may need to change this later.
	filePath := directory + fileName

	_, err = os.Create(filePath)
	if err != nil {
		fmt.Println("There was an error creating the reference", err)
		os.Exit(1)
	}
	fmt.Println("Reference", refName, "created successfully")
}

func init() {
	// add command to wut
	rootCmd.AddCommand(newCmd)
}
