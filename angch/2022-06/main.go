package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var verbose = false

func day6(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
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

func main() {
	part1, part2 := day6("test.txt")
	fmt.Println(part1, part2)
	if part1 != 5 || part2 != 23 { // Tested first case only
		log.Fatal("Test failed")
	}
	fmt.Println(day6("input.txt"))
}
