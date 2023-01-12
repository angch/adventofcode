package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var verbose = false

func day6io(f io.Reader) (int, int) {
	part1, part2 := 0, 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()

		seen := make([]byte, 0, 14)
		for i := 0; i < len(t); i++ {
			if len(seen) >= 4 {
				k := make(map[byte]bool)
				for _, v := range seen[len(seen)-4:] {
					k[v] = true
				}
				if len(k) == 4 && part1 == 0 {
					part1 = i
				}
			}
			if len(seen) == 14 {
				k := make(map[byte]bool)
				for _, v := range seen {
					k[v] = true
				}
				if len(k) == 14 && part2 == 0 {
					part2 = i
					break
				}
			}
			seen = append(seen, t[i])
			if len(seen) > 14 {
				seen = seen[1:]
			}
		}
	}

	return part1, part2 // 635 not
}
func day6io2(f io.Reader) (int, int) {
	part1, part2 := 0, 0

	t, _ := io.ReadAll(f)
	count1 := make([]int, 26)
	count1overflow := 0
	count2 := make([]int, 26)
	count2overflow := 0

	for i := 0; i < len(t); i++ {
		c := t[i] - 'a'

		// Part1
		if part1 == 0 {
			count1[c]++
			if count1[c] == 2 {
				count1overflow++
			}

			if i >= 4 {
				c2 := t[i-4] - 'a'
				count1[c2]--
				if count1[c2] == 1 {
					count1overflow--
				}
			}
			if i >= 3 && count1overflow <= 0 {
				part1 = i + 1
			}
		}

		// Part2
		count2[c]++
		if count2[c] == 2 {
			count2overflow++
		}
		if i >= 14 {
			c2 := t[i-14] - 'a'
			count2[c2]--
			if count2[c2] == 1 {
				count2overflow--
			}
		}
		if i >= 13 && count2overflow <= 0 {
			return part1, i + 1
		}
	}
	return part1, part2
}

func day6(file string) (int, int) {
	f, _ := os.Open(file)
	defer f.Close()
	return day6io2(f)
}

func main() {
	part1, part2 := day6("test.txt")
	fmt.Println(part1, part2)
	if part1 != 5 || part2 != 23 { // Tested first case only
		log.Fatal("Test failed")
	}
	fmt.Println(day6("input.txt"))
}
