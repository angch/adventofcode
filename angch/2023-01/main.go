package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var words = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func day1(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		first, last := -1, -1
		firstk, lastk := 999, -1
		for k, v := range t {
			if v >= '0' && v <= '9' {
				if first < 0 {
					first = int(v) - '0'
					firstk = k
				} else {
					last = int(v) - '0'
					lastk = k
				}
			}
		}
		if last == -1 && first > -1 {
			last = first
			lastk = firstk
		}
		part1 += first*10 + last

		for k, v := range words {
			i := strings.Index(t, v)
			if i >= 0 && i < firstk {
				first = k
				firstk = i
			}
			i = strings.LastIndex(t, v)
			if i >= 0 && i > lastk {
				last = k
				lastk = i
			}
		}
		part2 += first*10 + last

		// Sanity
		if first <= 0 || first > 9 || last <= 0 || last > 9 {
			log.Fatal("OOB")
		}
	}
	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, _ := day1("test.txt")
	_, part2 := day1("test2.txt")
	if part1 != 142 || part2 != 281 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day1("input.txt"))
}
