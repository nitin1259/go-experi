package main

import "fmt"

func reverse_main() {
	fmt.Println("reverse string")

	fmt.Println("Reverse String: ", reverseString("Kapila pashu Aahar"))

	fmt.Println("Reverse String: ", reverseStr("Kapilap"))

}

func reverseString(str string) string {
	if len(str) <= 1 {
		return str
	}

	return reverseString(str[1:]) + string(str[0])
}

func reverseStr(str string) string {
	rns := []rune(str)

	for i := 0; i < len(rns)/2; i++ {
		rns[i], rns[len(rns)-i-1] = rns[len(rns)-i-1], rns[i]
	}

	return string(rns)
}
