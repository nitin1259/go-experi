package main

import (
	"fmt"
	"sync"
)

func printHelloOld(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the waitgroup that the goroutine has completed
	fmt.Println("Hello")
	ch <- true // Signal that "hello" has been printed
}

func printWorldOld(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the waitgroup that the goroutine has completed
	<-ch            // Wait for "hello" to be printed
	fmt.Println("World")
}

func printHelloAndWorldInSeperateGoroutine() {
	ch := make(chan bool) // Channel to synchronize printing
	var wg sync.WaitGroup // WaitGroup to wait for goroutines to complete
	wg.Add(2)             // Add the number of goroutines to wait for

	go printHelloOld(ch, &wg) // Launch the "hello" goroutine
	go printWorldOld(ch, &wg) // Launch the "world" goroutine

	wg.Wait() // Wait for all goroutines to complete
	close(ch) // Close the channel
}

func printHello(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the waitgroup that the goroutine has completed
	for i := 0; i < 10; i++ {
		<-ch // Wait for signal to print
		fmt.Println("Hello")
		ch <- true // Signal next goroutine to print
	}
}

func printWorld(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the waitgroup that the goroutine has completed
	for i := 0; i < 10; i++ {
		<-ch // Wait for signal to print
		fmt.Println("World")
		ch <- true // Signal next goroutine to print
	}
}

func printHelloAndWorld10TimesInSeperateGoroutine() {
	ch := make(chan bool) // Channel to synchronize printing
	var wg sync.WaitGroup // WaitGroup to wait for goroutines to complete
	wg.Add(2)             // Add the number of goroutines to wait for

	go printHello(ch, &wg) // Launch the "hello" goroutine
	go printWorld(ch, &wg) // Launch the "world" goroutine

	ch <- true // Start the printing by sending a signal to the "hello" goroutine

	wg.Wait() // Wait for all goroutines to complete
	close(ch) // Close the channel
}
