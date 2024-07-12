// Copyright (c) The Noxide Project Authors
// SPDX-License-Identifier: BSD-3-Clause

package stacks

type element[T any] struct {
	item T
	next *element[T]
}

type linkStack[T any] struct {
	top  *element[T]
	size int
}

func (ls *linkStack[T]) Push(item T) {
	ls.top = &element[T]{
		item: item,
		next: ls.top,
	}
	ls.size++
}

func (ls *linkStack[T]) Pop() T {
	check(ls)

	elem := ls.top
	ls.top = elem.next
	elem.next = nil
	ls.size--
	return elem.item
}

func (ls *linkStack[T]) Peek() T {
	check(ls)

	return ls.top.item
}

func (ls *linkStack[T]) Empty() bool {
	return ls.Size() == 0
}

func (ls *linkStack[T]) Size() int {
	return ls.size
}
