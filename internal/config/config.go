package config

type Config struct {
	Claude ClaudeConfig `mapstructure:"claude"`
	Git    GitConfig    `mapstructure:"git"`
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

var AppConfig Config

func InitConfig() {
	// Configuration initialization logic goes here
}
