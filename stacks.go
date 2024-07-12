// Copyright (c) The Noxide Project Authors
// SPDX-License-Identifier: BSD-3-Clause

// Package stacks provides various implementations of a stack data structure,
// each ideal for a certain set of use cases and conditions.
package stacks

// A Stack implements the basic stack data structure operations, including Push,
// Pop, Peek, Empty, and Size.
type Stack[T any] interface {
	Push(T)
	Pop() T
	Peek() T
	Empty() bool
	Size() int
}

// Simple creates an implementation of Stack backed by a slice. This is a good
// choice for small numbers of elements or where the size of the stack will not
// increase or decrease substantially.
func Simple[T any](items ...T) Stack[T] {
	ss := &sliceStack[T]{items: make([]T, 0, len(items))}
	initialize(ss, items...)
	return ss
}

// Linked creates an implementation of Stack backed by a linked-list. This is
// a good choice for use cases where the size of the stack changes dramatically
// and often. Unlike a slice, the backing data structure will not need to be
// copied as the size of the stack changes, as each element is allocated
// individually in the linked list.
func Linked[T any](items ...T) Stack[T] {
	ls := &linkStack[T]{
		top:  nil,
		size: 0,
	}
	initialize(ls, items...)
	return ls
}
