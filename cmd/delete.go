/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [referenceName]",
	Short: "Delete a reference",
	Long: `Delete a reference by providing it's name`,
	Run: func(cmd *cobra.Command, args []string) {
        deleteReference(args[0])
	},
}

func deleteReference(reference string) {
    // parse the argument passed to delete to get the file we are looking for.
    fileName := reference + ".pb"
    directory := "./proto/"
    entries, err := os.ReadDir(directory)

    // make sure the directory exists
    if err != nil {
        fmt.Println("Error reading directory", err)
        os.Exit(1)
    }

    // check if the entry exists
    for _, entry := range entries {
        if entry.Name() == fileName {
            // delete the file if it exists
            err := os.Remove(directory + fileName)
            // return an error if the file could not be deleted
            if err != nil {
                fmt.Println("Error deleting reference", err)
                os.Exit(1)
            }
            fmt.Println("Reference", reference, "deleted successfully")
            return
        }
    }
    // if the reference does not exist, return an error
    fmt.Println("Reference", reference, "does not exist")
    return

}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
