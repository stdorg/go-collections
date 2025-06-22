# go-collections

A generic, efficient, and thread-safe collections library for Go. Designed to provide a variety of data structures with both concurrent-safe and non-concurrent implementations using Go generics.

## Features

- **Generic Interfaces**: Type-safe APIs for any comparable type.
- **Concurrent-Safe Implementations**: Safe for use in multi-goroutine environments.
- **Non-Concurrent Implementations**: High performance for single-threaded use cases.
- **Convenient API**: Common operations such as Add, Remove, Contains, Clear, Len, Chan, Slice, and more.

## Installation

```sh
go get github.com/stdorg/go-collections
```

## Usage Examples

### Importing a Collection

```go
import "github.com/stdorg/go-collections/<collection>"
```

### Example: Concurrent-Safe Collection

```go
c := <collection>.NewSafe[Type]()

// ...
```

### Example: Non-Concurrent Collection

```go
c := <collection>.NewUnsafe[Type]()

// ...
```

## Thread Safety

- Use `NewSafe` when multiple goroutines need to access or modify the collection concurrently. These implementations use synchronization primitives to ensure thread safety.
- Use `NewUnsafe` for single-threaded or read-only scenarios where thread safety is not required. These versions offer better performance by avoiding synchronization overhead.

## License

MIT License

Copyright &copy; 2024 Othon Hugo ([github.com/othonhugo](https://github.com/othonhugo))