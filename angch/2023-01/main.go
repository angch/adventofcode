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

func day1(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		first, last := -1, -1
		firstk, lastk := 999, -1
		for k, v := range t {
			v2 := int(v) - '0'
			if v2 >= 0 && v2 <= 9 {
				if first < 0 {
					first, firstk = v2, k
				}
				last, lastk = v2, k
			}
		}
		part1 += first*10 + last

		for k, v := range words {
			i := strings.Index(t, v)
			if i >= 0 && i < firstk {
				first, firstk = k, i
			}
			i = strings.LastIndex(t, v)
			if i >= 0 && i > lastk {
				last, lastk = k, i
			}
		}
		part2 += first*10 + last
	}
	return
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
