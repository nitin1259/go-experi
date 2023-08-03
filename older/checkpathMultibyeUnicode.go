package main

import (
	"fmt"
	"unicode/utf8"
)

// func main() {

// 	temp := []int{30, 29, 33, 40, 12, 34, 10, 50, 12}

// 	fmt.Println(dailyTemperatures(temp))

// }

func IsPathContainsMultibyteUnicode(path string) bool {
	for _, r := range path {
		fmt.Println(string(r), "->", r)
		if i := utf8.RuneLen(r); i == 4 || i == -1 {
			return true
		}
	}
	return false
}

func cmd_main() {

	path1 := "£1500 Investor "
	path2 := "S H A Y E 🌿 (@onmyhand) • Instagram photos and videos_files.pdf"

	// if pathContainsMultibyteChars(path1) {
	// 	fmt.Println(path1, "contains multibyte characters")
	// } else {
	// 	fmt.Println(path1, "does not contain multibyte characters")
	// }

	// if pathContainsMultibyteChars(path2) {
	// 	fmt.Println(path2, "contains multibyte characters")
	// } else {
	// 	fmt.Println(path2, "does not contain multibyte characters")
	// }

	if IsPathContainsMultibyteUnicode(path1) {
		fmt.Println(path1, "contains multibyte characters")
	} else {
		fmt.Println(path1, "does not contain multibyte characters")
	}

	if IsPathContainsMultibyteUnicode(path2) {
		fmt.Println(path2, "contains multibyte characters")
	} else {
		fmt.Println(path2, "does not contain multibyte characters")
	}

}
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[j] = i - j
		}
		stack = append(stack, i)
	}
	return res
}
