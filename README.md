# stacks

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/stacks.svg)](https://pkg.go.dev/cattlecloud.net/go/stacks)
[![License](https://img.shields.io/github/license/cattlecloud/stacks?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/stacks/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/stacks/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/stacks/actions/workflows/ci.yaml)

`stacks` is a Go library with generics implementing multiple Stack data structures
for optimizing by use-case.

```
stack (noun):
  1. a LIFO data structure
```

### Getting Started

The `stacks` package can be added to a Go project with `go get`.

```shell
go get cattlecloud.net/go/stacks@latest
```

```go
import "cattlecloud.net/go/stacks"
```

### Examples

##### backed by slice

The `Simple` stack type is backed by a normal Go slice and is a good choice for
small numbers of elements or where the size of the stack will not increase or
decrease repeatedly - this is important so as to minmize re-copying the data as
the underlying slice is resized.

```go
s := stacks.Simple[string]()
s.Push("how")
s.Push("are")
s.Push("you")
fmt.Println(s.Pop()) // "you"
```

##### backed by linked list

The `Linked` stack type is backed by a linked-list and is a good choice for large
numbers of elements or where the size of the stack does increase or decrease
repeatedly. The nature of the linked list avoids the need to re-copy elements as
the size of the data structure changes over time.

```go
s := stacks.Linked[string]()
s.Push("foo")
s.Push("bar")
s.Push("baz")
fmt.Println(s.Pop()) // "baz"
```

##### methods

```go
s.Push[T](item T)
s.Pop() T
s.Peek() T
s.Empty() bool
s.Size() int
```

### License

The `cattlecloud.net/go/stacks` module is opensource under the [BSD-3-Clause](LICENSE) license.
