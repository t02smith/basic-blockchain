# Basic Blockchain

A basic proof of work blockchain based upon [this guide](https://dev.to/nheindev/build-the-hello-world-of-blockchain-in-go-bli) by Noah Hein.

> Project based in Go 1.16

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

Typical usage will follow:

```bash
# 1. generate your wallets
basic-blockchain create wallet -n 3

# 2. generate a new blockchain
basic-blockchain create blockchain

# 3. send currency between your wallets
basic-blockchain send
```
