package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type segment struct {
	segments [7]int
}

var digits = [10]segment{
	{[7]int{1, 1, 1, 0, 1, 1, 1}}, // 0
	{[7]int{0, 0, 1, 0, 1, 0, 0}}, // 1
	{[7]int{1, 0, 1, 1, 1, 0, 1}}, // 2
	{[7]int{1, 0, 1, 1, 0, 1, 1}}, // 3
	{[7]int{0, 1, 1, 1, 0, 1, 0}}, // 4
	{[7]int{1, 1, 0, 1, 0, 1, 1}}, // 5
	{[7]int{1, 1, 0, 1, 1, 1, 1}}, // 6
	{[7]int{1, 0, 1, 0, 0, 1, 0}}, // 7
	{[7]int{1, 1, 1, 1, 1, 1, 1}}, // 8
	{[7]int{1, 1, 1, 1, 1, 0, 1}}, // 9
}

var counts = [10]int{}

var segmentDigits = [7][]int{}

func solve(words2 []string) map[string]int {
	possiblemaps := [10]string{}
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		possiblemaps[v] = "abcdefg"
	}
	pmap2 := make(map[rune][]int, 0)
	for _, v := range "abcdefg" {
		pmap2[v] = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	for _, v := range words2 {
		if len(v) == 2 { // 1
			possiblemaps[1] = v
			for _, v2 := range v {
				pmap2[v2] = []int{1}
			}
		}
	}

	translate := make(map[string]int)
	for _, v := range words2 {
		if len(v) == 2 { // 1
			translate[v] = 1
			possiblemaps[1] = v
			for _, v2 := range v {
				pmap2[v2] = []int{1}
			}
		} else if len(v) == 7 {
			translate[v] = 8
			possiblemaps[8] = v
		} else if len(v) == 4 {
			translate[v] = 4
			possiblemaps[4] = v
		} else if len(v) == 3 {
			translate[v] = 7
			possiblemaps[7] = v
		}
	}

	// do 5
	consider := make(map[string]bool, 0)
	for _, v := range words2 {
		if len(v) == 5 {
			consider[v] = true
			// consider = append(consider, v)
		}
	}
	// log.Println(5, consider)
	for k := range consider {
		c := 0

		for _, v := range k {
			for _, v2 := range possiblemaps[1] {
				if v == v2 {
					c++
					continue
				}
			}
		}
		if c == 2 {
			// log.Println(k, "is", 3)
			translate[k] = 3
			possiblemaps[3] = k
		} else {
			c = 0
			// log.Println("4 is", possiblemaps[4])
			for _, v := range k {
				for _, v2 := range possiblemaps[4] {
					if v == v2 {
						c++
						continue
					}
				}
			}
			if c == 2 {
				// log.Println(k, "is", 2)
				translate[k] = 2
				possiblemaps[2] = k
			} else {
				// log.Println(k, "is", 5)
				translate[k] = 5
				possiblemaps[5] = k
			}
		}
	}

	// do 6
	consider = make(map[string]bool)
	for _, v := range words2 {
		if len(v) == 6 {
			consider[v] = true
		}
	}
	// log.Println(6, "x", consider)
	for k := range consider {
		c := 0

		for _, v := range k {
			for _, v2 := range possiblemaps[1] {
				if v == v2 {
					c++
					continue
				}
			}
		}
		if c == 1 {
			// log.Println(k, "is", 6)
			translate[k] = 6
			possiblemaps[6] = k
		} else {
			c = 0
			// log.Println("4 is", possiblemaps[4])
			for _, v := range k {
				for _, v2 := range possiblemaps[4] {
					if v == v2 {
						c++
						continue
					}
				}
			}
			// log.Println("x", c)
			if c == 4 {
				// log.Println(k, "is", 9)
				translate[k] = 9
				possiblemaps[9] = k
			} else {
				// log.Println(k, "is", 0)
				translate[k] = 0
				possiblemaps[0] = k
			}
		}
	}
	return translate
}

func day8(filepath string) {
	for i := 0; i < 7; i++ {
		a := make([]int, 0)
		for j := 0; j < 10; j++ {
			if digits[j].segments[i] == 1 {
				a = append(a, j)
			}
		}
		segmentDigits[i] = a
	}
	// log.Println(segmentDigits)

	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	words := make([]string, 0)
	words2 := make([]string, 0)
	output := make([][]string, 0)
	lines := make([][]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		split := strings.Split(t, " | ")

		// pos := make([]int, 0)
		s := split[1]
		s = strings.ReplaceAll(s, "\n", " ")
		s = strings.ReplaceAll(s, "|", " ")
		s = strings.ReplaceAll(s, "  ", " ")
		s = strings.ReplaceAll(s, "  ", " ")

		x := strings.Split(s, " ")
		line := make([]string, 0)
		for _, v := range x {
			w := strings.Split(v, "")
			sort.Strings(w)
			xx := strings.Join(w, "")
			words = append(words, xx)
			words2 = append(words2, xx)
			line = append(line, xx)
		}
		output = append(output, line)
		// lines = append(lines, line)

		s = split[0]
		s = strings.ReplaceAll(s, "\n", " ")
		s = strings.ReplaceAll(s, "|", " ")
		s = strings.ReplaceAll(s, "  ", " ")
		s = strings.ReplaceAll(s, "  ", " ")

		x = strings.Split(s, " ")
		for _, v := range x {
			// log.Println(v)
			w := strings.Split(v, "")
			sort.Strings(w)
			xx := strings.Join(w, "")
			words2 = append(words2, xx)
		}
		lines = append(lines, words2)
	}

	for i := 0; i < 10; i++ {
		c := 0
		for _, v := range digits[i].segments {
			if v == 1 {
				c++
			}
		}
		counts[i] = c
	}
	// log.Println(counts)

	// 1, 4, 7, or 8

	countme := []int{1, 4, 7, 8}
	part1, part2 := 0, 0
	for _, v2 := range words {
		for _, v := range countme {
			if len(v2) == counts[v] {
				part1++
				break
			}
		}
		// digit := translate[v2]
		// log.Println(v2, "is", digit)
	}

	for k2, v2 := range lines {
		n := 0
		translate := solve(v2)
		for _, v := range output[k2] {
			n = n*10 + translate[v]
		}
		// log.Println(v2, "is", n)
		part2 += n
	}

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	day8("test.txt")
	day8("test2.txt")
	day8("input.txt")
}
