// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

// Package lispylist implements a classical lisp list data structure based on a Cons cell or Pair.
// Unlike languages such as Python and Ruby, Go does not have a built in list type.
// There is a List container but it's rarely used and just doubly linked list.
//
// The basic data structure here is a "List". sometimes called a "cons" or "pair" which has two cells, a Head and a Tail.
// In classic lisp the head of a list is the "car" and the tail is the "cdr".
// A Pair is often written as a dotted Pair "(a . b)" where a is the Head and b us the Tail.
// Lists are build up from dotted Pairs via cons, with the last element of the list having nil as it's Tail.
// MakeList(1, 2, 3, 4) === (1 . (2 . (3 . (4. nil)))) == Cons(1, Cons(2, Cons(3, Cons(4, nil))))
// In summary, list.Head contains a list element and the list.Tail contains a pointer to the next list element.
// Most of the functions here are just straight functions and not methods.
// Currently the sole exception is Print for which a convenience

package lispylist

import (
	"fmt"
	"math/rand"
	"time"
)

// List is the basic data structure used for all lists.
// I would have liked to call this Pair and make a type
// type List *Pair
// However, methods can not have pointer types a receivers.
// Might revisit this decision
type List struct {
	Head interface{}
	Tail interface{}
}

// Cons creates a single element of a List
// Cons(1, nil) creates a single element list
func Cons(a, b interface{}) *List {
	//fmt.Printf("cons: (%v . %v)\n", a, b)
	return &List{a, b}
}

// List creates a list of whatever is passed in as rest.
func MakeList(rest ...interface{}) *List {
	h := Cons(rest[0], nil)
	l := h
	for _, v := range rest[1:] {
		l.Tail = Cons(v, nil)
		l = l.Tail.(*List)
	}
	return h
}

// Last returns the last element of a list as a list.
func Last(x *List) *List {
	for x.Tail != nil {
		x = x.Tail.(*List)
	}
	return x
}

// Nconc is like append but it shares structure with the lists that are passed to it.
func Nconc(x, y *List) *List {
	switch {
	case x == nil && y == nil:
		return nil
	case x != nil && y == nil:
		return x
	case x == nil && y != nil:
		return y
	case x.Tail != nil:
		l := Last(x)
		l.Tail = y
		return x
	default:
		x.Tail = y
		return x
	}
}

// Flatten takes a nested list and flattens it. This version uses an accumulator.
func Flatten(tree interface{}) *List {
	var rec func(x interface{}, acc *List) *List
	rec = func(x interface{}, acc *List) *List {
		if x == nil {
			return acc
		}
		switch v := x.(type) {
		case *List:
			return rec(v.Head, rec(v.Tail, acc))
		default:
			if acc == nil {
				return Cons(v, nil)
			}
			return Cons(v, acc)
		}
	}
	return rec(tree, nil)
}

// FlattenAlt is an alternate version of Flatten
func FlattenAlt(tree *List) *List {
	if tree == nil {
		return nil
	}
	Head, Headok := tree.Head.(*List)
	Tail, _ := tree.Tail.(*List)
	if Headok {
		return Nconc(FlattenAlt(Head), FlattenAlt(Tail))
	}
	return Nconc(Cons(tree.Head, nil), FlattenAlt(Tail))
}

// Print prints a list
func (list *List) Print() {
	var p func(list *List)
	p = func(list *List) {
		fmt.Printf("(")
		prev := list
		for {
			if prev == nil {
				break
			}
			if prev.Head != nil {
				h, ok := prev.Head.(*List)
				if ok {
					p(h)
				} else {
					_, ok := prev.Head.(string)
					if ok {
						fmt.Printf("%q", prev.Head)
					} else {
						fmt.Printf("%v", prev.Head)
					}
				}
			}
			if prev.Tail != nil {
				prev = prev.Tail.(*List)
				fmt.Printf(" ")
			} else {
				break
			}
		}
		fmt.Printf(")")
	}
	p(list)
	fmt.Printf("\n")
}

func Print(list *List) {
	list.Print()
}

// A few utility functions mainly used for testing
var s = rand.NewSource(time.Now().UTC().UnixNano())
var r = rand.New(s)

// rbetween returns random int [a, b]
func rbetween(a int, b int) int {
	return r.Intn(b-a+1) + a
}

func VerifyFlattenedList(l *List) bool {
	i := 1
	var v func(l *List) bool
	v = func(l *List) bool {
		//fmt.Printf("%v\n", l)
		//#fmt.Printf("%v == %v\n", l.Head.(int), i)
		if l.Head.(int) != i {
			return true
		}
		if l.Tail == nil {
			return false
		}
		i++
		return v(l.Tail.(*List))
	}
	return v(l)
}

func GenIntList(start, length int) *List {
	idx := start + length - 1
	l := Cons(idx, nil)
	idx--
	for idx >= start {
		l = Cons(idx, l)
		idx--
	}
	return l
}

func GenNestedList(v, n, depth int) *List {
	var gnl func(d int) *List
	gnl = func(d int) *List {
		d--
		if d > 0 {
			lst := Cons(nil, nil)
			a := gnl(rbetween(1, d))
			b := gnl(rbetween(1, d))
			lst.Head = a
			lst.Tail = MakeList(b)
			return lst
		} else {
			l := rbetween(1, n/rbetween(1, depth))
			lst := GenIntList(v, l)
			v += l
			return lst
		}
	}
	d := rbetween(1, depth)
	return gnl(d)
}
