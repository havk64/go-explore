package main

import (
	"fmt"
	"time"
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

func process(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	start := time.Now()
	slice := []int{0,1,2,3,4,5,6,7,8,9}

	for n := range process(firstStage(slice)) {
		fmt.Println(n)
	}
	fmt.Printf("%v\n",time.Since(start))
}
