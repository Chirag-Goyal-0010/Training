package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. Basic Goroutine
func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 2. Channel Communication
func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send i to channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close channel when done
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}
}

// 3. Select Statement
func selectExample(ch1, ch2 chan int) {
	for i := 0; i < 5; i++ {
		select {
		case ch1 <- i:
			fmt.Printf("Sent %d to ch1\n", i)
		case ch2 <- i * 2:
			fmt.Printf("Sent %d to ch2\n", i*2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}

// 4. Worker Pool
type Job struct {
	ID       int
	Duration time.Duration
}

func worker(id int, jobs <-chan Job, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(job.Duration)
		results <- job.ID
	}
}

// 5. Mutex for Synchronization
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// 6. WaitGroup Example
func processWithWaitGroup(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Processing %d\n", id)
	time.Sleep(100 * time.Millisecond)
}

// 7. Context with Timeout
func processWithTimeout(done chan bool) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operation timed out")
	case <-done:
		fmt.Println("Operation completed")
	}
}

func main() {
	// 1. Basic Goroutine
	fmt.Println("=== Basic Goroutine ===")
	go sayHello("Alice")
	go sayHello("Bob")
	time.Sleep(100 * time.Millisecond)

	// 2. Channel Communication
	fmt.Println("\n=== Channel Communication ===")
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go producer(ch)
	go consumer(ch, &wg)
	wg.Wait()

	// 3. Select Statement
	fmt.Println("\n=== Select Statement ===")
	ch1 := make(chan int)
	ch2 := make(chan int)
	go selectExample(ch1, ch2)
	for i := 0; i < 5; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("Received from ch1: %d\n", x)
		case x := <-ch2:
			fmt.Printf("Received from ch2: %d\n", x)
		}
	}

	// 4. Worker Pool
	fmt.Println("\n=== Worker Pool ===")
	numJobs := 5
	numWorkers := 3
	jobs := make(chan Job, numJobs)
	results := make(chan int, numJobs)
	var wg2 sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg2.Add(1)
		go worker(w, jobs, results, &wg2)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{
			ID:       j,
			Duration: time.Duration(j*100) * time.Millisecond,
		}
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg2.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Job %d completed\n", result)
	}

	// 5. Mutex Example
	fmt.Println("\n=== Mutex Example ===")
	counter := SafeCounter{}
	for i := 0; i < 10; i++ {
		go counter.Increment()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Final count: %d\n", counter.GetCount())

	// 6. WaitGroup Example
	fmt.Println("\n=== WaitGroup Example ===")
	var wg3 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg3.Add(1)
		go processWithWaitGroup(&wg3, i)
	}
	wg3.Wait()
	fmt.Println("All processes completed")

	// 7. Timeout Example
	fmt.Println("\n=== Timeout Example ===")
	done := make(chan bool)
	go processWithTimeout(done)
	time.Sleep(1 * time.Second)
	done <- true
	time.Sleep(100 * time.Millisecond)

	// 8. Buffered Channel
	fmt.Println("\n=== Buffered Channel ===")
	bufferedCh := make(chan int, 3)
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	close(bufferedCh)
	for v := range bufferedCh {
		fmt.Printf("Buffered channel value: %d\n", v)
	}
} 