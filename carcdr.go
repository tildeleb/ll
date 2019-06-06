// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

package lispylist

import (
	_ "fmt"
)

// Car returns the head of the list.
func Car(x List) interface{} {
	if x == nil || x.Head == nil {
		return nil
	}
	return x.Head.(List)
	/*
		if ok {
			return l
		}
		return x.Head
	*/
}

// Cdr returns the tail of the list.
func Cdr(x List) interface{} {
	if x == nil || x.Tail == nil {
		return nil
	}
	return x.Tail.(List)
	/*
		l, ok := x.Tail.(List)
		if ok {
			return l
		}
		return x.Tail
	*/
}

func Caar(x List) interface{} {
	return (Car(Car(x).(List)))
}

func Cadr(x List) interface{} {
	return (Car(Cdr(x).(List)))
}

func Cddr(x List) interface{} {
	return (Cdr(Cdr(x).(List)))
}

func Cdar(x List) interface{} {
	return (Cdr(Car(x).(List)))
}

func Caaar(x List) interface{} {
	return (Car(Car(Car(x).(List)).(List)))
}

func Caadr(x List) interface{} {
	return (Car(Car(Cdr(x).(List)).(List)))
}

func Cadar(x List) interface{} {
	return (Car(Cdr(Car(x).(List)).(List)))
}

func Caddr(x List) interface{} {
	return (Car(Cdr(Cdr(x).(List)).(List)))
}

func Cdaar(x List) interface{} {
	return (Cdr(Car(Car(x).(List)).(List)))
}

func Cdadr(x List) interface{} {
	return (Cdr(Car(Cdr(x).(List)).(List)))
}

func Cdddr(x List) interface{} {
	return (Cdr(Cdr(Cdr(x).(List)).(List)))
}

func Caaaar(x List) interface{} {
	return (Car(Car(Car(Car(x).(List)).(List)).(List)))
}

func Caaadr(x List) interface{} {
	return (Car(Car(Car(Cdr(x).(List)).(List)).(List)))
}

func Caadar(x List) interface{} {
	return (Car(Car(Cdr(Car(x).(List)).(List)).(List)))
}

func Caaddr(x List) interface{} {
	return (Car(Car(Cdr(Cdr(x).(List)).(List)).(List)))
}

func Cadaar(x List) interface{} {
	return (Car(Cdr(Car(Car(x).(List)).(List)).(List)))
}

func Cadadr(x List) interface{} {
	return (Car(Cdr(Car(Cdr(x).(List)).(List)).(List)))
}

func Caddar(x List) interface{} {
	return (Car(Cdr(Cdr(Car(x).(List)).(List)).(List)))
}

func Cadddr(x List) interface{} {
	return (Car(Cdr(Cdr(Cdr(x).(List)).(List)).(List)))
}

func Cdaaar(x List) interface{} {
	return (Cdr(Car(Car(Car(x).(List)).(List)).(List)))
}

func Cdaadr(x List) interface{} {
	return (Cdr(Car(Car(Cdr(x).(List)).(List)).(List)))
}

func Cdadar(x List) interface{} {
	return (Cdr(Car(Cdr(Car(x).(List)).(List)).(List)))
}

func Cdaddr(x List) interface{} {
	return (Cdr(Car(Cdr(Cdr(x).(List)).(List)).(List)))
}

func Cddaar(x List) interface{} {
	return (Cdr(Cdr(Car(Car(x).(List)).(List)).(List)))
}

func Cddadr(x List) interface{} {
	return (Cdr(Cdr(Car(Cdr(x).(List)).(List)).(List)))
}

func Cddddr(x List) interface{} {
	return (Cdr(Cdr(Cdr(Cdr(x).(List)).(List)).(List)))
}
