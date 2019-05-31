// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

// Package lisplist implements a clasical lisp list datastructure.
// Lists are build up from pairs via cons.

package main

import "fmt"

// Pair is the basic data structure used for all lists.
type pair struct {
	head interface{}
	tail interface{}
}

// Cons creates a pair
func cons(a, b interface{}) *pair {
	return &pair{a, b}
}

// Last returns the last element of a list as a list.
func last(x *pair) *pair {
	for x.tail != nil {
		x = x.tail.(*pair)
	}
	return x
}

// nconc is like append but it shares structure with the lists that are passed to it.
func nconc(x, y *pair) *pair {
	//fmt.Printf("nconc: ")
	//fmt.Printf("x=%#v", x)
	//fmt.Printf("y=%#v\n", y)
	switch {
	case x == nil && y == nil:
		return nil
	case x != nil && y == nil:
		return x
	case x == nil && y != nil:
		return y
	case x.tail != nil:
		//fmt.Printf("x tail\n")
		l := last(x)
		l.tail = y
		return x
	default:
		x.tail = y
		return x
	}
}

/*
(defun flatten (tree)
  (if (atom tree)
      (mklist tree)
    (nconc (flatten (car tree))
       (if (cdr tree) (flatten (cdr tree))))))
*/

// Flatten takes a nested list and flattens it.
func flatten(tree *pair) *pair {
	if tree == nil {
		return nil
	}
	car, carok := tree.head.(*pair)
	cdr, _ := tree.tail.(*pair)
	if carok {
		//fmt.Printf("carok\n")
		return nconc(flatten(car), flatten(cdr))
	}
	//fmt.Printf("number\n")
	return nconc(cons(tree.head, nil), flatten(cdr))
}

// List creates a list of whatever is passed in as rest.
func list(rest ...interface{}) *pair {
	h := cons(rest[0], nil)
	l := h
	//fmt.Printf("%#v\n", l)
	for _, v := range rest[1:] {
		l.tail = cons(v, nil)
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
			if prev.head != nil {
				h, ok := prev.head.(*pair)
				if ok {
					p(h)
				} else {
					fmt.Printf("%v", prev.head)
				}
			}
			if prev.tail != nil {
				prev = prev.tail.(*pair)
				fmt.Printf(" ")
			} else {
				fmt.Printf(")")
				return
			}
		}
	}
	p(list)
	fmt.Printf("\n")
}

func print(list *pair) {
	list.Print()
}

func main() {
	h := cons(1, nil)
	l := h
	l.tail = cons(2, nil)
	l = l.tail.(*pair)
	l.tail = cons(3, nil)
	l1 := list(1) //, 2)
	l2 := list(4) //, 5, 6)
	l3 := list(1, list(2, 3, list(4, 5, 6)), list(7, 8))
	//b := cons(2, a.tail)
	//cons(3, b.tail)
	print(h)
	print(l1)
	print(l2)
	print(last(l2))
	print(l2)
	print(nconc(l1, l2))
	print(l3)
	//print(nconc(l1, l2))
	print(l1)
	print(l2)
	//fmt.Printf("%#v\n", l2)
	f := flatten(l3)
	//fmt.Printf("%#v\n", f)
	//fmt.Printf("%#v\n", f.tail)
	print(f)
}
