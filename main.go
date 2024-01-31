package main

import "fmt"

func getSecondLastIndexOf(input string, ch byte) int {
	lastIndex := -1
	secondLastIndex := -1

	for i := 0; i < len(input); i++ {
		if input[i] == ch {
			secondLastIndex = lastIndex
			lastIndex = i
		}
	}

	return secondLastIndex
}

func main() {
	input := "hello world"
	ch := 'l'

	result := getSecondLastIndexOf(input, byte(ch))

	if result != -1 {
		fmt.Printf("Second-to-last index of '%c' in '%s' is: %d\n", ch, input, result)
	} else {
		fmt.Printf("No second-to-last index of '%c' found in '%s'\n", ch, input)
	}
}
