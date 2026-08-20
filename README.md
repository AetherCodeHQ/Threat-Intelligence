# Threat Intelligence

![CI](https://github.com/Qyroxen/Threat-Intelligence/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Threat-Intelligence/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Threat-Intelligence?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Threat-Intelligence)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Threat-Intelligence)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Threat-Intelligence?style=social)](https://github.com/Qyroxen/Threat-Intelligence/stargazers)

## What is it?

Threat Intelligence is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Threat-Intelligence.git
cd Threat-Intelligence
go build -o threatintelligence .

# Run
./threatintelligence --help
```

## CLI Usage

```bash
# Basic usage
./threatintelligence

# With flags
./threatintelligence --verbose --output json

# Get help
./threatintelligence --help
```

## Examples

```bash
# Example 1
./threatintelligence example1

# Example 2
./threatintelligence example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o threatintelligence .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Threat-Intelligence/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Threat-Intelligence?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Threat-Intelligence/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Threat-Intelligence?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Threat-Intelligence/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Threat-Intelligence" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Threat-Intelligence/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Threat-Intelligence" alt="Pull Requests">
  </a>
</p>
