// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

// Package lisplist implements a clasical lisp list datastructure.
// Unlike languages such as Python and Ruby Go doesn't have a built in list type.
// There is a list container but it's rarely used and just doubly linked list.
//
// The basic datastructure here is a "pair" which has two cells a head and a tail.
// A pair is often written as a dotted pair "(a . b)" where a is the head and b us the tail.
// Lists are build up from dotted pairs via cons.
// List(1, 2, 3, 4) === (1 . (2 . (3 . (4. nil))))
// So the head contains the list element and the tail contains a pointer to the next pair
// The list is terminated by nil as the tail.
// This code demonstrates:
// 1. Ability to not only write flatten but design a list data structure
// 2. Knowledge of creating a generic datastructure
// 3. Knowledge of the 2 nils issue when using interface{}
// 4. Use of recursive tail calls which Go usually optimizes (not verified)
// 5. Use of varagrs/rest argument

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Pair is the basic data structure used for all lists.
type pair struct {
	head interface{}
	tail interface{}
}

// Cons creates a pair
func Cons(a, b interface{}) *pair {
	//fmt.Printf("cons: (%v . %v)\n", a, b)
	return &pair{a, b}
}

// Last returns the last element of a list as a list.
func Last(x *pair) *pair {
	for x.tail != nil {
		x = x.tail.(*pair)
	}
	return x
}

// Nconc is like append but it shares structure with the lists that are passed to it.
func Nconc(x, y *pair) *pair {
	switch {
	case x == nil && y == nil:
		return nil
	case x != nil && y == nil:
		return x
	case x == nil && y != nil:
		return y
	case x.tail != nil:
		l := Last(x)
		l.tail = y
		return x
	default:
		x.tail = y
		return x
	}
}

// flatten2 takes a nested list and flattens it.
func flatten2(tree interface{}) *pair {
	var rec func(x interface{}, acc *pair) *pair
	rec = func(x interface{}, acc *pair) *pair {
		if x == nil {
			return acc
		}
		switch v := x.(type) {
		case *pair:
			return rec(v.head, rec(v.tail, acc))
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
func flatten(tree *pair) *pair {
	if tree == nil {
		return nil
	}
	head, headok := tree.head.(*pair)
	tail, _ := tree.tail.(*pair)
	if headok {
		return Nconc(flatten(head), flatten(tail))
	}
	return Nconc(Cons(tree.head, nil), flatten(tail))
}

// List creates a list of whatever is passed in as rest.
func List(rest ...interface{}) *pair {
	h := Cons(rest[0], nil)
	l := h
	for _, v := range rest[1:] {
		l.tail = Cons(v, nil)
		l = l.tail.(*pair)
	}
	return h
}

// Print prints a list
func (list *pair) Print() {
	var p func(list *pair)
	p = func(list *pair) {
		fmt.Printf("(")
		prev := list
		for {
			if prev == nil {
				break
			}
			if prev.head != nil {
				h, ok := prev.head.(*pair)
				if ok {
					p(h)
				} else {
					_, ok := prev.head.(string)
					if ok {
						fmt.Printf("%q", prev.head)
					} else {
						fmt.Printf("%v", prev.head)
					}
				}
			}
			if prev.tail != nil {
				prev = prev.tail.(*pair)
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

func Print(list *pair) {
	list.Print()
}

var s = rand.NewSource(time.Now().UTC().UnixNano())
var r = rand.New(s)

// rbetween returns random int [a, b]
func rbetween(a int, b int) int {
	return r.Intn(b-a+1) + a
}

func genList(start, length int) *pair {
	idx := start + length - 1
	l := Cons(idx, nil)
	idx--
	for idx >= start {
		l = Cons(idx, l)
		idx--
	}
	return l
}

func genNestedList(v, n, depth int) *pair {
	var gnl func(d int) *pair
	gnl = func(d int) *pair {
		d--
		if d > 0 {
			lst := Cons(nil, nil)
			a := gnl(rbetween(1, d))
			b := gnl(rbetween(1, d))
			lst.head = a
			lst.tail = List(b)
			return lst
		} else {
			l := rbetween(1, n/rbetween(1, depth))
			lst := genList(v, l)
			v += l
			return lst
		}
	}
	d := rbetween(1, depth)
	return gnl(d)
}

func main() {
	lst := genNestedList(1, 20, 5)
	Print(lst)
	Print(flatten2(lst))
	return
	Print(flatten2(List(1, 2, 3)))
	Print(genList(1, 3))
	h := Cons(1, nil)
	l := h
	l.tail = Cons(2, nil)
	l = l.tail.(*pair)
	l.tail = Cons(3, nil)
	l1 := List(1)
	l2 := List(2)
	l3 := List(3, 4)
	l4 := List(5, 6, 7)
	s := []int{1, 2, 3, 4, 5, 6}
	l5 := List(s)
	l6 := List("how", "now", "brown", "cow")
	l7 := List(1, List(2, 3, List(4, 5, 6)), List(7, 8))
	l8 := List(List(List(List(1, 2), 3), 4), List(5, 6), List(7, 8, 9), List(10, 11))
	Print(h)
	Print(l1)
	Print(l2)
	Print(Last(l2))
	Print(l2)
	Print(Nconc(l1, l2))
	Print(l3)
	Print(l4)
	Print(Nconc(l3, l4))
	Print(l5)
	Print(l6)
	Print(l7)
	fmt.Print("14: ")
	Print(l4)
	Print(flatten2(l4))
	fmt.Print("17: ")
	Print(l7)
	Print(flatten2(l7))
	f := flatten2(l8)
	fmt.Print("18: ")
	Print(f)
	Print(nil)
}
