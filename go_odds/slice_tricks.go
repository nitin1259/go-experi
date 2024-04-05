package main

import (
	"fmt"
)

func slice_tricks() {
	fmt.Println("starting main ...")

	/*

		var sl []int
		fmt.Println(sl, sl == nil, len(sl), cap(sl))
		sl = append(sl, 2, 3, 4)
		fmt.Println(sl, sl == nil, len(sl), cap(sl))
		sl = append(sl, 6, 7)
		fmt.Println(sl, sl == nil, len(sl), cap(sl))
		sl = append(sl, 9, 10)
		fmt.Println(sl, sl == nil, len(sl), cap(sl))

	*/

	a := [3]int{1, 2, 3}
	b := a[:]
	b = append(b, 4)
	fmt.Println(a, b)

	// -----------

	s := []int{1, 2, 3}
	p := &s[1]
	// fmt.Println(p, *p, &s[1])
	s = append(s, 4)
	*p = 5
	// fmt.Println(p, *p, &s[1])
	fmt.Println(s)

}
