# go-project-template-cli

There is a certain amount of work to be done before you can implement the features of your Go powered CLI app.

A few of those best practices, libraries, and structural elements are bundled in this repository. You can use it to bootstrap your own codebase.

The code is platform agnostic while the build tools are tailored to *NIX systems (Linux, MacOS X). Your local build environment should include these tools:

* `sh` and `sed` (see initial-setup.sh)
* `make` (see [Makefile](Makefile]))
* `upx` (optional)
* the `go` compiler (obviously)

## How to Use

Open a shell and execute this block:

``` bash
PROJECT_NAME=myawesometool # define your project name
git clone https://github.com/jojomi/go-project-template-cli.git "${PROJECT_NAME}" # clone this repo
cd "${PROJECT_NAME}" # change to project directory
./intitial-setup.sh # follow instructions
```

After that, replace this file ([`README.md`](README.md)) with more useful content for your project and sail away!

Good luck with your project and never forget where you came from ;).

## Best Practices Used

Non-exhaustive list of best practices in this template:

* [spf13/cobra](https://github.com/spf13/cobra) for CLI-interaction, commands, and flags
* [rs/zerolog](https://github.com/rs/zerolog) for (structured) logging
* reproducible builds (see [here](https://blog.filippo.io/reproducing-go-binaries-byte-by-byte/))

If you want to dig deeper, have a look [here](go.mod).

## How to contribute

Additions and changes are welcome. Discussing them in a issue[link] prior to a MR[link] would be useful. Forking this repository is encouraged as well of course.