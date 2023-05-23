# 🛠️ Dotfiles Installer

![Build Status](https://github.com/Mephisto-Grumpy/dotfiles-installer/actions/workflows/go.yml/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/Mephisto-Grumpy/dotfiles-installer)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Mephisto-Grumpy/dotfiles-installer)
![GitHub](https://img.shields.io/github/license/Mephisto-Grumpy/dotfiles-installer)
![GitHub repo size](https://img.shields.io/github/repo-size/Mephisto-Grumpy/dotfiles-installer)

Welcome to Dotfiles Installer! This tool helps you to install and manage your dotfiles with ease.

## 🚀 Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### 📋 Prerequisites

- Go (version 1.20 or later)

### 🏗️ Installation

Clone the repository:

```bash
git clone https://github.com/Mephisto-Grumpy/dotfiles-installer.git
```

Build the project:

```bash
cd dotfiles-installer
make # or make build
```

This will generate a `dotfiles-installer` binary in your current directory.

### 🏃🏻‍♂️ Running the tests

```bash
go test -race -coverprofile=coverage.out -covermode=atomic -tags test ./...
```

### 🐳 Docker for testing

```bash
docker build -t dotfiles-installer .
docker run --rm -it -v $(pwd)/bin/*:/home/ubuntu/bin dotfiles-installer
```

## 📖 Usage

```bash
./dotfiles-installer -h # or ./dotfiles-installer --help
```

## 🛠️ Built With

- [Go](https://golang.org/) - The Go Programming Language
- [GitHub Actions](https://github.com/features/actions) - CI/CD
- [CMake](https://cmake.org/) - Build System

## ✍🏻 Authors

- **Mephisto** - _Initial work_ - [Mephisto-Grumpy](https://github.com/Mephisto-Grumpy)
- **PunGrumpy** - _Development_ - [PunGrumpy](https://github.com/PunGrumpy)

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
