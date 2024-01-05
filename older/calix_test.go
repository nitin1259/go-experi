package main

/*

Write a function that takes nonempty array of distinct integers and an integer representing a target sum. If any two numbers in the array sum up to the target sum.
The function should return them in an array. If no two numbers sum up to the sum, the function should return an empty array.

Array = [ 3,5,-4,8,11,1,-1,6]
Target sum = 10

Sample output
[-1.11]
*/

import "fmt"

func Calix_main() {
	fmt.Println("running main...")
	input := []int{3,5,-4,8,11,1,-1,6}
	target := 10

	result := findTwoNumWithSum(input, target)

	fmt.Printf("result : %+v", result)

}

func findTwoNumWithSum(arr []int, target int)[]int{
		numMap := make(map[int]bool)
		for _, num := range arr {

			comp := target - num

			if numMap[comp]{
				return []int{comp, num}
			}

			numMap[num]=true
		}
		return []int{}
}

