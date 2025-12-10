package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/TobiasAagaard/gitgen/internal/config"
	"github.com/TobiasAagaard/gitgen/testutils"
)

func TestSaveCreatesConfigFile(t *testing.T) {
	cleanup := testutils.UseTempHomeDir(t)
	defer cleanup()

	cfg := config.Config{
		App: config.AppConfig{AIProvider: "claude"},
		Claude: config.ClaudeConfig{
			APIkey:      "test-key",
			Model:       "claude-3-5-sonnet-20241022",
			MaxTokens:   4096,
			Temperature: 0.2,
		},
		Git: config.GitConfig{
			BranchPrefix: "feat",
			CommitStyle:  "conventional",
		},
	}
	if err := config.Save(cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("home: %v", err)
	}
	path := filepath.Join(home, ".config", "gitgen", "config.yaml")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected config file at %s, got: %v", path, err)
	}
}

func TestLoadReadsConfigAndValidatesKey(t *testing.T) {
	cleanup := testutils.UseTempHomeDir(t)
	defer cleanup()

	cfg := config.Config{
		App: config.AppConfig{AIProvider: "claude"},
		Claude: config.ClaudeConfig{
			APIkey:      "valid-key",
			Model:       "claude-3-5-sonnet-20241022",
			MaxTokens:   1234,
			Temperature: 0.1,
		},
		Git: config.GitConfig{
			BranchPrefix: "chore",
			CommitStyle:  "conventional",
		},
	}
	if err := config.Save(cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	if err := config.Load(); err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	got := config.AppConfigState
	if got.App.AIProvider != "claude" {
		t.Errorf("AIProvider = %s, want claude", got.App.AIProvider)
	}
	if got.Claude.APIkey != "valid-key" {
		t.Errorf("API key mismatch, got %q", got.Claude.APIkey)
	}
	if got.Claude.MaxTokens != 1234 {
		t.Errorf("MaxTokens = %d, want 1234", got.Claude.MaxTokens)
	}
	if got.Git.BranchPrefix != "chore" {
		t.Errorf("BranchPrefix = %s, want chore", got.Git.BranchPrefix)
	}
}

func TestLoadFailsWhenMissingClaudeKey(t *testing.T) {
	cleanup := testutils.UseTempHomeDir(t)
	defer cleanup()

	cfg := config.Config{
		App: config.AppConfig{AIProvider: "claude"},
		Claude: config.ClaudeConfig{
			APIkey: "",
			Model:  "claude-3-5-sonnet-20241022",
		},
	}
	if err := config.Save(cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	if err := config.Load(); err == nil {
		t.Fatalf("expected error when API key missing, got nil")
	}
}
