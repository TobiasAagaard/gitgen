package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

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

	comparison := compareVersions(currentVersion, latestVersion)

	fmt.Printf("Current version: %s\n", currentVersion)
	fmt.Printf("Latest version:  %s\n", latestVersion)

	if comparison < 0 {
		fmt.Printf("\nNew version available: %s → %s\n", currentVersion, latestVersion)
		fmt.Println("Updating...")

		installCmd := exec.Command("go", "install", "github.com/TobiasAagaard/gitgen@"+latest.TagName)
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			return fmt.Errorf("failed to update gitgen: %w", err)
		}

		fmt.Printf("\n✓ Successfully updated to version %s\n", latestVersion)
		fmt.Println("Restart your shell.")
	} else if comparison > 0 {
		fmt.Println("\n✓ You are running a newer version than the latest release.")
	} else {
		fmt.Println("\n✓ You are already running the latest version.")
	}

	return nil
}

func compareVersions(current, latest string) int {
	current = strings.Split(current, "-")[0]
	latest = strings.Split(latest, "-")[0]

	currentParts := strings.Split(current, ".")
	latestParts := strings.Split(latest, ".")

	maxLen := len(currentParts)
	if len(latestParts) > maxLen {
		maxLen = len(latestParts)
	}

	for i := 0; i < maxLen; i++ {
		currentNum := 0
		latestNum := 0

		if i < len(currentParts) {
			var err error
			currentNum, err = strconv.Atoi(currentParts[i])

			if err != nil {
				currentNum = 0
			}
		}
		if i < len(latestParts) {
			var err error
			latestNum, err = strconv.Atoi(latestParts[i])
			if err != nil {
				latestNum = 0
			}
		}

		if currentNum < latestNum {
			return -1
		} else if currentNum > latestNum {
			return 1
		}
	}
	return 0
}

func getLatestRelease() (*githubRelease, error) {
	url := "https://api.github.com/repos/TobiasAagaard/gitgen/releases/latest"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("gitgen/%s (%s/%s)", version.ShortInfo(), runtime.GOOS, runtime.GOARCH))

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusTooManyRequests {
		// Consider including Retry-After / X-RateLimit-Reset details for better UX.
		return nil, fmt.Errorf("GitHub API rate limit exceeded. Try again later")
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 64<<10))
		return nil, fmt.Errorf("GitHub API returned %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}
	return &release, nil
}
