package cmd

import (
	"fmt"
	"os"

	"github.com/TobiasAagaard/gitgen/internal/config"
	"github.com/TobiasAagaard/gitgen/internal/ui"
	"github.com/TobiasAagaard/gitgen/pkg/version"
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Run to setup or reconfigure Gitgen from terminal",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ui.RunFirstTimeSetup()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of Gitgen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Info())
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for updates to Gitgen",
	RunE:  runUpdateCommand,
}

func init() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(updateCmd)
}
