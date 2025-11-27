/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitgen",
	Short: "Gitgen is a tool to generate ai generated git commit and branch names/messages",
	Long: `Gitgen is a tool to generate ai generated git commit and branch names/messages.
It uses the AI model of your choice to generate commit messages and branch names based on the changes made in the repository.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
