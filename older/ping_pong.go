package main

import (
	"fmt"
	"time"
)

func ping_pong() {
	fmt.Println("starting main...")

	ch := make(chan struct{})

	go printPong(ch)
	go printPing(ch)

	ch <- struct{}{}

	time.Sleep(1 * time.Millisecond)
	close(ch)

}

func printPing(ch chan struct{}) {
	for {
		<-ch
		fmt.Println("Ping...")
		ch <- struct{}{}
	}
}

func printPong(ch chan struct{}) {
	for {
		<-ch
		fmt.Println("Pong...")
		ch <- struct{}{}
	}
}
