## Day 3

### 1. Sequential vs Concurrent

- **Sequential Execution**: In sequential execution, tasks are performed one after another in a predetermined order. Each task completes before the next one begins.

  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Start")

      // Sequential tasks
      task1()
      task2()

      fmt.Println("End")
  }

  func task1() {
      fmt.Println("Executing Task 1")
  }

  func task2() {
      fmt.Println("Executing Task 2")
  }
  ```

- **Concurrent Execution**: In concurrent execution, tasks can overlap and execute simultaneously, possibly speeding up the overall process.

### 2. Goroutines - Main Goroutine, Anonymous Goroutine

- **Main Goroutine**: The `main()` function represents the main goroutine that starts when your program runs.

  ```go
  package main

  import (
      "fmt"
      "time"
  )

  func main() {
      fmt.Println("Main Goroutine starts")

      // Start a new goroutine
      go func() {
          fmt.Println("Goroutine inside main")
      }()

      time.Sleep(1 * time.Second) // Wait to allow the goroutine to execute
      fmt.Println("Main Goroutine ends")
  }
  ```

- **Anonymous Goroutine**: A goroutine can be launched anonymously using an anonymous function.

### 3. Go Runtime Scheduler

- **Go Runtime Scheduler**: Manages the execution of goroutines on OS threads. It handles the multiplexing of goroutines onto a limited number of OS threads (`GOMAXPROCS`).

### d. Wait Groups and Examples

- **Wait Groups**: `sync.WaitGroup` is used to wait for a collection of goroutines to finish before proceeding.

  ```go
  package main

  import (
      "fmt"
      "sync"
      "time"
  )

  func main() {
      var wg sync.WaitGroup

      for i := 0; i < 3; i++ {
          wg.Add(1)
          go worker(i, &wg)
      }

      wg.Wait() // Wait for all workers to finish
      fmt.Println("All workers have completed")
  }

  func worker(id int, wg *sync.WaitGroup) {
      defer wg.Done()

      fmt.Printf("Worker %d started\n", id)
      time.Sleep(2 * time.Second)
      fmt.Printf("Worker %d completed\n", id)
  }
  ```

### 5. Channels

- **Channels**: Channels are used for communication and synchronization between goroutines.

  ```go
  package main

  import (
      "fmt"
      "time"
  )

  func main() {
      ch := make(chan string)

      go func() {
          time.Sleep(2 * time.Second)
          ch <- "Hello from goroutine"
      }()

      msg := <-ch // Receive from channel
      fmt.Println(msg)
  }
  ```

### 6. Buffered vs. Unbuffered Channels

- **Unbuffered Channel**: Synchronizes sender and receiver directly.

  ```go
  package main

  import (
      "fmt"
      "time"
  )

  func main() {
      ch := make(chan int)

      go func() {
          ch <- 42 // Send to channel
      }()

      value := <-ch // Receive from channel
      fmt.Println(value)
  }
  ```

- **Buffered Channel**: Allows multiple senders to put data in the channel without immediate blocking.

  ```go
  package main

  import (
      "fmt"
      "time"
  )

  func main() {
      ch := make(chan int, 1) // Buffered channel with capacity 1

      ch <- 42 // Send to channel

      value := <-ch // Receive from channel
      fmt.Println(value)
  }
  ```

### 7. Select Statement

- **Select Statement**: Used to wait on multiple channel operations simultaneously.

  ```go
  package main

  import (
      "fmt"
      "time"
  )

  func main() {
      ch1 := make(chan string)
      ch2 := make(chan string)

      go func() {
          time.Sleep(2 * time.Second)
          ch1 <- "one"
      }()

      go func() {
          time.Sleep(1 * time.Second)
          ch2 <- "two"
      }()

      for i := 0; i < 2; i++ {
          select {
          case msg1 := <-ch1:
              fmt.Println("Received from ch1:", msg1)
          case msg2 := <-ch2:
              fmt.Println("Received from ch2:", msg2)
          }
      }
  }
  ```

### 8. Mutex Lock

- **Mutex Lock**: Used to synchronize access to shared resources across multiple goroutines.

  ```go
  package main

  import (
      "fmt"
      "sync"
  )

  var counter int
  var mutex sync.Mutex

  func main() {
      var wg sync.WaitGroup
      wg.Add(2)

      go increment(&wg)
      go increment(&wg)

      wg.Wait()
      fmt.Println("Final Counter:", counter)
  }

  func increment(wg *sync.WaitGroup) {
      defer wg.Done()

      mutex.Lock()
      defer mutex.Unlock()

      for i := 0; i < 1000; i++ {
          counter++
      }
  }
  ```

### i. Concurrency Practice

- **Concurrency Practice**: Implementing concurrent programs helps reinforce understanding of goroutines, channels, and synchronization techniques.

  ```go
  package main

  import (
      "fmt"
      "sync"
      "time"
  )

  func main() {
      var wg sync.WaitGroup

      for i := 0; i < 5; i++ {
          wg.Add(1)
          go worker(i, &wg)
      }

      wg.Wait()
      fmt.Println("All workers have completed")
  }

  func worker(id int, wg *sync.WaitGroup) {
      defer wg.Done()

      fmt.Printf("Worker %d started\n", id)
      time.Sleep(time.Duration(id) * time.Second)
      fmt.Printf("Worker %d completed\n", id)
  }
  ```

In summary, understanding these concurrency topics and practicing them with examples will help you develop robust and efficient concurrent Go programs. Experiment with different scenarios and explore additional features of Go's concurrency model to gain deeper insights and mastery.
