// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

// Package lispylist implements a classical lisp list data structure based on a Cons cell or Pair.
// Unlike languages such as Python and Ruby, Go does not have a built in list type.
// There is a List container but it's rarely used and just doubly linked list.
//
// The basic data structure here is a "List". sometimes called a "cons" or "pair" which has two cells, a Head and a Tail.
// In classic lisp the head of a list is the "car" and the tail is the "cdr".
// A Pair is often written as a dotted Pair "(a . b)" where a is the Head and b us the Tail.
// Lists are build up from dotted Pairs via Cons, with the last element of the list having nil as it's Tail.
// MakeList(1, 2, 3, 4) === (1 . (2 . (3 . (4. nil)))) == Cons(1, Cons(2, Cons(3, Cons(4, nil))))
// In summary, list.Head contains a list element and the list.Tail contains a pointer to the next list element.
// Most of the functions here are straight functions and not methods.
// That's on purpose to support functional composition that reads well.
// Currently the sole exception is Print for which a convenience function Print exists.
//
// I am a bit conflicted about using classic Lisp names like car and cdr vs head and tail, nconc vs splice.
//
// All code has been tested with lists up to length of 100,000.
// For lists larger than 100,000 all the functions that are tail recursive would have to be verified.
//
// NB. There is code below that does not check that a type assertion is "ok".
// I am aware of this. However, even though it's not "ok" the nil is still assigned to the variable on the lhs.
// I actually coded Length with the check and it is substantially harder to read the code and the code is not
// really more robust with the check.

package lispylist

import (
	"fmt"
)

// List is the basic data structure used for all lists.
// I would have liked to call this Pair and make a type
// type List *Pair
// However, methods can not have pointer types a receivers.
// Might revisit this decision
type Pair struct {
	Head interface{}
	Tail interface{}
}
type List = *Pair

var NilList List = nil

// Cons creates a single element of a List
// Cons(1, nil) creates a single element list
func Cons(a, b interface{}) List {
	//fmt.Printf("cons: (%v . %v)\n", a, b)
	return &Pair{a, b}
}

// List creates a list of whatever is passed in as rest.
func MakeList(rest ...interface{}) List {
	if rest == nil || len(rest) == 0 {
		return nil
	}
	h := Cons(rest[0], nil)
	l := h
	for _, v := range rest[1:] {
		l.Tail = Cons(v, nil)
		l = l.Tail.(List)
	}
	return h
}

// Last returns the last element of a list as a list.
func Last(x List) List {
	for x.Tail != nil {
		x = x.Tail.(List)
	}
	return x
}

// Nconc is like append but it shares structure with the lists that are passed to it.
func Nconc(x, y List) List {
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
func Flatten(tree interface{}) List {
	var rec func(x interface{}, acc List) List
	rec = func(x interface{}, acc List) List {
		if x == nil {
			return acc
		}
		switch v := x.(type) {
		case List:
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

// FlattenAlt is an alternate version of Flatten.
func FlattenAlt(l List) List {
	if l == nil {
		return nil
	}
	Head, Headok := l.Head.(List)
	Tail, _ := l.Tail.(List)
	if Headok {
		return Nconc(FlattenAlt(Head), FlattenAlt(Tail))
	}
	return Nconc(Cons(l.Head, nil), FlattenAlt(Tail))
}

// Length as a tail recursive function.
func LengthAlt(l List) int {
	if l == nil {
		return 0
	}
	l, _ = l.Tail.(List)
	return (1 + LengthAlt(l))
}

// Length returns the number of top level Cons.
func Length(l List) int {
	var cnt int
	for {
		if l == nil {
			return cnt
		}
		l, _ = l.Tail.(List)
		cnt++
	}
}

// Compare to above, is this really better and safer?
func LengthWithCheck(l List) int {
	var cnt int
	var ok bool
	if l == nil {
		return cnt
	}
	for {
		cnt++
		if l.Tail == nil {
			return cnt
		}
		l, ok = l.Tail.(List)
		if !ok {
			panic("Length: bad list structure")
		}
	}
}

// Traverse a list depth first/head first
func Traverse(lst List, f func(e interface{})) {
	var trav func(l List, f func(e interface{}))
	trav = func(l List, f func(e interface{})) {
		if l == nil {
			return
		}
		if l.Head != nil {
			v, ok := l.Head.(List)
			if ok {
				trav(v, f)
			} else {
				f(l.Head)
			}
		}
		if l.Tail != nil {
			v, ok := l.Tail.(List)
			if ok {
				trav(v, f)
			} else {
				f(l.Tail)
			}
		}
	}
	trav(lst, f)
}

var Pmax = 45 // The maximum number of list elements Print will print.

// Print prints a list up to Pmax elements long
// would be nice if pmax triggers to print the last few elements too
func Print(list interface{}) {
	var cnt = 0
	var prt = func(x interface{}) {
		_, ok := x.(string)
		if ok {
			fmt.Printf("%q", x)
		} else {
			fmt.Printf("%v", x)
		}
	}
	var p func(list interface{})
	p = func(alist interface{}) {
		list, ok := alist.(List)
		if !ok {
			prt(alist)
			return
		}
		fmt.Printf("(")
		prev := list
		for {
			if cnt >= Pmax {
				fmt.Printf("...")
				break
			}
			if prev == nil {
				break
			}
			if prev.Head != nil {
				h, ok := prev.Head.(List)
				if ok {
					p(h)
				} else {
					prt(prev.Head)
					cnt++
				}
			}
			if prev.Tail != nil {
				prev = prev.Tail.(List)
				fmt.Printf(" ")
			} else {
				break
			}
		}
		fmt.Printf(")")
	}
	//fmt.Printf("%v\n", list)
	p(list)
	fmt.Printf("\n")
}

func print(list List) {
	fmt.Printf("%v\n", list)
}
