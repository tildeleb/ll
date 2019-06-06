// Copyright © 2019 Lawrence E. Bakst. All rights reserved.
//
// I wanted to code this in Go.
// Unlike Ruby, Python, and others, Go doesn't have a built-in list type.
// So I wrote a simple list package based on the way lists were originally designed in Lisp.
// See the code at:
// https://github.com/tildeleb/lispylist
// The code there has a full test suite and doc.
// This code is https://github.com/tildeleb/lispylist/example/example.go

package main

import . "leb.io/lispylist"

// This function below was copied from FlattenAlt in my "lispylist" package.
// The assignment doesn't specify how to flatten a list.
// Nor does the assignment specify if the resulting flattened list can share structure with the argument.
// All the versions of flatten I coded do a depth first, head first traversal of the list.
// This version uses Spice and the resulting flattened list shares structure with its argument.
// See Flatten() or the Traverse() example in "ltest.go" for examples that doesn't share structure with the argument.
// All code has been tested with lists up to length of 100,000.
// For lists larger than 100,000 all the functions that are tail recursive would have to be verified as tail recursive.
// If you are working with lists larger than 100,000 you should probably consider an alternative data structure anyway.
func flatten(lst List) List {
	if lst == nil {
		return nil
	}
	Head, Headok := lst.Head.(List)
	Tail, _ := lst.Tail.(List)
	if Headok {
		return Splice(flatten(Head), flatten(Tail))
	}
	return Splice(Cons(lst.Head, nil), flatten(Tail))
}

// Each run generates a random nested list, prints it, flattens it, and prints it again.
// Since the lists are generated with increasing integers it's easy to verify the correct results.
func main() {
	list := GenNestedList(1, 20, 7)
	Print(list)
	Print(flatten(list))
}
