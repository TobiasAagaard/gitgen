# gitgen

 CLI tool to generate git commit messages and branch names

## Features

-  **AI-Powered**: Uses Claude AI to analyze your code changes and generate meaningful suggestions
-  **Commit Messages**: Generate commit messages following conventional commits or your preferred style
-  **Branch Names**: Generate descriptive branch names based on your changes
-  **Combined Mode**: Generate both commit message and branch name in one go
-  **Preview Mode**: Preview suggestions before executing
-  **Configurable**: Customize styles, prefixes, and AI parameters
-  **Best Practices**: Follows Cobra and Viper patterns for CLI apps

## Installation

```bash
go install github.com/TobiasAagaard/gitgen@latest
```

Or build from source:

```bash
git clone https://github.com/TobiasAagaard/gitgen.git
cd gitgen
go build -o gitgen
```

## Quick Start

 **Make some changes to your code and try it out:**
   ```bash
   # Preview commit message
   gitgen commit
   
   # Preview branch name
   gitgen branch
   
   # Preview both
   gitgen generate
   ```

## Usage

### Generate Commit Message

```bash
# Preview commit message for staged changes
gitgen commit

# Preview with different style
gitgen commit --style detailed

# Stage all changes and create commit
gitgen commit --stage --execute

# Create commit for already staged changes
gitgen commit --execute
```

### Generate Branch Name

```bash
# Preview branch name
gitgen branch

# Create branch
gitgen branch --execute

# Create branch and check it out
gitgen branch --execute --checkout

# Use custom prefix
gitgen branch --execute --prefix bugfix/
```

### Generate Both (Recommended)

```bash
# Preview both commit message and branch name
gitgen generate

# Create branch and commit in one go
gitgen generate --stage --execute --checkout

# With custom settings
gitgen generate --execute --style conventional --prefix feature/
```


### Commit Styles

- **conventional**: Follows [Conventional Commits](https://www.conventionalcommits.org/) format
  ```
  feat(auth): add OAuth2 authentication
  ```

- **simple**: Short and concise messages
  ```
  Add user authentication
  ```

- **detailed**: Comprehensive messages with context
  ```
  Add OAuth2 authentication
  
  Implemented OAuth2 flow with Google and GitHub providers. 
  Added user session management and token refresh logic.
  ```

## Commands

| Command | Description |
|---------|-------------|
| `gitgen init` | Initialize configuration file |
| `gitgen commit` | Generate and optionally create commit |
| `gitgen branch` | Generate and optionally create branch |
| `gitgen generate` | Generate both commit and branch |
| `gitgen --help` | Show help for any command |

## Flags

### Global Flags
- `--config string`: Config file path (default: `$HOME/.gitgen.yaml`)

### Commit Flags
- `-e, --execute`: Execute the git commit command
- `-s, --staged`: Use staged changes (default: true)
- `-a, --stage`: Stage all changes before committing
- `--style string`: Commit message style

### Branch Flags
- `-e, --execute`: Execute the git branch command
- `-c, --checkout`: Checkout branch after creating
- `-p, --prefix string`: Branch name prefix
- `-s, --staged`: Use staged changes

### Generate Flags
- `-e, --execute`: Execute both commands
- `-c, --checkout`: Checkout new branch
- `-a, --stage`: Stage all changes
- `--style string`: Commit message style
- `-p, --prefix string`: Branch name prefix


### Typical Workflow

```bash
# 1. Make changes to your code
vim myfile.go

# 2.  Preview what gitgen suggests
gitgen generate

# 3. If you like it, execute
gitgen generate --stage --execute --checkout

# Result: New branch created, changes staged and committed! 
```

### Preview Before Committing

```bash
# Stage your changes
git add . 

# Preview commit message
gitgen commit

# If satisfied, execute
gitgen commit --execute
```

### Create Feature Branch

```bash
# Make changes
vim feature. go

# Generate and create feature branch
gitgen branch --execute --prefix feature/ --checkout
```

## Requirements

- Go 1.21 or higher
- Git installed and configured
- Claude API key ([Get one here](https://console.anthropic.com/))

## Development

```bash
# Clone repository
git clone https://github.com/TobiasAagaard/gitgen.git
cd gitgen

# Install dependencies
go mod download

# Build
go build

# Run tests
go test ./... 
```

## Architecture

```
gitgen/
├── cmd/                    # Cobra commands
│   ├── root.go            # Root command
│   ├── commit.go          # Commit command
│   ├── branch.go          # Branch command
│   ├── generate.go        # Generate both command
│   └── init.go            # Init command
├── internal/
│   ├── config/            # Viper configuration
│   │   └── config.go
│   ├── git/               # Git operations
│   │   └── git.go
│   └── ai/                # Claude AI integration
│       └── claude.go
└── main.go                # Entry point
```



## License

MIT License - see LICENSE file for details
