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

func main() {
	tick := time.Tick(50 * time.Millisecond)
	slice := []int{0,1,2,3,4,5,6,7,8,9}
	outbound := firstStage(slice)
	for {
		select {
		case ch, ok := <-outbound:
			fmt.Printf("%#v, %t\n", ch, ok)
			if ok == false { return }
		case <-tick:
			fmt.Println("    .")
		default:
			fmt.Printf(".")
		}
	}
}
