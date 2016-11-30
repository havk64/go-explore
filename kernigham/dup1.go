// Dup1 prints each line that appears more than once in the stdin
// Using bufio package.
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// Scan advances the Scanner to next token, which will then be
	// available through Text or Bytes method.
	for input.Scan() {
		counts[input.Text()]++
	}

	// for eventual errors from input.Err()
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
