package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func day7(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		l, r, _ := strings.Cut(t, ": ")
		testv, _ := strconv.Atoi(l)
		rs := strings.Split(r, " ")
		n := make([]int, 0, len(rs))
		for _, v := range rs {
			nums, _ := strconv.Atoi(v)
			n = append(n, nums)
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
					// ispart2 = true
					for range len(rs[i]) {
						start *= 10
					}
					start += n[i]
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
						part2 += testv
						break a
					}
				}
				part1 += testv
				break a
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
	}
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
