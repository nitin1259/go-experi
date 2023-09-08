package main

import (
	"fmt"
	"sync"
	"time"
)

func EvenOddProgram() {
	fmt.Println("Printing even odd as seperate thread")

	evenChan := make(chan int)
	oddChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go printEven1(&wg, evenChan, oddChan)
	go printOdd1(&wg, evenChan, oddChan)

	evenChan <- 0
	wg.Wait()
}

func printEven1(wg *sync.WaitGroup, evenChan chan int, oddChan chan int) {

	defer wg.Done()
	for num := range evenChan {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Even thread: ", num)
		oddChan <- num + 1
	}
}

func printOdd1(wg *sync.WaitGroup, evenChan chan int, oddChan chan int) {

	defer wg.Done()
	for num := range oddChan {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Oddd thread: ", num)
		evenChan <- num + 1
	}
}
