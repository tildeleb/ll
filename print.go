// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

package ll

import (
	"fmt"
)

var Pmax = 45 // The maximum number of list elements Print will print.

// Print prints a list up to Pmax elements long
// would be nice if pmax triggers to print the last few elements too
func Print(list interface{}) {
	var cnt = 0
	var prt = func(x interface{}) {
		_, ok := x.(string)
		if ok {
			fmt.Printf("%q", x)
		} else {
			fmt.Printf("%v", x)
		}
	}
	var p func(list interface{})
	p = func(alist interface{}) {
		list, ok := alist.(List)
		if !ok {
			prt(alist)
			return
		}
		fmt.Printf("(")
		prev := list
		for {
			if cnt >= Pmax {
				fmt.Printf("...")
				break
			}
			if prev == nil {
				break
			}
			if prev.Head != nil {
				h, ok := prev.Head.(List)
				if ok {
					p(h)
				} else {
					prt(prev.Head)
					cnt++
				}
			}
			if prev.Tail != nil {
				prev = prev.Tail.(List)
				fmt.Printf(" ")
			} else {
				break
			}
		}
		fmt.Printf(")")
	}
	//fmt.Printf("%v\n", list)
	p(list)
	fmt.Printf("\n")
}

func print(list List) {
	fmt.Printf("%v\n", list)
}
