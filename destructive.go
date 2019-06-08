// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

package lispylist

import (
	_ "fmt"
)

func Rplaca(x List, y interface{}) List {
	x.Head = y
	return x
}

func Rplacd(x List, y interface{}) List {
	x.Tail = y
	return x
}

// Nconc is like append but it shares structure with the lists that are passed to it.
func Nconc(x, y List) List {
	switch {
	case x == nil && y == nil:
		return nil
	case x != nil && y == nil:
		return x
	case x == nil && y != nil:
		return y
	case x.Tail != nil:
		l := Last(x)
		l.Tail = y
		return x
	default:
		x.Tail = y
		return x
	}
}

func nrev(x, y List) List {
	if Cdr(x) == nil {
		return Rplacd(x, y)
	}
	return nrev(Cdr(x).(List), Rplacd(x, y))
}

func Nreverse(x List) List {
	if x == nil {
		return nil
	}
	return nrev(x, nil)
}
