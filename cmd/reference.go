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

func createReferenceCmd() *cobra.Command {
    // make sure the proto directory exists
    entries, err := os.ReadDir("./reference")
    // throw an error if it doesn't
    if err != nil {
        fmt.Println("Error reading directory", err)
        os.Exit(1)
    }

    var referenceCmd *cobra.Command

    // iterate through the entries in the directory
    for _, entry := range entries {
        if strings.HasSuffix(entry.Name(), ".pb") {
            referenceName := strings.TrimSuffix(entry.Name(), ".pb")
            referenceCmd = generateReferenceCmd(referenceName)
        }
    }

    // if there are no references, create a default command that tells the user
    // how to create a new reference.
    if referenceCmd == nil {
        referenceCmd = &cobra.Command{
            Use:   "reference",
            Short: "reference",
            Long: 
    `    This is the default reference command, please create a new reference
    using:

    wut new-ref [referenceName]`,
            Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("This is the reference command, create a new " +
                    "reference with:\n\twut new-ref [referenceName]")
            },
        }
    }

    return referenceCmd
}

func generateReferenceCmd(referenceName string) *cobra.Command {
    referenceCmd := &cobra.Command{
        Use:   referenceName,
        Short: fmt.Sprintf("%s reference", referenceName),
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("This is the reference command " + cmd.Use)
        },
    }
    // The flags for this command are defined here so that the default command
    // has no flags.

    // referenceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

    // Flags that I would like to implement would be the following:
    // 1. --string, -s, "string to search for"
    // 2. --tag, -t, "tag to search for"
    // 3. --list, -l, "list of all topics stored in the current reference"
    // 4. --add, -a, "add a new topic to the current reference"
    // 5. --remove, -r, "remove a topic from the current reference"
    return referenceCmd
}

func init() {
    rootCmd.AddCommand(createReferenceCmd())
}
