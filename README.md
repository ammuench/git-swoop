# git-swoop 🐦

A quick little CLI utility to *swoop* to another branch, pull down the latest from remote, and then return to where you started.

## What it does

I found myself frequently needing to jump to our develop branch, pull the latest changes, and then return to my working branch so I could easily merge updates. This tool automates that workflow to make my life a bit easier, and hopefully yours too!

## Install

`git-swoop` can be installed by running:

```bash
go install github.com/ammuench/git-swoop@latest
```

It should then be available in your terminal as `git-swoop`. If you cannot find it, you may need to configure your `GOBIN` or `GOPATH`. [You can read the docs to configure that here](https://go.dev/ref/mod#go-install).

Alternatively, you can clone the repo and build it from source:

```bash
git clone https://github.com/ammuench/git-swoop.git
cd git-swoop
make build
make install
```

`git-swoop` was built with `go1.23.5`, but may work with older versions.

## Usage

### `git-swoop <target-branch-name>`

Tries to check out your target branch, pull latest from remote, and return you to your original branch.

Will try to return you to your original branch if it encounters an error.

Will always log which branch you are on when it exits.

### Example

```bash
# While on feature/my-feature branch
git-swoop main

# This will checkout main, pull latest changes, and return to feature/my-feature
```

### Flags

#### `git-swoop --help` (aliases: `-h`, `-help`)

Prints out a list of all flags and commands.

#### `git-swoop --version` (aliases: `-v`, `-version`)

Prints out version and license info.

## License

`git-swoop` is published under a [GPL V3.0 License that can be found here](https://github.com/ammuench/git-swoop/blob/main/LICENSE.md).

## Contributing

If you have anything you'd like to add, please just open a PR with some comments.

