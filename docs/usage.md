# Understanding `*http.Request` in Go

In Go, we use a pointer to access `*http.Request` for several reasons:

## 1. Efficiency
- **Avoids Copying**: `http.Request` is a large struct with many fields. Passing it by value (i.e., not using a pointer) would involve copying the entire struct, which can be inefficient in terms of memory and performance. By passing a pointer (`*http.Request`), you avoid copying the struct and only pass a reference to it, which is more efficient.

## 2. Mutability
- **Modification**: The `http.Request` struct contains information about the incoming request, including headers, form values, and URL parameters. Although in most cases, you don't modify the request itself, some parts of the request, like form values, might be parsed or altered. Using a pointer allows you to modify the request's fields if needed.

## 3. Consistency with Go's Idioms
- **Go Convention**: It is idiomatic in Go to use pointers for large structs or when you need to modify the struct. This practice is consistent across many Go packages and standard library functions.

### Example

Consider the `http.Request` struct:

```go
type Request struct {
    Method string
    URL *url.URL
    Header Header
    Body io.ReadCloser
    // other fields
}
