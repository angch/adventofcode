package main

import (
	"bufio"
	"fmt"
	"os"
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

// This is the only one we need, tbh, but we'll derive it manually from
// the segments above, instead of hardcoding it
var digitSegmentCount = [10]int{}

// This is how it looks like, if we hardcoded it:
// var counts = [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

func init() {
	for i, digit := range digits {
		c := 0
		for _, v := range digit.segments {
			c += v
		}
		digitSegmentCount[i] = c
	}
}

func solve(outputWords []string) map[string]int {
	possiblemaps := [10]string{}
	translate := make(map[string]int)

	consider5 := make(map[string]bool)
	consider6 := make(map[string]bool)

	for _, v := range outputWords {
		// Based on the length of the word, we can determine which
		// of the digits it corresponds to, except for the 5 and 6
		// digits, which we need to consider separately
		switch len(v) {
		case 2: // Only "1" uses 2 segments
			translate[v] = 1
			possiblemaps[1] = v
		case 3: // Only "7" uses 3 segments
			translate[v] = 7
			possiblemaps[7] = v
		case 4: // Only "4" uses 4 segments
			translate[v] = 4
			possiblemaps[4] = v
		case 5: // 5 segments, only "2", "3", "5"
			consider5[v] = true
		case 6: // 6 segments, only "0", "6", "9"
			consider6[v] = true
		case 7: // Only "8" uses 7 segments
			translate[v] = 8
			possiblemaps[8] = v
		}
	}

	// Now we need to figure out which of the words with 5 segments
	// ie, "2", "3", "5", based on how many of the shard segments
	// with "1" and "4"
	for k := range consider5 {
		c := 0
		for _, v := range k {
			// Count the number of same segments with "1"
			for _, v2 := range possiblemaps[1] {
				if v == v2 {
					c++
				}
			}
		}

		// Only "3" has 2 segments shared with "1"
		if c == 2 {
			translate[k] = 3
			possiblemaps[3] = k
		} else {
			c = 0

			// Now count the number of same segments with "4"
			// Only "2" and "5" is left
			for _, v := range k {
				for _, v2 := range possiblemaps[4] {
					if v == v2 {
						c++
					}
				}
			}

			// Only 2 segments matched with "4", which is "2"
			if c == 2 {
				translate[k] = 2
				possiblemaps[2] = k
			} else {
				// The only other match is "5"
				translate[k] = 5
				possiblemaps[5] = k
			}
		}
	}

	// Checking digits with 6 segments, ie "0", "6", "9"
	// against "1" and "4"
	for k := range consider6 {
		c := 0
		for _, v := range k {
			for _, v2 := range possiblemaps[1] {
				if v == v2 {
					c++
				}
			}
		}
		// Only "6" shares one segment with "1"
		if c == 1 {
			translate[k] = 6
			possiblemaps[6] = k
		} else {
			c = 0
			// Count number of segments matched with "4"
			// It's either "0" or "9"
			for _, v := range k {
				for _, v2 := range possiblemaps[4] {
					if v == v2 {
						c++
					}
				}
			}

			// 4 segments matched with "4", which is "9"
			if c == 4 {
				translate[k] = 9
				possiblemaps[9] = k
			} else {
				translate[k] = 0
				possiblemaps[0] = k
			}
		}
	}
	return translate
}

func day8(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineOutputs := make([][]string, 0)
	lines := make([][]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		split := strings.Split(t, " | ")

		outputSegments := strings.Split(split[1], " ")
		lineOutputs = append(lineOutputs, outputSegments)

		lineSegments := append(outputSegments, strings.Split(split[0], " ")...)
		lines = append(lines, lineSegments)
	}

	// 1, 4, 7, or 8
	countme := []int{1, 4, 7, 8}
	part1, part2 := 0, 0
	for _, v1 := range lineOutputs {
		for _, v2 := range v1 {
			for _, v3 := range countme {
				if len(v2) == digitSegmentCount[v3] {
					part1++
					break
				}
			}
		}
	}

	for k, line := range lines {
		n := 0
		translate := solve(line)
		for _, v := range lineOutputs[k] {
			n = n*10 + translate[v]
		}
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
