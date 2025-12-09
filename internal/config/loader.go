package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func configHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "gitgen"), nil
}

func ConfigExists() bool {
	dir, err := configHome()
	if err != nil {
		return false
	}
	_, err = os.Stat(filepath.Join(dir, "config.yaml"))
	return err == nil
}

func Load() error {
	dir, err := configHome()
	if err != nil {
		return err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	viper.SetDefault("app.ai_provider", "claude")
	viper.SetDefault("claude.model", "claude-3-5-sonnet-20241022")
	viper.SetDefault("claude.max_tokens", 4096)
	viper.SetDefault("claude.temperature", 0.2)
	viper.SetDefault("git.branch_prefix", "feat")
	viper.SetDefault("git.commit_style", "conventional")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var configCheck Config
	if err := viper.Unmarshal(&configCheck); err != nil {
		return err
	}
	if configCheck.Claude.APIkey == "" {
		return errors.New("missing claude api key")
	}

	AppConfigState = configCheck
	return nil
}

func Save(cfg Config) error {
	dir, err := configHome()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}

	viper.Set("app.ai_provider", cfg.App.AIProvider)
	viper.Set("claude.api_key", cfg.Claude.APIkey)
	viper.Set("claude.model", cfg.Claude.Model)
	viper.Set("claude.max_tokens", cfg.Claude.MaxTokens)
	viper.Set("claude.temperature", cfg.Claude.Temperature)
	viper.Set("git.branch_prefix", cfg.Git.BranchPrefix)
	viper.Set("git.commit_style", cfg.Git.CommitStyle)

	path := filepath.Join(dir, "config.yaml")
	return viper.WriteConfigAs(path)
}
