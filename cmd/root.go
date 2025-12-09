package cmd

import (
	"fmt"
	"os"

	"github.com/TobiasAagaard/gitgen/internal/config"
	"github.com/TobiasAagaard/gitgen/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitgen",
	Short: "Gitgen is a tool to generate ai generated git commit and branch names/messages",
	Long: `Gitgen is a tool to generate ai generated git commit and branch names/messages.
	It uses the AI model of your choice to generate commit messages and branch names based on the changes made in the repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.ConfigExists() {
			if err := config.Load(); err != nil {
				return err
			}
		} else {
			if err := ui.RunFirstTimeSetup(); err != nil {
				return err
			}
		}
		fmt.Println("GitGen is configured. Ready.")
		return nil
	},
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Run to setup or reconfigure Gitgen",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ui.RunFirstTimeSetup()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
