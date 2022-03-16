# Contributing

## Getting Started

We'd love to accept your contributions to this project! If you are a first time contributor, please review our [Contributing Guidelines](https://go-vela.github.io/docs/community/contributing_guidelines/) before proceeding.

### Prerequisites

* [Review the commit guide we follow](https://chris.beams.io/posts/git-commit/#seven-rules) - ensure your commits follow our standards
* Review our [style guide](https://go-vela.github.io/docs/community/contributing_guidelines/#style-guide) to ensure your code is clean and consistent.
* Check out [Make](https://www.gnu.org/software/make/) - start up local development

### Setup

* [Fork](/fork) this repository

* Clone this repository to your workstation:

```bash
# Clone the project
git clone git@github.com:go-vela/cli.git $HOME/go-vela/cli
```

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/cli
```

* Point the original code at your fork:

```bash
# Add a remote branch pointing to your fork
git remote add fork https://github.com/your_fork/cli
```

### Running Locally

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/cli
```

* Build the repository code:

```bash
# execute the `build` target with `make`
make build

# This command will output binaries to the following locations:
#
# * $HOME/go-vela/cli/release/darwin/amd64/vela
# * $HOME/go-vela/cli/release/linux/amd64/vela
# * $HOME/go-vela/cli/release/linux/arm64/vela
# * $HOME/go-vela/cli/release/linux/arm/vela
# * $HOME/go-vela/cli/release/windows/amd64/vela
```

* Run the repository code:

```bash
# run the Go binary for your specific operating system and architecture
release/darwin/amd64/vela
```

### Development

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/cli
```

* Write your code and tests to implement the changes you desire.
 
* Test the repository code (ensures your changes don't break existing functionality):

```bash
# execute the `test` target with `make`
make test
```

* Clean the repository code (ensures your code meets the project standards):

```bash
# execute the `test` target with `make`
make clean
```

* Push to your fork:

```bash
# Push your code up to your fork
git push fork master
```

* Make sure to follow our [PR process](https://go-vela.github.io/docs/community/contributing_guidelines/#development-workflow) when opening a pull request


Thank you for your contribution!
