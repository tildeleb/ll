# ll (lispy lists in Go)

Copyright © 2019 Lawrence E. Bakst. All rights reserved.

## Introduction

This one was just for fun.

Unlike languages such as Python and Ruby, Go does not have a built in list type.

Package ll implements a classical lisp list data structure based on a Cons cell or Pair. 

The basic data structure here is a "List". sometimes called a "cons" or "pair" which has two cells, a Head and a Tail. In classic lisp the Head of a list is accessed via the "car" fucntion and the Tail is accessed via the "cdr" function. A Pair is often written as a dotted pair "(a . b)" where a is the Head and b is the Tail. Lists are build up from dotted Pairs via cons, with the last element of the list having nil as it's Tail.

    MakeList(1, 2, 3, 4) === (1 . (2 . (3 . (4. nil)))) == Cons(1, Cons(2, Cons(3, Cons(4, nil))))
    
In summary, list.Head contains a list element and the list.Tail contains a pointer to the next list element. This it's a singly linked list. Most of the functions here are straight functions and not methods. Currently the sole exception is Print for which a convenience function Print also exists. Functions are in keeping with how other languages implement lists.


## Functions
	func Cons(a, b interface{}) List
	func Ncons(a interface{}) List
	func Xcons(a, b interface{}) List
	
	// New creates a list of whatever is passed in as rest.
	// In Lisp this would be called List but that name is used for the type here.
	func New(rest ...interface{}) List
	
	// Last returns the last element of a list as a list.
	func Last(x List) List
	
	// Flatten takes a nested list and flattens it. This version uses an accumulator.
	func Flatten(tree interface{}) List
	
	// FlattenAlt is an alternate version of Flatten.
	func FlattenAlt(l List) List
	
	// Length as a tail recursive function.
	func LengthAlt(l List) int {
	
	// Length returns the number of top level Cons.
	func Length(l List) int {
	
	// Compare to above, is this really better and safer?
	func LengthWithCheck(l List) int {
	
	// Traverse a list depth first/head first
	func Traverse(lst List, f func(e interface{})) {
	
	func Append2(x, y List) List
	func Append(rest ...interface{}) List
	func Reverse(x List) List
	
	// Substr substitutes x for all instances of y in list z.
	func Substr(x, y interface{}, z List) List
	
	// Member returns the toplevel list that starts with x
	func Member(x interface{}, y List) List 

	// Car returns the head of the list.
	func Car(x List) interface{}
	
	// Cdr returns the tail of the list.
	func Cdr(x List) interface{}
	
	func CxR(s string, x List) interface{}
	
	// All functions below where each x in xxxx can be "a" or "d"
	func CxxxxR(s string, x List) interface{}
	
	func Rplaca(x List, y interface{}) List {
	func Rplacd(x List, y interface{}) List
	
	// Nconc is like append but it shares structure with the lists that are passed to it.
	func Nconc(x, y List) List
	
	func Nreverse(x List) List
	func Print(list interface{})
	
	/ A few utility functions mainly used for testing.
	func VerifyFlattenedList(l List) bool
	
	// Generate a list of ints starting at start and length long.
	func GenIntList(start, length int) List
	
	// Generate a nested list of ints starting at start and length long with up to depth nesting.
	func GenNestedList(astart, length, depth int) List


## Bugs
Many of the functions are tail recursive but Go doesn't yet optimize tail calls and it's unclear if it ever will.

## Future Plans
This was a fun project and I plan to add some of the other classic lisp functions such as map, and so on.


