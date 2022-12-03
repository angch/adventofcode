package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func tr(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 'a' + 1
	}
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 1 + 26
	}
	log.Fatal(b)
	return 0
}

func day3(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	group := map[byte]byte{}
	g := 2
a:
	for scanner.Scan() {
		g = (g + 1) % 3
		if g == 0 {
			// We looped back to the first member in the group, so we can look for
			// the group's dupes
			for k, v := range group {
				if v == 7 {
					part2 += int(k)
					break
				}
			}
			group = map[byte]byte{}
		}

		t := scanner.Text()

		for i := 0; i < len(t); i++ {
			p := group[tr(t[i])]
			p |= 1 << g
			group[tr(t[i])] = p
		}

		c := make(map[byte]bool)
		for i := 0; i < len(t)/2; i++ {
			c[tr(t[i])] = true
		}
		for i := len(t) / 2; i < len(t); i++ {
			pr := tr(t[i])
			if c[pr] {
				part1 += int(pr)
				continue a
			}
		}
	}
	// Check the last group
	for k, v := range group {
		if v == 7 {
			part2 += int(k)
			break
		}
	}

	return part1, part2
}

func main() {
	part1, part2 := day3("test.txt")
	fmt.Println(part1, part2)
	if part1 != 157 || part2 != 70 {
		log.Fatal("Test failed")
	}
	fmt.Println(day3("input.txt"))
}
