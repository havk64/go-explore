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
	tick := time.Tick(1 * time.Nanosecond)
	slice := []int{0,1,2,3,4,5,6,7,8,9}

	ch := firstStage(slice)
	outbound := process(ch)
	for {
		select {
		case res, ok := <-outbound:
			if ok {
				fmt.Println(res)
			} else {
				fmt.Printf("%v\n", time.Since(start))
				return
			}
		case <-tick:
			fmt.Printf(".")
		}
	}

}
