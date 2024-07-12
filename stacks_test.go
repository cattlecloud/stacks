// Copyright (c) The Noxide Project Authors
// SPDX-License-Identifier: BSD-3-Clause

package stacks

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestBasic(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		basicTest(t, Simple[int]())
	})

	t.Run("linked", func(t *testing.T) {
		basicTest(t, Linked[int]())
	})
}

func basicTest(t *testing.T, s Stack[int]) {
	must.Empty(t, s)
	must.Size(t, 0, s)

	s.Push(1)

	must.NotEmpty(t, s)
	must.Eq(t, 1, s.Peek())
	must.Size(t, 1, s)

	s.Push(2)
	must.Eq(t, 2, s.Peek())
	must.Size(t, 2, s)

	s.Push(3)
	must.Eq(t, 3, s.Peek())
	must.Size(t, 3, s)

	must.Eq(t, 3, s.Pop())
	must.Size(t, 2, s)
	must.Eq(t, 2, s.Peek())

	must.Eq(t, 2, s.Pop())
	must.Size(t, 1, s)
	must.Eq(t, 1, s.Peek())

	s.Push(4)
	s.Push(5)

	must.Eq(t, 5, s.Peek())
	must.Eq(t, 5, s.Pop())
	must.Eq(t, 4, s.Peek())

	must.Eq(t, 4, s.Pop())
	must.Eq(t, 1, s.Pop())

	must.Empty(t, s)
	must.Eq(t, 0, s.Size())
}

func TestPopEmpty(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		s := Simple[int]()
		emptyTest(t, s, s.Pop)
	})

	t.Run("linked", func(t *testing.T) {
		s := Linked[int]()
		emptyTest(t, s, s.Pop)
	})
}

func TestPeekEmpty(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		s := Simple[int]()
		emptyTest(t, s, s.Peek)
	})

	t.Run("linked", func(t *testing.T) {
		s := Linked[int]()
		emptyTest(t, s, s.Peek)
	})
}

func emptyTest[T any](t *testing.T, s Stack[int], f func() T) {
	must.Empty(t, s)
	var ok = false

	defer func() {
		if r := recover(); r != nil {
			ok = true
			s := r.(string)
			must.Eq(t, "stacks: is empty", s)
		}
	}()

	f()
	must.True(t, ok)
}

func TestInitial(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		s := Simple(1, 2, 3)
		initTest(t, s)
	})

	t.Run("linked", func(t *testing.T) {
		s := Linked(1, 2, 3)
		initTest(t, s)
	})
}

func initTest(t *testing.T, s Stack[int]) {
	must.Size(t, 3, s)
	must.Eq(t, 3, s.Pop())
	must.Eq(t, 2, s.Pop())
	must.Eq(t, 1, s.Pop())
	must.Empty(t, s)
}
