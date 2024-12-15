stacks
======

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/stacks.svg)](https://pkg.go.dev/cattlecloud.net/go/stacks)
[![License](https://img.shields.io/github/license/cattlecloud/stacks?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/stacks/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/stacks/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/stacks/actions/workflows/ci.yaml)

`stacks` provides multiple implementations of a generic Stack data structure for
optimizing by varioius use cases.

```
stack (noun):
  1. a LIFO data structure
```

# Getting Started

The `stacks` package can be added to a Go project with `go get`.

```shell
go get cattlecloud.net/go/stacks@latest
```

# Examples

##### Create a contiguous stack.

The `Simple` constructor can be used to create `Stack` backed by a `[]T`. This
is useful for most use cases where the size of the stack will not be problematic
in normal circumstances. If you are storing small types such as numeric types or
pointers, this is probably the best type of stack to use.

```go
s := stacks.Simple[*person]()
s.Push(&person{Name: "Alice"})
s.Push(&person{Name: "Bob"})
```

##### Create a linked-list stack.

The `Linked` constructor can be used to create a `Stack` backed by a linked list
of `T` elements. This is usefor for cases where the `T` type is large and would
not be reasonable to maintain a large contigous slice of `T` elements.

```go
s := stacks.Linked[blob]()
s.Push(blob{...})
```

##### Pop elements.

Elements can be popped off the top of the stack. It is a logic bug to pop an empty
stack, so it is common to check if the stack is empty first.

```go
empty := s.Empty()
```

```go
item := s.Pop()
```

##### Peek at top element.

You can get a reference to the top element without removing it. Again, calling
peek on an empty stack is a logic bug.

```go
item := s.Peek()
```

##### Check size of the stack.

The stack can report the number of elements it contains.

```go
size := s.Size()
```

# License

The `cattlecloud.net/go/stacks` module is opensource under the [BSD-3-Clause](LICENSE) license.
