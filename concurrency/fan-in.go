package main

import (
	"fmt"
	"time"
	"sync"
)

func firstStage(num []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()
	return out
}

func secondStage(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	start := time.Now()
	slice := []int{0,1,2,3,4,5,6,7,8,9}

	in := firstStage(slice)
	c1 := secondStage(in)
	c2 := secondStage(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
	fmt.Printf("%v\n", time.Since(start))
}
