// Copyright © 2019 Lawrence E. Bakst. All rights reserved.
package main

import (
	"flag"
	_ "fmt"
	. "leb.io/lispylist"
)

var start = flag.Int("start", 1, "start value for numbers")
var n = flag.Int("n", 20, "maximum value for numbers")
var depth = flag.Int("depth", 7, "max depth for nesting")

// Yet another way to flatten using Traverse
// Each run generates a random nested list, prints it, flattens it, and prints it again.
// Since the lists are generated with increasing integers it's easy to verify the correct results.

func flatten(list *List) *List {
	l := MakeList()
	var f = func(v interface{}) {
		l = Splice(l, Cons(v, nil))
	}
	Traverse(list, f)
	return l
}

func main() {
	flag.Parse()
	list := GenNestedList(*start, *n, *depth)
	Print(list)
	l := flatten(list)
	Print(l)
	//fmt.Printf("length=%d\n", LengthAlt(l))
}
