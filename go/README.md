# Go

- Source: The Go Programming Language, Donovan & Kernighan 2016

## Setup

- `export GOPATH=path/to/workspace/`

## Language Properties

- Go is a compiled language; to compile, use `go build`. To build and run, use `go run`
- Go code is organized into packages:
  
A package consists of one or more .go source files defining what the package
does.  Source file begins with `package` declaration, stating which package the
file belongs to, e.g. `package main`.  `package main` and the function `main` in
here is special. It defines standalone executable, and not a library.  Then
comes `import` statements for other packages.

- Newlines are important, they replace semicolons
- Use `gofmt` to rewrite code into standard format
- Use `goimports` to manage import declarations, not part of the standard distribution.
