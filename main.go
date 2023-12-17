package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	// evenCh := make(chan int)
	// oddCh := make(chan int)

	ch := make(chan int)

	wg.Add(2)

	go generateOdd(&wg, ch)
	go generateEven(&wg, ch)
	

	for num := range ch{
		fmt.Println(num)
	}
	wg.Wait()
	close(ch)
}

func generateEven(wg *sync.WaitGroup, ch chan int){
defer wg.Done()
	for i:=0; i< 10; i++{
		if i%2==0{
			ch<-i
		}
	}

}

func generateOdd(wg *sync.WaitGroup, ch chan int){

	wg.Done()
	for i:=1; i< 10; i++{
		if i%2==1{
			ch<-i
		}
	}
	
}