// Copyright © 2019 Lawrence E. Bakst. All rights reserved.
package main

import (
	"flag"
	"fmt"
	. "leb.io/ll"
)

var start = flag.Int("start", 1, "start value for numbers")
var max = flag.Int("max", 20, "maximum value for numbers")
var depth = flag.Int("depth", 7, "max depth for nesting")

// Yet another way to flatten using Traverse
// Each run generates a random nested list, prints it, flattens it, and prints it again.
// Since the lists are generated with increasing integers it's easy to verify the correct results.

func flatten(list List) List {
	l := New()
	var f = func(v interface{}) {
		l = Splice(l, Cons(v, nil))
	}
	Traverse(list, f)
	return l
}

func main() {
	flag.Parse()
	list := GenNestedList(*start, *max, *depth)
	Print(list)
	//Print(Last(list))
	v := Rbetween(*start, *max)
	fmt.Printf("v=%d\n", v)
	Print(Member(v, list))
	return

	n := NilList
	x := New(1, 2, 3, 4)
	y := New(5, 6, 7, 8)
	z := New(9, 10, 11, 12)
	a := Append(x, y, z)
	Print(a)
	b := Substr(99, 12, a)
	Print(b)
	//Print(Nreverse(b))
	Print(b)
	Print(Member(6, b))
	Print(Append(n, y, z))
	Print(Append(n, n, z))
	Print(Append(n, n, n))
	Print(Append(n, y, n))
	Print(Append(x, n, n))
	Print(Append(x, n, z))
	return
	for list != nil {
		car := CxR("a", list)
		cdr := CxR("d", list)
		fmt.Printf("car=")
		Print(car)
		fmt.Printf("cdr=")
		Print(cdr)
		if cdr == nil {
			break
		}
		list = cdr.(List)
	}
	return
	l := flatten(list)
	Print(l)
	Print(NilList)
	l = nil
	Print(l)
	return
	l2 := New(1, 2, 3, 4)
	//fmt.Printf("%#v\n", l)
	fmt.Printf("car=")
	Print(Car(l2))
	//fmt.Printf("cdr=")
	Print(Caddr(l))
	//print(l)
	//fmt.Printf("length=%d\n", LengthAlt(l))
}
