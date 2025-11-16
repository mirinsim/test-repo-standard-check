# test-repo-standard-check

Simple utility CLI tool - testing repo standards and bootstrap workflow

[![Lint](https://github.com/mirinsim/test-repo-standard-check/workflows/Lint/badge.svg)](https://github.com/mirinsim/test-repo-standard-check/actions?query=workflow%3ALint)
[![Test](https://github.com/mirinsim/test-repo-standard-check/workflows/Test/badge.svg)](https://github.com/mirinsim/test-repo-standard-check/actions?query=workflow%3ATest)

## Features

- Simple CLI interface built with Cobra
- Clean project structure following Go best practices
- Comprehensive CI/CD with linting and testing
- TDD-ready with example tests

## Installation

### From source

```bash
git clone https://github.com/mirinsim/test-repo-standard-check.git
cd test-repo-standard-check
make install
```

### Using go install

```bash
go install github.com/mirinsim/test-repo-standard-check@latest
```

## Usage

```bash
# Show version
test-repo-standard-check version

# Show help
test-repo-standard-check --help
```

## Development

### Prerequisites

- Go 1.22 or later
- Make (optional, for convenience commands)

### Quick Start

```bash
# Clone and enter directory
git clone https://github.com/mirinsim/test-repo-standard-check.git
cd test-repo-standard-check

# Install dependencies
make deps

# Run tests
make test

# Run linter
make lint

# Build binary
make build

# Install to $GOPATH/bin
make install
```

### Available Make Targets

- `make build` - Build the binary
- `make test` - Run tests with coverage
- `make lint` - Run linters
- `make fmt` - Format code
- `make deps` - Download dependencies
- `make install` - Install binary to $GOPATH/bin
- `make clean` - Remove build artifacts

## Project Structure

```
.
├── cmd/              # CLI commands (Cobra)
├── internal/         # Private application code
│   └── version/      # Version information
├── .github/          # GitHub workflows and settings
├── Makefile          # Development tasks
└── main.go           # Application entry point
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
