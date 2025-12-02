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

func dupe(a string) bool {
	left := a[0 : len(a)/2]
	right := a[len(a)/2:]
	return left == right
}

func dupe2(a string) bool {
a:
	for i := 2; i < 10; i++ {
		l := len(a) / i

		t1 := 0
		t2 := l
		if l == 0 {
			break
		}

		for ; t2 < len(a); t1, t2 = t1+l, t2+l {
			if t2+l > len(a) {
				continue a
			}
			left := a[t1:t2]
			right := a[t2 : t2+l]
			if left != right {
				continue a
			}
		}
		return true
	}
	return false
}

func day2(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()

		foos := strings.Split(t, ",")
		for _, t2 := range foos {
			rg := strings.Split(t2, "-")

			from, _ := strconv.Atoi(rg[0])
			to, _ := strconv.Atoi(rg[1])

			for t3_ := from; t3_ <= to; t3_++ {
				t3 := strconv.Itoa(t3_)
				// fmt.Println(t3)

				if dupe(t3) {
					// log.Println("bad", t3)
					part1 += t3_
				}
				if dupe2(t3) {
					// log.Println("bad", t3)
					part2 += t3_
				}
			}
		}

	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day2("test.txt")
	fmt.Println(part1, part2)
	if part1 != 1227775554 || part2 != 4174379265 {
		// if part1 != 0 || part2 != 0 {
		log.Fatal("Test failed ", part1, part2)
	}
	part1, part2 = day2("input.txt")
	fmt.Println(part1, part2)
	fmt.Println("Elapsed time:", time.Since(t1))
}
