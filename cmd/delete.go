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
        // parse the argument passed to delete to get the file we are looking for.
        fileName := args[0] + ".pb"
        directory := "./proto/"
        entries, err := os.ReadDir(directory)

        // make sure the directory exists
        if err != nil {
            fmt.Println("Error reading directory", err)
            os.Exit(1)
        }

        // make sure the entry exists
        for _, entry := range entries {
            if entry.Name() == fileName {
                err := os.Remove(directory + fileName)
                if err != nil {
                    fmt.Println("Error deleting reference", err)
                    os.Exit(1)
                }
                fmt.Println("Reference", args[0], "deleted successfully")
                return
            }
        }
        fmt.Println("Reference", args[0], "does not exist")
        return

	},
}


func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
