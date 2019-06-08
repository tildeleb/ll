// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

// Package ll (lispylist) implements a classical lisp list data structure based on a Cons cell or Pair.
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
//
// This package uses the classic Lisp names like car, cdr, nconc, and so on.
// My take on more modern names is in modern.go
//
// All code has been tested with lists up to length of 100,000.
// Many of the functions here are tail recursive in the original style of Lisp.
// For lists larger than 100,000 all the functions that are tail recursive optimization would have to be verified.
//
// NB. There is code below that does not check that a type assertion is "ok".
// I am aware of this. However, even though it's not "ok" the nil is still assigned to the variable on the lhs.
// I actually coded Length with the check and it is substantially harder to read the code and the code is not
// really more robust with the ok check.

package ll

import (
	_ "fmt"
)

// Pair is the basic data structure used for all lists.
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

func Ncons(a interface{}) List {
	return &Pair{a, nil}
}

func Xcons(a, b interface{}) List {
	return &Pair{b, a}
}

// New creates a list of whatever is passed in as rest.
// In Lisp this would be called List but that name is used for the type here.
func New(rest ...interface{}) List {
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

//(defun evenp (n) (if (zerop n) t (oddp (1- n)))) (defun oddp (n) (evenp (1- n)))

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

func Append2(x, y List) List {
	if x == nil {
		return y
	}
	cdr, _ := Cdr(x).(List)
	return Cons(Car(x), Append2(cdr, y))
}

func Append(rest ...interface{}) List {
	if rest == nil || len(rest) == 0 {
		return nil
	}
	if len(rest) == 1 {
		return rest[0].(List)
	}
	length := len(rest)
	length--
	val := rest[length].(List)
	rest = rest[:length]
	length--
	if length >= 0 {
		for i, _ := range rest {
			r, _ := rest[length-i].(List)
			val = Append2(r, val)
		}
	}
	return val
}

func Reverse(x List) List {
	l := x
	r := NilList
	for {
		r = Cons(Car(l), r)
		l, _ = Cdr(l).(List)
		if l == nil {
			return r
		}
	}
}

func substr(x, y, z interface{}) interface{} {
	if z == nil {
		return nil
	}
	lst, ok := z.(List)
	if !ok {
		if y == z {
			return x
		} else {
			return z
		}
	}
	// z is a list so we can take the car and cdr
	// double tail recursive
	return Cons(substr(x, y, Car(lst)), substr(x, y, Cdr(lst)))
}

// Substr substitutes x for all instances of y in list z.
func Substr(x, y interface{}, z List) List {
	l, _ := substr(x, y, z).(List)
	return l
}

func Member(x interface{}, y List) List {
	//fmt.Printf("x=%v, Car(y)=%v\n", x, Car(y))
	if y == nil {
		return nil
	}
	if x == Car(y) {
		return y
	}
	cdr, _ := Cdr(y).(List)
	return Member(x, cdr)
}
