
# gocompat [![godoc](https://godoc.org/github.com/smola/gocompat?status.svg)](https://godoc.org/github.com/smola/gocompat) [![Build Status](https://travis-ci.org/smola/gocompat.svg)](https://travis-ci.org/smola/gocompat) [![codecov.io](https://codecov.io/github/smola/gocompat/coverage.svg)](https://codecov.io/github/smola/gocompat)

**gocompat** is a tool to check compatibility between Go API versions.

## Usage

### Listing all symbols

**gocompat** considers an API as all exported symbols in a given set of packages as well as all exported symbols reachable from them. You can check this for the current package as follows:

```
gocompat reach .
```

### Comparing two versions

**gocompat** can compare the API of two git references in a repository. For example:

```
gocompat compare v0.1.0 master ./...
```

### What is a backwards compatible change?

There is almost no API change in Go that is fully backwards compatibility ([see this post for more](https://blog.merovius.de/2015/07/29/backwards-compatibility-in-go.html)). By default, gocompat uses a strict approach in which most changes to exported symbols are considered incompatible.

However, most users will probably want to use compatibility guarantees analogous to the [Go 1 compatibility promise](https://golang.org/doc/go1compat). You can use the `--go1compat` for that, which is a shorthand for `--exclude=SymbolAdded --exclude=FieldAdded --exclude=MethodAdded`. For example:

```
gocompat compare --go1compat v1.0.0 v1.1.0 ./...
```

## License

Released under the terms of the Apache License Version 2.0, see [LICENSE](LICENSE).