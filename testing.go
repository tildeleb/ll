// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

package lispylist

import (
	_ "fmt"
	"math/rand"
	"time"
)

// Support for random numbers without locking
var s = rand.NewSource(time.Now().UTC().UnixNano())
var r = rand.New(s)

// rbetween returns random int [a, b].
func rbetween(a int, b int) int {
	return r.Intn(b-a+1) + a
}

// A few utility functions mainly used for testing.
func VerifyFlattenedList(l List) bool {
	i := 1
	var v func(l List) bool
	v = func(l List) bool {
		if l.Head.(int) != i {
			return true
		}
		if l.Tail == nil {
			return false
		}
		i++
		return v(l.Tail.(List))
	}
	return v(l)
}

// Generate a list of ints starting at start and length long.
func GenIntList(start, length int) List {
	idx := start + length - 1
	l := Cons(idx, nil)
	idx--
	for idx >= start {
		l = Cons(idx, l)
		idx--
	}
	return l
}

// Generate a nested list of ints starting at start and length long with up to depth nesting.
func GenNestedList(astart, length, depth int) List {
	var start = astart
	var lst List = nil
	var gnl func(d int) List
	gnl = func(d int) List {
		d--
		if d > 0 {
			l := Cons(nil, nil)
			a := gnl(rbetween(1, d))
			b := gnl(rbetween(1, d))
			l.Head = a
			l.Tail = MakeList(b)
			return l
		} else {
			n := rbetween(1, length/rbetween(1, depth)+1)
			l := GenIntList(start, n)
			start += n
			return l
		}
	}
	for start < length+astart {
		x := rbetween(1, 5)
		if x < 3 {
			n := rbetween(start, length+astart)
			lst = Splice(lst, GenIntList(start, n))
			start += n
		} else {
			d := rbetween(1, depth)
			l := gnl(d)
			lst = Splice(lst, l)
		}
	}
	return lst
}
