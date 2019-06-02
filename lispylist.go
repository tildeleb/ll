// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

// Package lisplist implements a clasical lisp list datastructure based on a Cons.
// Unlike languages such as Python and Ruby, Go does not have a built in list type.
// There is a List container but it's rarely used and just doubly linked list.
//
// The basic datastructure here is a "List". sometimes called a "cons" which has two cells, a Head and a Tail.
// In classic list the head of a list is the "car" and the tail is the "cdr".
// A Pair is often written as a dotted Pair "(a . b)" where a is the Head and b us the Tail.
// Lists are build up from dotted Pairs via cons, with the last element of the list having nil as it's Tail.
// MakeList(1, 2, 3, 4) === (1 . (2 . (3 . (4. nil))))
// So the Head contains the list element and the Tail contains a pointer to the next list element or the rest of the list.
// Most of the functions here are just straight functions and not methods.
// Currently the sole exception is Print.

package lispylist

import (
	"fmt"
)

// List is the basic data structure used for all lists.
// I would have liked to call this Pair and make a type
// type List *Pair
// I can't do that because methods can not have pointer types a receivers.
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

// flatten2 takes a nested list and flattens it.
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

// flatten takes a nested list and flattens it.
func FlattenSlow(tree *List) *List {
	if tree == nil {
		return nil
	}
	Head, Headok := tree.Head.(*List)
	Tail, _ := tree.Tail.(*List)
	if Headok {
		return Nconc(FlattenSlow(Head), FlattenSlow(Tail))
	}
	return Nconc(Cons(tree.Head, nil), FlattenSlow(Tail))
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
