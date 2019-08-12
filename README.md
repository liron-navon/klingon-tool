[![CircleCI][circle ci badge]][circle ci]

# klingon-tool

A tool to translate names from english to klingon, and identifying the species of characters from star track.

This project uses gomod for managing dependencies, and makefile as a script runner.

![preview][preview]

### Requirements:
[golang][go] 1.11 or above (I used gomod)

### Running the app:
1. Download the program
2. call `make start` or `go run main.go <NAME_CHARACTER>`, or run the executable on a `mac` like `./klingon-tool <NAME_CHARACTER>`.

### Commands:
* `make test`: run tests
* `make start`: run the app
* `make build`: builds an executable for the client os

### Tests:
Tests were written with go's testing library, and testify as an assertion library.
You can run the tests using `make test` or `go test -cover ./...`

#### Notes:
I took the liberty to not make the app fail in case no name was passed to it using the command line, and allowing the user to enter a name after running the app.

[circle ci badge]: https://circleci.com/gh/liron-navon/klingon-tool.svg?style=svg
[circle ci]: https://circleci.com/gh/liron-navon/klingon-tool
[go]: https://golang.org/
[preview]: https://i.imgur.com/gRt5GRj.png
