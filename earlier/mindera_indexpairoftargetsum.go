package main

import "fmt"

func mindera_indexpairoftargetsum() {

	// []int{8,2,3,8,4,5,1,9,7,6,1,8,7}
	// 8->1,1
	// 1-> 8, 8 // skip
	// 1 -> 8, 8
	// 8 -> 1,1
	// target  = 9

	// map [1-> complement num]int-> index of num

	nums := []int{8,2,3,4,5,1,9,7,6,1,8,7}
	target := 9

	pais := findPairs(nums, target)

	if len(pais)>0{
		fmt.Printf("%v", pais)
	}else{
		fmt.Println("no pairs found")
	}

}

func findPairs(nums []int, target int)[][]int{
	result := [][]int{}

	indexMap := make(map[int][]int)

	for i, num := range nums{
		// fmt.Println(i, indexMap)
		complement := target -num // 9 - 8


		if complement == target{
			continue
		}
		if valIndex, ok:= indexMap[complement]; !ok{
			indexMap[complement] = []int{i}
			
		}else{
			valIndex = append(valIndex, i)
			indexMap[complement] = valIndex
			// result = append(result, []int{i, valIndex})
		}
		
	}

	for i, num := range nums{

		if val, ok:= indexMap[num] ; ok{
			for _, ind := range val{
				result = append(result, []int{i, ind})
			}
		}
	}
	return result
}