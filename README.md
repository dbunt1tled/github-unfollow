# GitHub Followers Manager

A powerful Go-based command-line tool to manage your GitHub following relationships. Follow users who follow a specific user but don't follow you, or unfollow users who don't follow you back.

[![Go Version](https://img.shields.io/badge/go-1.24+-blue?logo=go)](https://golang.org/)
[![Go Reference](https://pkg.go.dev/badge/github.com/dbunt1tled/github-unfollow.svg)](https://pkg.go.dev/github.com/dbunt1tled/github-unfollow)
[![Release](https://img.shields.io/github/v/release/dbunt1tled/github-unfollow)](https://github.com/dbunt1tled/github-unfollow/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/dbunt1tled/github-unfollow)](https://goreportcard.com/report/github.com/dbunt1tled/github-unfollow)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## ✨ Features

- 🔍 **Follow Management**: Follow followers of a specific user who don't follow you back
- 🗑️ **Unfollow Cleanup**: Unfollow users who don't follow you back
- ⚡ **Efficient Processing**: Handle large numbers of followers/following with pagination
- 🛡️ **Safe Execution**: Interactive confirmation before bulk operations
- 🚀 **Concurrent Processing**: Configurable worker pool for fast execution
- 🔐 **Secure Authentication**: OAuth2 token-based GitHub API authentication
- ⚙️ **Flexible Configuration**: Environment variables and command-line flags
- ⏱️ **Rate Limiting**: Configurable delays to respect GitHub API limits

## 📋 Table of Contents

- [Installation](#-installation)
- [Configuration](#-configuration)
- [Usage](#-usage)
- [Project Structure](#-project-structure)
- [Contributing](#-contributing)
- [License](#-license)

## 🚀 Installation

### Prerequisites

- Go 1.24 or higher
- A GitHub account
- A GitHub Personal Access Token with the `user:follow` scope

### Option 1: Build from source

1. Clone the repository:
   ```bash
   git clone https://github.com/dbunt1tled/github-unfollow.git
   cd github-unfollow
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the application:
   ```bash
   go build -o github-followers ./main.go
   ```

### Option 2: Install via Go
```bash
go install github.com/dbunt1tled/github-unfollow@latest
```

## ⚙️ Configuration

### GitHub Personal Access Token

1. Go to GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Generate a new token with the `user:follow` scope
3. Copy the token for configuration

### Environment Variables

Create a `.env` file in the project root or set environment variables:

```env
# Required
GITHUB_USERNAME=your_github_username
GITHUB_TOKEN=your_github_personal_access_token

# Optional (with defaults)
WORKER_COUNT=1      # Number of concurrent workers (default: 1)
QUEUE_SIZE=3        # Size of the worker queue (default: 3)
TIME_DELAY_MS=2000  # Delay between API calls in milliseconds (default: 2000)
```

> **Note**: You can copy `.env.example` to `.env` and fill in your credentials.

## 🚦 Usage

The tool provides two main commands: `follow` and `unfollow`.

### Follow Command

Follow followers of a specific user who don't already follow you:

```bash
# Basic usage - follow followers of 'username' who don't follow you
./github-followers follow username

# Skip confirmation prompt
./github-followers follow username --force
./github-followers follow username -f
```

### Unfollow Command

Unfollow users who don't follow you back:

```bash
# Show users who don't follow you back (dry run)
./github-followers unfollow

# Actually unfollow users who don't follow you back without confirm
./github-followers unfollow --f
./github-followers unfollow --force
```

## 🏗️ Project Structure

```
github-unfollow/
├── cmd/                 # Main application entry point
├── internal/
│   ├── cli/            # Command-line interface implementation
│   ├── config/         # Configuration management
│   ├── git_hub_manager/# GitHub API client and operations
│   ├── helper/         # Utility and helper functions
│   └── worker/         # Worker pool implementation
├── .env.example        # Example environment variables
├── .gitignore          # Git ignore file
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
└── main.go             # Application entry point
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Unfollow Command

Unfollow users who don't follow you back:

```bash
# Basic usage - unfollow users who don't follow you back
./github-unfollow unfollow

# Skip confirmation prompt
./github-unfollow unfollow --force
./github-unfollow unfollow -f
```

**Example workflow:**
1. Tool fetches your followers list
2. Tool fetches your following list
3. Identifies users you follow who don't follow you back
4. Shows the list and asks for confirmation (unless `--force` is used)
5. Unfollows the users concurrently with rate limiting

### Command-Line Flags

Both commands support the following flags:

- `--force, -f`: Skip interactive confirmation prompt (use with caution)
- `--help, -h`: Show help information

### Help Command

```bash
# Show general help
./github-unfollow --help

# Show help for specific commands
./github-unfollow follow --help
./github-unfollow unfollow --help
```

## Configuration Options

### Worker Pool Settings

- **WORKER_COUNT**: Number of concurrent workers for API calls (default: 1)
  - Higher values = faster execution but more API rate limit pressure
  - Recommended: 1-5 for most use cases

- **QUEUE_SIZE**: Size of the task queue (default: 3)
  - Buffer size for pending tasks
  - Should be >= WORKER_COUNT

- **TIME_DELAY_MS**: Delay between API calls in milliseconds (default: 2000)
  - Helps avoid GitHub API rate limits
  - Lower values = faster execution but higher risk of rate limiting
  - Recommended: 1000-3000ms

### Example Configurations

**Conservative (safe for large operations):**
```env
WORKER_COUNT=1
QUEUE_SIZE=3
TIME_DELAY_MS=3000
```

**Balanced:**
```env
WORKER_COUNT=2
QUEUE_SIZE=3
TIME_DELAY_MS=2000
```

## Examples


### Using environment variables
```bash
# Set token and username via environment
export GITHUB_USERNAME="myusername"
export GITHUB_TOKEN="ghp_xxxxxxxxxxxxxxxxxxxx"
export WORKER_COUNT=2
export TIME_DELAY_MS=1500

./github-unfollow unfollow
```

## Architecture

The project is organized into the following components:

- **`main.go`**: Application entry point
- **`cmd/`**: Cobra CLI commands
  - `root.go`: Root command setup
  - `follow.go`: Follow command implementation
  - `unfollow.go`: Unfollow command implementation
- **`internal/`**: Internal packages
  - `config/`: Configuration management with Viper
  - `git_hub_manager/`: GitHub API client wrapper
  - `worker/`: Worker pool for concurrent processing
  - `helper/`: Utility functions

## Security

- Your GitHub token is only used to authenticate with the GitHub API
- The token requires only the `user:follow` scope - no access to repositories or other data
- Tokens are loaded from environment variables, never hardcoded
- Never share your `.env` file or commit it to version control
- The `.env` file is included in `.gitignore` by default

## Rate Limiting

GitHub API has rate limits:
- 5,000 requests per hour for authenticated requests
- Each follow/unfollow operation uses 1 API call
- The tool includes built-in delays (configurable via `TIME_DELAY_MS`)
- Use conservative settings for large operations

## Troubleshooting

### Common Issues

**"Authentication failed" error:**
- Verify your `GITHUB_TOKEN` is correct and has `user:follow` scope
- Check that the token hasn't expired

**"Rate limit exceeded" error:**
- Increase `TIME_DELAY_MS` value
- Reduce `WORKER_COUNT`
- Wait for the rate limit to reset (1 hour)

**"User not found" error:**
- Verify the username exists and is public
- Some users may have restricted follower lists

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This tool is provided as-is, without any warranties. Use it at your own risk. The maintainers are not responsible for any issues caused by using this tool. Always review the list of users before confirming bulk operations.

## Support

If you find this tool useful, please consider giving it a ⭐ on GitHub!