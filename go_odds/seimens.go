package main

import (
	"fmt"
	"time"
)

func main_seimens(){

	// 1 to 5 
	// m1 -> ^2
	// m2 -> *10
	// later 1 to 5 

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// var wg sync.WaitGroup

	// wg.Add(1);

	go Prod(ch1)

	// wg.Add(2)
	go method1(ch1, ch2)
	go method2(ch2, ch3)

	// wg.Add(1)
	go cons(ch3)

	time.Sleep(1*time.Second)
}

func Prod(out chan<- int){
	defer close(out)
	// for i:=1; i<=5; i++{
	// 	out <- i
	// }

	out<-10
}

func method1(in <-chan int, out chan<- int){

	defer close(out)
	for num := range in{
		fmt.Println("method1:", num)
		out<-num*num
	}

}

func method2(in <-chan int, out chan<- int){
	defer close(out)
	for num := range in{
		fmt.Println("method2:", num)
		out<-num*10
	}
}

func cons(in <-chan int){
	for num := range in{
		fmt.Println(num)
	}
}
