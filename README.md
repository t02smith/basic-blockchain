# Basic Blockchain

A basic proof of work blockchain based upon [this guide](https://dev.to/nheindev/build-the-hello-world-of-blockchain-in-go-bli) by Noah Hein.

![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)

## Installation

```bash
# 1. clone the repository
git clone https://github.com/t02smith/basic-blockchain.git

# 2. build the application
go install github.com/t02smith/basic-blockchain
```

## Command Line Interface

**[Cobra](https://github.com/spf13/cobra)** is used to implement the CLI and you can find out
 more about the commands by running:

```bash
basic-blockchain --help
```

```bash
A proof of work blockchain based upon the guide by Noah Hein.

Usage:
  basic-blockchain [command]

Available Commands:
  balance     Get the balance at an address
  completion  Generate the autocompletion script for the specified shell
  create      Create a new blockchain
  help        Help about any command
  send        Send currency from one account to another
  show        Print the blockchain

Flags:
  -h, --help   help for basic-blockchain
```
