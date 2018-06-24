# Go

Googles's Go programming language. 

Learned for [Censored Planet](https://censoredplanet.com),
namely to implement [Satellite](https://www.usenix.org/system/files/conference/atc16/atc16_paper-scott.pdf) and
[Iris](https://people.eecs.berkeley.edu/~pearce/papers/dns_usenix_2017.pdf), as well as maintain
[Augur](https://people.eecs.berkeley.edu/~pearce/papers/augur_oakland_2017.pdf).

## Sources

- [The Go Programming Language, Donovan & Kernighan](https://www.gopl.io/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/list)

## Setup

- Important reads: [Installation](https://golang.org/doc/install) and [setup](https://golang.org/doc/code.html).

- `GOPATH` is an environment variable defining the user workspace. The three main directories
are `src` for source files, `pkg` for OS specific package objects, and `bin` for binaries from
compiling source files. Add `bin` to PATH:

```bash
export PATH=$PATH:/usr/local/go/bin
```

- Import statements are resolved through the `GOPATH`, meaning that any installed packages will
and must go here. Using `go get` will automatically do this.

## Notes

- Go is a compiled language; to compile, use `go build`. To build and run, use `go run`.

- Go code is organized into packages. Every source file belongs to a package, and convention is for 
every package source file to be in its own directory. Don't put source files belonging to different
packages in the same directory. When in a directory containing source files,
`go install` will build the package to the workspace's `bin`.

- Package `main` is special. It indicates that the package we are making is an executable, and not 
just a library of functions.

- Newlines are important, they replace semicolons.

- Use `gofmt` to rewrite code into standard format.

## Practice

File | Description | Concepts
---- | ----------- | --------
`hello.go` | Customary hello world program | Code structure basics, printing
`loop.go` | Echoes command line arguments to stdout | Variable declarations, for/while/range-based loops
`if.go` | Input lines through stdin, output lines that have duplicates | If statements, stdin, maps
`file.go` | Input lines through a file, output lines that have duplicates | Files
`spinner.go` | Spawn thread that outputs spinner to stdout, to show program is still running | Goroutine
`clock.go` | Server that handles clients concurrently. Sends current time at second intervals to client | Goroutine, networking
`netcat.go` | Netcat like program to connect to server, and pipe to stdout | Channels
`echo.go` | Concurrent server that echoes whatever its client sends | Goroutine, channels, networking
