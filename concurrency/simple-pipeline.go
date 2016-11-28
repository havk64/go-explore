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

func finalStage(s []int) {
	for n := range secondStage(firstStage(s)) {
		fmt.Println(n)
	}
}

func main() {
	start := time.Now()
	slice := []int{0,1,2,3,4,5,6,7,8,9}
	finalStage(slice)
	fmt.Printf("%v\n",time.Since(start))
}
