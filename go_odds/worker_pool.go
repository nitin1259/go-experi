package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Odelya, Parker, and Peter are running a cafe.
// Given that making coffee takes more time than accepting orders,
// Odelya will assist with accepting orders from customers and then pass those orders to the kitchen,
// where Parker and Peter prepare the coffee.

type order struct {
	num int
}

func (o order) String() string {
	return fmt.Sprintf("order-%02d", o.num)
}

func worker_pool_main() {

	orderChan := make(chan order, 3)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		PrepareFoodWorker("Parker", orderChan)
	}()

	go func() {
		defer wg.Done()
		PrepareFoodWorker("Peter", orderChan)
	}()

	// Odelya start taking order and placing into queue
	for i := 0; i < 10; i++ {
		waitForOrders()
		o := order{num: i}
		fmt.Printf("%s : I have will pass %s order to the channel. \n", "Odelya", o)
		orderChan <- o
	}

	fmt.Println("No more oder to place")
	close(orderChan)

	fmt.Println("Wait for workers to gracefully stop")
	wg.Wait()

	fmt.Println("All done")

}

func PrepareFoodWorker(chef string, ch <-chan order) {
	for o := range ch {
		fmt.Printf("%s: I recieved the order with num: %d \n", chef, o.num)
		prepareOrder()
		fmt.Printf("%s: I prepared the order: %d \n", chef, o.num)
	}
}

func prepareOrder() {
	processingTime := time.Duration(2+rand.Intn(2)) * time.Second
	time.Sleep(processingTime)
}

func waitForOrders() {
	processingTime := time.Duration(rand.Intn(2)) * time.Second
	time.Sleep(processingTime)
}
