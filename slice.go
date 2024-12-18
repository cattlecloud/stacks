// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package stacks

type sliceStack[T any] struct {
	items []T
}

func (ss *sliceStack[T]) Push(item T) {
	ss.items = append(ss.items, item)
}

func (ss *sliceStack[T]) Pop() T {
	check(ss)

	item := ss.items[len(ss.items)-1]
	ss.items = ss.items[0 : len(ss.items)-1]
	return item
}

func (ss *sliceStack[T]) Peek() T {
	check(ss)

	return ss.items[len(ss.items)-1]
}

func (ss *sliceStack[T]) Empty() bool {
	return ss.Size() == 0
}

func (ss *sliceStack[T]) Size() int {
	return len(ss.items)
}
