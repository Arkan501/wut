/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
    "log"
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
    directory := "./reference/"
	entries, err := os.ReadDir(directory)
	// making sure we can read the directory
	if err != nil {
        log.Fatal("Error reading directory", err)
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
        log.Fatal("There was an error creating the reference", err)
	}
	fmt.Println("Reference", fileName, "created successfully")
}

func init() {
	rootCmd.AddCommand(newCmd)
}
