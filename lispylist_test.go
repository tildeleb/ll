// Copyright © 2019 Lawrence E. Bakst. All rights reserved.
package lispylist_test

import (
	_ "fmt"
	. "leb.io/lispylist"
	"testing"
)

func benchmarkFlatten(b *testing.B, l []List) {
	idx := 0
	length := len(l)
	for i := 0; i < b.N; i++ {
		Flatten(l[idx])
		idx++
		if idx >= length {
			idx = 0
		}
	}
}

func benchmarkFlattenAlt(b *testing.B, l []List) {
	idx := 0
	length := len(l)
	for i := 0; i < b.N; i++ {
		FlattenAlt(l[idx])
		idx++
		if idx >= length {
			idx = 0
		}
	}
}

func BenchmarkFlatten(b *testing.B) {
	benchmarkFlatten(b, ll)
}

func BenchmarkFlattenAlt(b *testing.B) {
	benchmarkFlattenAlt(b, ll)
}

var n = 10000
var ll = make([]List, n)

func init() {
	for i := 0; i < n; i++ {
		ll[i] = GenNestedList(1, 50, 5)
	}
}

func TestFlatten(t *testing.T) {
	for i := 0; i < 1000; i++ {
		l := GenNestedList(1, 100, 9)
		//Print(l)
		l = Flatten(l)
		//Print(l)
		if VerifyFlattenedList(l) {
			t.Fail()
		}
	}
}

func ExamplOfList() {
	list := MakeList(1, 2, 3, 4, 5)
	Print(list)
	// Output:
	// (1 2 3 4 5)
}

func ExampleOfFlatten() {
	list := MakeList(1, MakeList(2, 3, MakeList(4, 5, 6)), MakeList(7, 8))
	Print(list)
	Print(Flatten(list))
	// Output:
	// (1 (2 3 (4 5 6)) (7 8))
	// (1 2 3 4 5 6 7 8)
}

func ExampleOfStrings() {
	l := MakeList("how", "now", "brown", "cow")
	Print(l)
	// Output:
	// ("how" "now" "brown" "cow")
}

func ExampleNconc() {
	l1 := MakeList(1)
	l2 := MakeList(2)
	l3 := MakeList(3, 4)
	l4 := MakeList(5, 6, 7)
	Print(Nconc(l1, l2))
	Print(Nconc(l3, l4))
	// Output:
	// (1 2)
	// (3 4 5 6 7)
}

func Example001() {
	h := Cons(1, nil)
	Print(h)
	l := h
	l.Tail = Cons(2, nil)
	l = l.Tail.(List)
	l.Tail = Cons(3, nil)
	Print(h)
	// Output:
	// (1)
	// (1 2 3)
}

func Example002() {
	Print(Cons(1, Cons(2, Cons(3, Cons(4, nil)))))
	// Output:
	// (1 2 3 4)
}

/*
func main() {
	lst := genNestedList(1, 20, 5)
	Print(lst)
	Print(flatten2(lst))
	return
	Print(flatten2(List(1, 2, 3)))
	Print(genList(1, 3))
	h := Cons(1, nil)
	l := h
	l.Tail = Cons(2, nil)
	l = l.Tail.(*Pair)
	l.Tail = Cons(3, nil)
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
*/
