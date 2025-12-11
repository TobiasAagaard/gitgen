package cmd

import (
	"testing"

	"github.com/TobiasAagaard/gitgen/internal/config"
	"github.com/TobiasAagaard/gitgen/testutils"
)

func TestRootRunsWithValidConfig(t *testing.T) {
	cleanup := testutils.UseTempHomeDir(t)
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
