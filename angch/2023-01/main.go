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

var words2 = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func reverse(t2 string) string {
	t := []rune(t2)
	for i := 0; i < len(t)/2; i++ {
		t[i], t[len(t)-i-1] = t[len(t)-i-1], t[i]
	}
	return string(t)
}

func init() {
	for k, _ := range words2 {
		words2[k] = reverse(words2[k])
	}
}

func day1(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		// first := (t[0]-'0')*10 + t[len(t)-1] - '0'
		first, last := -1, -1
		for _, v := range t {
			if v >= '0' && v <= '9' {
				if first < 0 {
					first = int(v) - '0'
				} else {
					last = int(v) - '0'
				}
			}
		}
		if last == -1 {
			last = first
		}
		// fmt.Println(first*10 + last)
		part1 += first*10 + last
		// if k == len(t)-1 {

		// fmt.Println(first)
		_ = t
	}

	return part1, part2
}

func day1b(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		// first := (t[0]-'0')*10 + t[len(t)-1] - '0'
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

		for k, v := range words {
			i := strings.Index(t, v)
			if i >= 0 && i < firstk {
				first = k
				firstk = i
			}
		}
		t2 := reverse(t)
		for k, v := range words2 {
			i := strings.Index(t2, v)
			if i >= 0 && len(t)-i > lastk {
				last = k
				lastk = len(t) - i
			}
		}
		// fmt.Println(first, last, t)

		// fmt.Println(first*10 + last)
		part2 += first*10 + last
		if first <= 0 || first > 9 || last <= 0 || last > 9 {
			log.Fatal("OOB")
		}
		// if k == len(t)-1 {

		// fmt.Println(first)
		_ = t
	}

	// Not 53318
	// Not 53327
	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day1("test.txt")
	// fmt.Println(part1, part2)
	if part1 != 142 || part2 != 0 {
		log.Fatal("Test failed")
	}

	part1, part2 = day1b("test2.txt")
	if part1 != 0 || part2 != 281 {
		log.Fatal("Test failed", part1, part2)
	}

	fmt.Println(day1("input.txt"))
	fmt.Println(day1b("input.txt"))
}
