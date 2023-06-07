[![Build Status](https://github.com/wkhere/accept/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/wkhere/accept/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/wkhere/accept/badge.svg?branch=master)](https://coveralls.io/github/wkhere/accept?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/wkhere/accept.svg)](https://pkg.go.dev/github.com/wkhere/accept)

## accept

Check the Accept-* http headers if the given value is accepted.

Example:
```go
if accept.Encoding(req.Header.Get("Accept-Encoding"), "deflate") {
    w.Header().Set("Content-Encoding", "deflate")
    // go with the deflate compression..
}
```
