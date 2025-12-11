package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/TobiasAagaard/gitgen/pkg/version"
	"github.com/spf13/cobra"
)

type githubRelease struct {
	TagName string `json:"tag_name"`
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

	currentVersion = strings.TrimPrefix(currentVersion, "v")
	latestVersion := strings.TrimPrefix(latest.TagName, "v")

	if latestVersion > currentVersion {
		fmt.Printf("New version available: %s\n", latest.TagName)
		fmt.Println("Update by using command gitgen update")
	} else {
		fmt.Println("You are running the latest version.")
	}

	fmt.Printf("Current version: %s\n", currentVersion)
	fmt.Printf("Latest version:  %s\n", latestVersion)

	fmt.Println("Updating...")
	installCmd := exec.Command("go", "install", "github.com/TobiasAagaard/gitgen@latest")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to update gitgen: %w", err)
	}

	fmt.Printf("\n✓ Successfully updated to version %s\n", latestVersion)
	fmt.Println("Restart your shell or run 'hash -r' to use the new version.")
	return nil
}

func getLatestRelease() (*githubRelease, error) {
	url := "https://api.github.com/repos/TobiasAagaard/gitgen/releases/latest"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("gitgen/%s (%s/%s)", version.ShortInfo(), runtime.GOOS, runtime.GOARCH))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API returned %d: %s", resp.StatusCode, string(body))
	}
	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}
	return &release, nil
}
