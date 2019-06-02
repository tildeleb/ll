// I wanted to code this in Go.
// Unlike Ruby, Python, and others, Go doesn't have a built-in list type.
// So I wrote a simple list package based on the way lists were originally designed in Lisp.
// See the code at:
// https://github.com/tildeleb/lispylist
// The code there has a full test suite and doc.

package main

import . "leb.io/lispylist"

// This function is from the "lispylist" package.
// There are many ways to code flatten. I believe this version is one of the easiest to understand.
// Nconc is just a destructive version of append, so the argument is modified not copied.
// This version of flatten is tail recursive.
// Go usually optimizes tail recursion to a goto so the stack doesn't grow.
// I didn't verify this.
func flatten(lst *List) *List {
	if lst == nil {
		return nil
	}
	Head, Headok := lst.Head.(*List)
	Tail, _ := lst.Tail.(*List)
	if Headok {
		return Nconc(flatten(Head), flatten(Tail))
	}
	return Nconc(Cons(lst.Head, nil), flatten(Tail))
}

// Each run generates a random nested list, prints it, flattens it, and prints it again.
// Since the lists are generated with increasing integers it's easy to verify the correct results.
func main() {
	list := GenNestedList(1, 21, 7)
	Print(list)
	Print(flatten(list))
}
