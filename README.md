# nosharedparamtype

A Go linter that enforces explicit type declarations for all function parameters.

## Overview

In Go, multiple consecutive parameters of the same type can share a single type
declaration. While this is valid syntax, it can reduce code clarity and
therefore readability.

## Examples

### Incorrect code examples

```go
func add(x, y int) int {
    return x + y
}

func calculate(a, b int, c float64, d, e string) float64 {
    // ...
}

func (s *Server) handle(req, resp http.Request) error {
    // ...
}
```

### Correct code

```go
func add(x int, y int) int {
    return x + y
}

func calculate(a int, b int, c float64, d string, e string) float64 {
    // ...
}

func (s *Server) handle(req http.Request, resp http.Request) error {
    // ...
}
```
