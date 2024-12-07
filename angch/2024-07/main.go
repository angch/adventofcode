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
	a:
		for {
			start := n[0]
			ispart2 := false
			for i := 1; i < len(n); i++ {
				if ops[i-1] == 0 {
					start += n[i]
				} else if ops[i-1] == 1 {
					start *= n[i]
				} else {
					ispart2 = true
					vv := fmt.Sprintf("%d%d", start, n[i])
					v, _ := strconv.Atoi(vv)
					start = v
				}
			}
			if start == testv {
				// log.Println(ops, start, testv, part1)

				if ispart2 {
					part2 += testv
				} else {
					part1 += testv
				}

				break a
			}

			again := false
			for i := 0; i < len(ops); i++ {
				ops[i]++
				if ops[i] == 3 {
					ops[i] = 0
				} else {
					again = true
					break
				}
			}
			if !again {
				break
			}
		}

		// log.Println(testv, n)
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
