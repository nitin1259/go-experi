package main

import "fmt"

func sb_main() {
	/*
Input: [-3, -4, 5, -1, 2, -4, 6, -1]
         Output: 8

	Input: [-2, 3, -1, 2]
         Output: 4
	*/
	// max sum of sub array
	fmt.Println(maxSubArray([]int{-2, 4, -3, 5, -5, -5, 2, 10}))
}

func maxSubArray(num []int)int{
	maxSum := num[0]
	currSum := num[0]
	for i:=1; i<len(num); i++{
		currSum = maxi(num[i], currSum+num[i])
		maxSum = maxi(maxSum, currSum)
	}
	return maxSum
}

func maxi(a, b int)int{
	if a>b{
		return a
	}
	return b
}