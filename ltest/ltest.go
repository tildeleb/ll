package main

import "flag"
import _ "fmt"
import . "leb.io/lispylist"

var start = flag.Int("start", 1, "start value for numbers")
var n = flag.Int("n", 50, "maximum value for numbers")
var depth = flag.Int("depth", 9, "max depth for nesting")

// Yet another way to flatten using Traverse
// Each run generates a random nested list, prints it, flattens it, and prints it again.
// Since the lists are generated with increasing integers it's easy to verify the correct results.
func main() {
	flag.Parse()
	l := MakeList()
	var f = func(v interface{}) {
		l = Splice(l, Cons(v, nil))
	}
	list := GenNestedList(*start, *n, *depth)
	Print(list)
	Traverse(list, f)
	Print(l)
}
