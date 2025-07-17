# GO Patterns

[![Go Report Card](https://goreportcard.com/badge/github.com/DucTran999/common-pattern)](https://goreportcard.com/report/github.com/DucTran999/common-pattern)
[![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go)](https://golang.org)
[![codecov](https://codecov.io/gh/DucTran999/common-pattern/graph/badge.svg?token=5XBMMBKCPD)](https://codecov.io/gh/DucTran999/common-pattern)
[![License](https://img.shields.io/github/license/DucTran999/common-pattern)](LICENSE)

A personal collection of patterns, concurrency models, and load balancing techniques explored while learning Go.

---

## Project Structure

The repository is organized into the following directories:

- `concurrency/`: Illustrates Go's concurrency primitives, including goroutines, channels, and synchronization patterns.

- `behavioral/`: Implements common behavioral design patterns such as Strategy, Bridge, and Command.

- `creational/`: Contains examples of creational design patterns like Factory, Singleton, and Builder.

## Prerequisites

Ensure the following tools are installed on your machine:

- [**Go 1.23+**](https://go.dev/dl/) — The project requires Go version 1.23 or later.
- [**Taskfile CLI**](https://taskfile.dev/) — Used for task automation and scripting.

---

## Installation

Clone the repository:

```bash
git clone https://github.com/DucTran999/common-pattern.git
cd common-pattern
```

---

## 🧪 Testing

This project uses Go's built-in testing framework with mocks and table-driven tests.

```bash
go test -v ./...
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! If you have suggestions for improvements or new patterns to include, please open an issue or submit a pull request.
