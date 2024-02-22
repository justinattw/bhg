// Port scanner using a worker pool

package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	// Use `range` to continuously receive from the `ports` channel.
	for p := range ports {
		fmt.Println(p) // Boilerplate - just print the port numbers.
		wg.Done()
	}
}

func main() {
	// Create a channel using make().
	// The param `100` allows the channel to be buffered, meaning you can send an item without waiting for a receiver to read the item.
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	// Use for loop to start desired number of workers (in this case, 100)
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i // Send a port on the `ports` channel to the worker.
	}
	wg.Wait()
	close(ports) // Close channel after all work has been completed.
}
