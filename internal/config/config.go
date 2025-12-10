package config

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Claude ClaudeConfig `mapstructure:"claude"`
	Git    GitConfig    `mapstructure:"git"`
}

type AppConfig struct {
	AIProvider string `mapstructure:"ai_provider"`
}

type ClaudeConfig struct {
	APIkey      string  `mapstructure:"api_key"`
	Model       string  `mapstructure:"model"`
	MaxTokens   int     `mapstructure:"max_tokens"`
	Temperature float64 `mapstructure:"temperature"`
}

type GitConfig struct {
	BranchPrefix string `mapstructure:"branch_prefix"`
	CommitStyle  string `mapstructure:"commit_style"`
}

var AppConfigState Config

func InitConfig() {
	// Configuration initialization logic goes here
}
