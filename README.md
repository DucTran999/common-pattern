# GO Pattern
A personal collection of patterns, concurrency models, and load balancing techniques explored while learning Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/DucTran999/common-pattern)](https://goreportcard.com/report/github.com/DucTran999/common-pattern)
[![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go)](https://golang.org)
[![License](https://img.shields.io/github/license/DucTran999/common-pattern)](LICENSE)
![CodeRabbit Pull Request Reviews](https://img.shields.io/coderabbit/prs/github/DucTran999/common-pattern?utm_source=oss&utm_medium=github&utm_campaign=DucTran999%2Fcommon-pattern&labelColor=171717&color=FF570A&link=https%3A%2F%2Fcoderabbit.ai&label=CodeRabbit+Reviews)

# Table of Contents
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

# Project Structure
The repository is organized into the following directories:

- `concurrency/`: Examples demonstrating Go's concurrency primitives, including goroutines, channels, and synchronization techniques.

- `creational/`: Implementation of the Singleton design pattern, ensuring a class has only one instance and providing a global point of access to it.

- `dsa/`: Data Structures and Algorithms implemented in Go, covering common structures like linked lists, trees, and sorting algorithms.

- `load-balancing/`: Strategies and implementations for distributing workloads across multiple resources, including round-robin and least connections algorithms.

- `utils/`: Utility functions and helpers to support the main implementations.

- `scripts/`: Automation scripts for tasks such as testing, building, and deployment.

# Getting Started

## Prerequisites
- Go 1.23 or later installed on your machine.
- Taskfile CLI installed (https://taskfile.dev/)

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