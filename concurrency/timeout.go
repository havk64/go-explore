package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Starting...")
	timeout := make(chan string, 2)
	tick := time.Tick(1 * time.Second)
	ch := make(chan int, 1)
	go func() {
		time.Sleep(5 * time.Second)
		timeout <- "timeout"
	}()
	for {
		select {
		case <-ch:
			fmt.Println("Channel")
		case <-tick:
			fmt.Printf(".\n")
		case f := <-timeout:
			fmt.Println(f, time.Since(start))
			return
		}
	}
}
