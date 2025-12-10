# GitGen CLI
![CodeRabbit Pull Request Reviews](https://img.shields.io/coderabbit/prs/github/TobiasAagaard/gitgen?utm_source=oss&utm_medium=github&utm_campaign=TobiasAagaard%2Fgitgen&labelColor=171717&color=FF570A&link=https%3A%2F%2Fcoderabbit.ai&label=CodeRabbit+Reviews)

GitGen helps generate AI-assisted Git branch names and commit messages. On first run, it guides you through a quick setup to choose the AI provider (currently Claude) and store your API key locally.

## Install

```zsh
# Ensure Go bin is on PATH (once)
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc && source ~/.zshrc

# Install latest version from your repo
go install github.com/TobiasAagaard/gitgen@latest

# Verify
gitgen --help
```

## Update

```zsh
go install github.com/TobiasAagaard/gitgen@latest
```

## Usage

```zsh
# Run the CLI
gitgen

# First Run gitgen or gitgen setup
- Choose the AI provider: `only supports claude API`
- Enter your Claude API key (hidden input)
- Optionally set the model (default: `claude-3-5-sonnet-20241022`)

Config is saved to `~/.config/gitgen/config.yaml`.
```

## Documentation

- Commands:
	- `gitgen`, `gitgen setup`
		- Runs the app. If configuration is missing or invalid, it launches interactive setup.
- Config file: `~/.config/gitgen/config.yaml`
	- `claude.api_key`: Claude API key
	- `claude.model`: Claude model name
	- `claude.max_tokens`: Token limit (default `4096`)
	- `claude.temperature`: Generation temperature (default `0.2`)
	- `git.branch_prefix`: Branch prefix (default `feat`)
	- `git.commit_style`: Commit style (default `conventional`)

## Troubleshooting
- Command not found:
	- Ensure `~/go/bin` or `/usr/local/bin` is on your `PATH`.
- "no required module provides package" or missing `go.sum`:
	- Run `go mod tidy` in the repo.
- Case-sensitive filename issues on macOS:
	- Use lowercase file names consistently (e.g., `internal/config/config.go`, `internal/config/loader.go`, `internal/ui/firstrun.go`).


## CI
The GitHub Actions workflow uses `go-version-file: go.mod`, so CI installs Go `1.25.x` as declared in `go.mod`.
