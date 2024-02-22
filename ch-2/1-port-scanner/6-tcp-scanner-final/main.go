package main

import (
	"fmt"
	"net"
	"sort"
)

// Worker now accepts two channels, `ports` and `results`
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0 // If port is closed, send 0 to `results`
		} else {
			conn.Close()
			results <- p // Else if port is open, send port # to `results`
		}
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int) // Use separate channel to communicate results to main thread
	var openports []int       // Use slice to store results for later sorting

	// Start 100 goroutines
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// Send to works in a separate goroutine, because the result-gathering loop needs to start before more than 100 items of work can continue
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// Receive on `results` channel 1024 times
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port) // Append to slice if port is open
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports) // Sort slice of open ports
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
