package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup // sync.WaitGroup acts as a synchronised counter

	// Change between 1024 and 65535
	for i := 1; i <= 1024; i++ {
		wg.Add(1) // Increment the sync.WaitGroup counter every time you create a goroutine to scan a port
		go func(j int) {
			defer wg.Done() // Deferred call decreemnts the counter whenever one unit of work has been performed

			address := fmt.Sprintf("127.0.0.1:%d", j)
			// address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait() // Blocks until all work has been done and counter has returned to zero.
}
