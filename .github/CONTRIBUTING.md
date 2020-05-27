# Contributing

We'd love to accept your contributions to this project! There are just a few guidelines you need to follow.

## Bugs

Bug reports should be opened up as [issues](https://help.github.com/en/github/managing-your-work-on-github/about-issues) on the [go-vela/community](https://github.com/go-vela/community) repository!

## Feature Requests

Feature Requests should be opened up as [issues](https://help.github.com/en/github/managing-your-work-on-github/about-issues) on the [go-vela/community](https://github.com/go-vela/community) repository!

## Pull Requests

**NOTE: We recommend you start by opening a new issue describing the bug or feature you're intending to fix. Even if you think it's relatively minor, it's helpful to know what people are working on.**

We are always welcome to new PRs! You can follow the below guide for learning how you can contribute to the project!

## Getting Started

### Prerequisites

* [Review the commit guide we follow](https://chris.beams.io/posts/git-commit/#seven-rules) - ensure your commits follow our standards
* [Golang](https://golang.org/dl/) - for source code and [dependency management](https://github.com/golang/go/wiki/Modules)

### Setup

* [Fork](/fork) this repository

* Clone this repository to your workstation:

```bash
# Clone the project
git clone git@github.com:go-vela/types.git $HOME/go-vela/types
```

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/types
```

* Point the original code at your fork:

```bash
# Add a remote branch pointing to your fork
git remote add fork https://github.com/your_fork/types
```

### Development

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/types
```

* Write your code
  - Please be sure to [follow our commit rules](https://chris.beams.io/posts/git-commit/#seven-rules)

* Write tests for your changes and ensure they pass:

```bash
# Test the code with `go`
go test ./...
```

* Ensure your code meets the project standards:

```bash
# Clean the code with `go`
go mod tidy
go fmt ./...
go vet ./...
```

* Push to your fork:

```bash
# Push your code up to your fork
git push fork master
```

* Open a pull request. Thank you for your contribution!
