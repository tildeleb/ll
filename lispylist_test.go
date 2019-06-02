package lispylist_test

//import "fmt"
import "time"
import "testing"
import "math/rand"

import . "leb.io/lispylist"

var s = rand.NewSource(time.Now().UTC().UnixNano())
var r = rand.New(s)

// rbetween returns random int [a, b]
func rbetween(a int, b int) int {
	return r.Intn(b-a+1) + a
}

func GenList(start, length int) *List {
	idx := start + length - 1
	l := Cons(idx, nil)
	idx--
	for idx >= start {
		l = Cons(idx, l)
		idx--
	}
	return l
}

func GenNestedList(v, n, depth int) *List {
	var gnl func(d int) *List
	gnl = func(d int) *List {
		d--
		if d > 0 {
			lst := Cons(nil, nil)
			a := gnl(rbetween(1, d))
			b := gnl(rbetween(1, d))
			lst.Head = a
			lst.Tail = MakeList(b)
			return lst
		} else {
			l := rbetween(1, n/rbetween(1, depth))
			lst := GenList(v, l)
			v += l
			return lst
		}
	}
	d := rbetween(1, depth)
	return gnl(d)
}

func benchmarkFlatten(b *testing.B, l []*List) {
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

func benchmarkFlattenSlow(b *testing.B, l []*List) {
	idx := 0
	length := len(l)
	for i := 0; i < b.N; i++ {
		FlattenSlow(l[idx])
		idx++
		if idx >= length {
			idx = 0
		}
	}
}

func BenchmarkFlatten(b *testing.B) {
	benchmarkFlatten(b, ll)
}

func BenchmarkFlattenSlow(b *testing.B) {
	benchmarkFlattenSlow(b, ll)
}

var n = 10000
var ll = make([]*List, n)

func init() {
	for i := 0; i < n; i++ {
		ll[i] = GenNestedList(1, 50, 5)
	}
}

func ExampleOfList() {
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

func Example001() {
	h := Cons(1, nil)
	Print(h)
	l := h
	l.Tail = Cons(2, nil)
	l = l.Tail.(*List)
	l.Tail = Cons(3, nil)
	Print(h)
	// Output:
	// (1)
	// (1 2 3)
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
