/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:  "new-ref",
    Short: "Create a new reference",
	Long: `create a new reference document`,
	Run: func(cmd *cobra.Command, args []string) {
		makeNewReference(args[0])
	},
}

func makeNewReference(referenceName string) {
    directory := "./proto/"
	entries, err := os.ReadDir(directory)
	// making sure we can read the directory
	if err != nil {
		fmt.Println("Error reading directory", err)
		os.Exit(1)
	}

	fileName := referenceName + ".pb"

	// checking if the reference already exists
	for _, entry := range entries {
		if entry.Name() == fileName {
			fmt.Println("The reference file", fileName, "already exists")
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
	fmt.Println("Reference", referenceName, "created successfully")
}

func init() {
	rootCmd.AddCommand(newCmd)
}
