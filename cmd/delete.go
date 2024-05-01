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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [referenceName]",
	Short: "Delete a reference",
	Long: `
    This command is to be used when you want to delete a reference that
    you no longer need. It deletes the WHOLE reference file, so be careful. if 
    you are just trying to delete something specific in the reference, you should
    use the {needsToBeImplemented} command instead.`,
	Run: func(cmd *cobra.Command, args []string) {
        deleteReference(args[0])
	},
}

func deleteReference(reference string) {
    // parse the argument passed to delete to get the file we are looking for.
    fileName := reference + ".pb"
    directory := "./reference/"
    entries, err := os.ReadDir(directory)

    // make sure the directory exists
    if err != nil {
        log.Fatal("Error reading directory", err)
    }

    // check if the entry exists
    for _, entry := range entries {
        if entry.Name() == fileName {
            // delete the file if it exists
            err := os.Remove(directory + fileName)
            // return an error if the file could not be deleted
            if err != nil {
                log.Fatal("Error deleting reference", err)
            }
            fmt.Println("Reference", fileName, "deleted successfully")
            return
        }
    }
    // if the reference does not exist, return an error
    fmt.Println("Reference", fileName, "does not exist")
    os.Exit(1)

}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
