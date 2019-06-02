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

func GenList(start, length int) List {
	idx := start + length - 1
	l := Cons(idx, nil)
	idx--
	for idx >= start {
		l = Cons(idx, l)
		idx--
	}
	return l
}

func GenNestedList(v, n, depth int) List {
	var gnl func(d int) List
	gnl = func(d int) List {
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

func benchmarkFlatten1(b *testing.B, l []List) {
	idx := 0
	length := len(l)
	for i := 0; i < b.N; i++ {
		Flatten1(l[idx])
		idx++
		if idx >= length {
			idx = 0
		}
	}
}

func benchmarkFlatten2(b *testing.B, l []List) {
	idx := 0
	length := len(l)
	for i := 0; i < b.N; i++ {
		Flatten2(l[idx])
		idx++
		if idx >= length {
			idx = 0
		}
	}
}

func BenchmarkFlatten1(b *testing.B) {
	benchmarkFlatten1(b, ll)
}

func BenchmarkFlatten2(b *testing.B) {
	benchmarkFlatten2(b, ll)
}

var n = 10000
var ll = make([]List, n)

func init() {
	for i := 0; i < n; i++ {
		ll[i] = GenNestedList(1, 50, 5)
	}
}
