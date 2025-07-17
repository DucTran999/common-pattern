# GO Patterns

A personal collection of patterns, concurrency models, and load balancing techniques explored while learning Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/DucTran999/common-pattern)](https://goreportcard.com/report/github.com/DucTran999/common-pattern)
[![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go)](https://golang.org)
[![codecov](https://codecov.io/gh/DucTran999/common-pattern/graph/badge.svg?token=5XBMMBKCPD)](https://codecov.io/gh/DucTran999/common-pattern)
[![License](https://img.shields.io/github/license/DucTran999/common-pattern)](LICENSE)

# Table of Contents

- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

# Project Structure

The repository is organized into the following directories:

- `concurrency/`: Illustrates Go's concurrency primitives, including goroutines, channels, and synchronization patterns.

- `behavioral/`: Implements common behavioral design patterns such as Strategy, Bridge, and Command.

- `creational/`: Contains examples of creational design patterns like Factory, Singleton, and Builder.

- `dsa/`: Data Structures and Algorithms in Go, including linked lists, trees, sorting algorithms, and more.

- `load-balancing/`: Techniques for distributing workloads across multiple resources, featuring strategies like round-robin and least-connections.

- `utils/`: General-purpose utility functions and helpers used across the project.

- `scripts/`: Automation scripts for testing, building, and deployment tasks.

# Getting Started

## Prerequisites

Ensure the following tools are installed on your machine:

- [**Go 1.23+**](https://go.dev/dl/) — The project requires Go version 1.23 or later.
- [**Taskfile CLI**](https://taskfile.dev/) — Used for task automation and scripting.

## Installation

Clone the repository:

```bash
git clone https://github.com/DucTran999/common-pattern.git
cd common-pattern
```

# Usage

Each directory contains standalone Go programs illustrating specific patterns or algorithms. Navigate to the desired directory and run the examples:

```sh
# use task command to run the app
task run
```

Then a menu will be shown in terminal:

```sh
0 issues.
========== MENU ==========
1) worker pool
2) generator
3) multiplexing
===== Load balancing =====
4) load balance alg: round-robin
0) Exit
==========================
Choose an option:
```

# Contributing

Contributions are welcome! If you have suggestions for improvements or new patterns to include, please open an issue or submit a pull request.

# License

This project is licensed under the MIT License.
