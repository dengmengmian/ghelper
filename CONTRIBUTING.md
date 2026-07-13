# Contributing to ghelper

Thank you for your interest in contributing! This document outlines the process for contributing to ghelper.

## Getting Started

1. Fork the repository on GitHub.
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/ghelper.git
   cd ghelper
   ```
3. Create a new branch for your change:
   ```bash
   git checkout -b my-feature
   ```

## Development Setup

- Go 1.21 or later is required.
- Dependencies are managed via Go modules (`go mod tidy`).

## Code Style

- Follow standard Go conventions (`gofmt`, `go vet`).
- Exported functions must have godoc comments starting with the function name.
- Write tests for new functionality.
- Avoid panics in library code; return errors instead.
- Do not use `log.Fatal` or `log.Panic` in library code.

## Running Tests

# All tests with race detection
go test -race ./...

# Specific package
go test -race ./gconvert/

# With coverage
go test -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

## Pull Request Checklist

- [ ] Code compiles (`go build ./...`)
- [ ] Tests pass (`go test -race ./...`)
- [ ] `go vet ./...` is clean
- [ ] New code has godoc comments
- [ ] New functionality has tests
- [ ] No panics or silent error swallowing in exported APIs

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
