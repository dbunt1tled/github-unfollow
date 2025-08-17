# GitHub Unfollow Tool

A Go-based command-line tool to help you clean up your GitHub following list by automatically identifying and unfollowing users who don't follow you back.

[![Go Version](https://img.shields.io/badge/go-1.21+-blue?logo=go)](https://golang.org/)
[![Go Reference](https://pkg.go.dev/badge/github.com/dbunt1tled/github-unfollow.svg)](https://pkg.go.dev/github.com/dbunt1tled/github-unfollow)
[![Release](https://img.shields.io/github/v/release/dbunt1tled/github-unfollow)](https://github.com/dbunt1tled/github-unfollow/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/dbunt1tled/github-unfollow)](https://goreportcard.com/report/github.com/dbunt1tled/github-unfollow)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- 🔍 Find users you follow who don't follow you back
- ⚡ Process large numbers of followers/following efficiently
- 🛡️ Safe execution with interactive confirmation
- 🚀 Concurrent processing with configurable worker pool
- 🔐 Secure token-based authentication
- ⚙️ Configurable through environment variables and command-line flags

## Prerequisites

- Go 1.16 or higher
- A GitHub account
- A GitHub Personal Access Token with the `user:follow` scope

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/github-unfollow.git
   cd github-unfollow
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the project root and add your GitHub credentials:
   ```
   GITHUB_USER=your_github_username
   GITHUB_TOKEN=your_github_personal_access_token
   ```

## Usage

1. Build the application:
   ```bash
   go build -o github-unfollow cmd/main.go
   ```

2. Run the application:
   ```bash
   ./github-unfollow
   ```

3. The tool will:
   - Fetch your followers and following lists
   - Identify users who don't follow you back
   - Show you the list of users to be unfollowed
   - Ask for confirmation before proceeding (unless `-force` is used)
   - Process unfollows concurrently using the configured worker pool

## Configuration

### Environment Variables
Create a `.env` file in the project root with the following variables:

```
# Required
GITHUB_USER=your_github_username
GITHUB_TOKEN=your_github_personal_access_token

# Optional (with defaults)
WORKER_COUNT=10  # Number of concurrent workers
QUEUE_SIZE=10     # Size of the worker queue
```

### Command-Line Flags

- `-force`: Skip the interactive confirmation prompt (use with caution)

Example:
```bash
# Run with default settings and interactive confirmation
./github-unfollow

# Run with force flag (skips confirmation)
./github-unfollow -force
```

## Security

- Your GitHub token is only used to authenticate with the GitHub API
- The token requires only the `user:follow` scope
- Never share your `.env` file or commit it to version control

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This tool is provided as-is, without any warranties. Use it at your own risk. The maintainers are not responsible for any issues caused by using this tool.
