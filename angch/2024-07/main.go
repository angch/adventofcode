package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Work(t string) (int, int) {
	l, r, _ := strings.Cut(t, ": ")
	testv, _ := strconv.Atoi(l)
	rs := strings.Split(r, " ")
	n := make([]int, len(rs))
	mult10 := make([]int, len(rs)) // Faster if we precalc 10^x
	for i, v := range rs {
		num, _ := strconv.Atoi(v)
		n[i] = num
		m := 1
		for range len(v) {
			m *= 10
		}
		mult10[i] = m
	}
	ops := make([]int, len(n)-1)
	prevvalues := make([]int, len(n))
	prevvalues[0] = n[0]
	last := 1
a:
	for {
		start := prevvalues[last-1]
		bad := len(ops) - 1
		for i := last; i < len(n); i++ {
			if ops[i-1] == 0 {
				start += n[i]
			} else if ops[i-1] == 1 {
				start *= n[i]
			} else {
				start = start*mult10[i] + n[i]
			}
			prevvalues[i] = start
			if start > testv {
				// Early exit, so we know all operators
				// after this is "bad", jump to the next
				// op.
				bad = i - 1
				break
			}
		}
		if start == testv {
			// Need to recheck, since we cached old values, unsure
			// if we're using part2 ops or not
			for _, v := range ops {
				if v == 2 {
					return 0, testv
				}
			}
			return testv, 0
		}

		// Iterate to the next ops, starting with the
		// one we know it's bad, no need to test everything
		for i := bad; i >= 0; i-- {
			ops[i]++
			if ops[i] == 3 {
				ops[i] = 0
			} else {
				// Remember where we're continuing, don't
				// need to recalc everything
				last = i + 1
				continue a
			}
		}
		break
	}
	return 0, 0
}

func day7(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	work := make(chan string, 8)
	result := make(chan [2]int, 8)
	wg := sync.WaitGroup{}
	workerNum := 16
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func(work chan string, result chan [2]int) {
			for t := range work {
				p1, p2 := Work(t)
				result <- [2]int{p1, p2}
			}
			wg.Done()
		}(work, result)
	}

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		for r := range result {
			part1 += r[0]
			part2 += r[1]
		}
		wg2.Done()
	}()

	for scanner.Scan() {
		t := scanner.Text()
		work <- t
	}

	close(work)
	wg.Wait()
	close(result)
	wg2.Wait()
	part2 += part1
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day7("test.txt")
	fmt.Println(part1, part2)
	if part1 != 3749 || part2 != 11387 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day7("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
