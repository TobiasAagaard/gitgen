package cmd

import (
	"os"
	"testing"

	"github.com/TobiasAagaard/gitgen/internal/config"
	"github.com/spf13/viper"
)

func ExecuteNoExitForTest() error {
	return rootCmd.Execute()
}

func useTempHomeDir(t *testing.T) func() {
	t.Helper()
	tmp := t.TempDir()
	oldHome, had := os.LookupEnv("HOME")
	_ = os.Setenv("HOME", tmp)
	return func() {
		if had {
			_ = os.Setenv("HOME", oldHome)
		} else {
			_ = os.Unsetenv("HOME")
		}
		viper.Reset()
	}
}

func TestRootRunsWithValidConfig(t *testing.T) {
	cleanup := useTempHomeDir(t)
	defer cleanup()

	cfg := config.Config{
		App: config.AppConfig{AIProvider: "claude"},
		Claude: config.ClaudeConfig{
			APIkey:      "testkey",
			Model:       "claude-3-5-sonnet-20241022",
			MaxTokens:   4096,
			Temperature: 0.2,
		},
		Git: config.GitConfig{
			BranchPrefix: "feature/",
			CommitStyle:  "conventional",
		},
	}
	if err := config.Save(cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}
	if err := ExecuteNoExitForTest(); err != nil {
		t.Fatalf("root Execute failed: %v", err)
	}
}
