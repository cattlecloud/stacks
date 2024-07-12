// Copyright (c) The Noxide Project Authors
// SPDX-License-Identifier: BSD-3-Clause

package stacks

// verify ensures the stack is not empty, and will panic if it is
func check[T any](s Stack[T]) {
	if s.Empty() {
		panic("stacks: is empty")
	}
}

// initialize pushes each item in items onto the stack
func initialize[T any](s Stack[T], items ...T) {
	for _, item := range items {
		s.Push(item)
	}
}
