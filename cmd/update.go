package cmd

import (
	"fmt"

	"github.com/TobiasAagaard/gitgen/pkg/version"
	"github.com/spf13/cobra"
)

type githubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}

func runUpdateCommand(cmd *cobra.Command, args []string) error {
	fmt.Println("Checking for updates...")

	latest, err := getLatestRelease()
	if err != nil {
		return fmt.Errorf("failed to check for updates: %w", err)
	}

	currentVersion := version.ShortInfo()
	if currentVersion == "dev" {
		fmt.Println("Running development version. Use 'go install github.com/TobiasAagaard/gitgen@latest' to update.")
		return nil
	}
	return nil
}

func getLatestRelease() (*githubRelease, error) {
	return nil, nil
}
