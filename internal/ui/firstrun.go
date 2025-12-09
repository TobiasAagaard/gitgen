// internal/ui/firstrun.go
package ui

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/TobiasAagaard/gitgen/internal/config"
)

func RunFirstTimeSetup() error {
	ai := ""
	if err := survey.AskOne(&survey.Select{
		Message: "Choose AI provider",
		Options: []string{"claude"},
		Default: "claude",
	}, &ai); err != nil {
		return err
	}

	var apiKey string
	if err := survey.AskOne(&survey.Password{
		Message: "Enter Claude API key",
	}, &apiKey, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	model := "claude-3-5-sonnet-20241022"
	if err := survey.AskOne(&survey.Input{
		Message: "Claude model (press Enter for default)",
		Default: model,
	}, &model); err != nil {
		return err
	}

	cfg := config.Config{
		App: config.AppConfig{AIProvider: ai},
		Claude: config.ClaudeConfig{
			APIkey:      apiKey,
			Model:       model,
			MaxTokens:   4096,
			Temperature: 0.2,
		},
		Git: config.GitConfig{
			BranchPrefix: "feat",
			CommitStyle:  "conventional",
		},
	}

	if err := config.Save(cfg); err != nil {
		return err
	}
	config.AppConfigState = cfg
	return nil
}
