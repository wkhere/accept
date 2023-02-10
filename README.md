## accept

Check the Accept-* http headers if the given value is accepted.

Example:
```go
if accept.Encoding(req.Header.Get("Accept-Encoding"), "deflate") {
    w.Header().Set("Content-Encoding", "deflate")
    // go with the deflate compression..
}
```
