/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wut",
	Short: "gives a short description of a topic, followed by an example",
	Long: `
    wut is a CLI tool that allows you to store and retrieve information
    about a topic. It gives you, the user, the power to create and manage
    references tailored to your needs. How useful wut is, is dependent on
    the user. How you use this tool is up to you.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wut.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


