package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	fillInterval time.Duration
	lastRefill   time.Time
	capacity     int
	tokens       int
}

func NewRateLimiter(capacity int, fillInterval time.Duration) *RateLimiter {

	return &RateLimiter{
		tokens:       capacity,
		capacity:     capacity,
		fillInterval: fillInterval,
		lastRefill:   time.Now(),
	}
}
func (rl *RateLimiter) AllowRequest() bool {

	timeNow := time.Now()
	fmt.Printf("****** fillInterval: %v, timeNow: %v, lastRefill: %v, newtoken: %v \n", rl.fillInterval, timeNow, rl.lastRefill, int(timeNow.Sub(rl.lastRefill)/rl.fillInterval))
	rl.tokens = rl.tokens + int(timeNow.Sub(rl.lastRefill)/rl.fillInterval)

	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity
	}

	if rl.tokens > 0 {
		rl.tokens--
		rl.lastRefill = timeNow
		return true
	}

	return false
}

var limiter = NewRateLimiter(10, 1*time.Second)

func rate_limiter_main() {

	for i := 0; i < 20; i++ {
		time.Sleep(300 * time.Millisecond)
		if !limiter.AllowRequest() {

			fmt.Printf("=====> request not served i: %d -> %+v \n ", i, limiter)
		} else {
			fmt.Printf("-----> request servedi: %d -> %+v \n ", i, limiter)
		}
	}

}
