package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func md5sum(key string, i int) string {
	hasher := md5.New()
	hasher.Write([]byte(key + strconv.Itoa(i)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func findmd5(key string, j int, skip int, start int, others <-chan int, result chan<- int) {
	target := strings.Repeat("0", j)
	for i := start; ; i += skip {
		select {
		// Did someone find anything?
		case max := <-others: // Async check if we "lost"
			fmt.Printf("Goroutine %d: Received bid for %d, currently at %d, waiting to give up\n", start, max, i)

			// Someone found something, can we do better?
			for ; i < max; i += skip {
				chksum := md5sum(key, i)
				if chksum[0:j] == target {
					// Found a better solution
					fmt.Printf("Goroutine %d: %s with %d zeros = %d\n", start, key, j, i)
					result <- i
					return
				}
			}
			// Nope, the other solution is better
			fmt.Printf("Goroutine %d: Giving up at %d\n", start, i)
			result <- i
			return
		default:
			// Nope, let's get to work:
			chksum := md5sum(key, i)
			if chksum[0:j] == target {
				// We found one!
				fmt.Printf("Goroutine %d: %s with %d zeros = %d\n", start, key, j, i)
				result <- i
				return
			}
		}
	}
	return
}

func main() {
	n := runtime.GOMAXPROCS(-1)

	input := "abcdef"
	zeros := 5
	if len(os.Args) > 2 {
		input = os.Args[1]
		zeros, _ = strconv.Atoi(os.Args[2])
	}

	// The typical pair of channels to send and receive messages from the goroutines
	// others is to signal the goroutines to see if they can beat this answer
	// result is to send the result to back to main
	others := make(chan int, n)
	result := make(chan int, n)

	fmt.Printf("%d goroutines\n", n)
	for i := 0; i < n; i++ {
		go findmd5(input, zeros, n, i, others, result)
	}

	// Wait for first solution
	first := <-result

	// We got the first solution, message the other goroutines to see if they can do better:
	for i := 0; i < n-1; i++ {
		others <- first
	}

	// Collect results from the rest of the goroutines
	for i := 0; i < n-1; i++ {
		r := <-result
		if r < first {
			// Found a better solution!
			first = r
		}
	}

	// Here you go:
	fmt.Println(first)
}
