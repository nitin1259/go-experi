package main

import (
	"fmt"
	"sync"
)

/**

Q1. Find the duplicate number in a sorted array that has n elements. The array will contain all the numbers ranging from 1 to n-1.
E.g
{1,2,2,3,4,5, 6} = 2
{1,2,3,4,4,5,6}  = 4


*/

func main() {

	// input := []int{1,2,3,4,5,6,6}

	// fmt.Println(findDuplicate(input))


	// Q2. Print numbers from 1 to 10 sequentially using 2 goroutines.

	ch := make(chan bool)

	var wg sync.WaitGroup
	
	wg.Add(2)

	go printOdd(ch, &wg)
	go printEven(ch, &wg)


	wg.Wait()
	close(ch)
}

func printOdd(ch chan bool, wg *sync.WaitGroup){
	defer wg.Done()
	for i:= 1;i<10; i++{
		if i%2==1{
			fmt.Println(i)
			ch<-true
		}else{
			<-ch
		}
	}
}


func printEven(ch chan bool, wg *sync.WaitGroup){
	defer wg.Done()
	for i:= 2;i<=10; i++{
		if i%2==0{
			<-ch
			fmt.Println(i)
		}else{
			ch <- true
		}
	}
}


func findDuplicate( input []int ) int{
	left, right := 0, len(input)-1
	for left < right {
		mid := (left + right) / 2
		// fmt.Println(left, right, mid)
		if input[mid] == input[mid+1] || input[mid] == input[mid-1]{
			return input[mid]
		}

		if input[mid] == mid+1{
			left = mid+1
		}else{
			right = mid-1
		}
	}

	return 0;
}