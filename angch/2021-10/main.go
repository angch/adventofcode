package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func score(t string) (int, string) {
	open := "[({<"
	close := "])}>"
	score := []int{57, 3, 1197, 25137}
	f := ""
	e := 0
	o := 0
	for _, v := range t {
		// fmt.Println(f, o)

		if strings.Contains(open, string(v)) {
			f += string(v)
			o++
		}
		if strings.Contains(close, string(v)) {
			last := f[len(f)-1]
			f = f[:len(f)-1]
			// log.Println(f, string(v))
			for i := 0; i < len(open); i++ {
				if byte(v) == close[i] {
					if open[i] != last {
						// fmt.Println("Expected", string(open[i]), "got", string(last))
						// fmt.Println("Expected", open[i], "got", last)
						if e == 0 {
							e += score[i]
						}
					}
				}
			}
			o--
		}
	}
	if len(f) == 0 {
		// fmt.Println("No open brackets", o)
		return e, f
	}
	// fmt.Println(f, o)
	return e, f
}

var scores = map[byte]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func day10(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	part1, part2 := 0, 0
	s3 := make([]int, 0)
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		s, incomplete := score(t)

		if incomplete != "" && s == 0 {
			// fmt.Println("Incomplete", incomplete)
			s2 := 0
			for i := len(incomplete) - 1; i >= 0; i-- {
				v := incomplete[i]
				s2 = s2*5 + scores[v]
				// fmt.Println("  ", v, scores[v], s2)
			}
			// log.Println("part2", s2)
			s3 = append(s3, s2)
		}

		// fmt.Println(t, s)
		part1 += s
	}
	sort.Ints(s3)
	// log.Println(s3)
	part2 = s3[len(s3)/2]

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	day10("test.txt")
	day10("input.txt")
}
