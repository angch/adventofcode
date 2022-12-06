package main

import (
	"bufio"
	"fmt"
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

		seen := make([]byte, 0, 4)
		seen2 := make([]byte, 0, 14)
		for i := 0; i < len(t); i++ {
			if len(seen) == 4 {
				k := make(map[byte]bool)
				for _, v := range seen {
					k[v] = true
				}
				if len(k) == 4 && part1 == 0 {
					// fmt.Println("part1", i)
					part1 = i
					// break

				}
			}
			if len(seen2) == 14 {
				k := make(map[byte]bool)
				for _, v := range seen2 {
					k[v] = true
				}
				if len(k) == 14 && part2 == 0 {
					// fmt.Println("part2", i)
					part2 = i
					break
				}
			}

			seen = append(seen, t[i])
			seen2 = append(seen2, t[i])
			if len(seen) > 4 {
				seen = seen[1:]
			}
			if len(seen2) > 14 {
				seen2 = seen2[1:]
			}
		}
	}

	return part1, part2 // 635 not
}

func main() {
	part1, part2 := day6("test.txt")
	fmt.Println(part1, part2)
	// if part1 != "CMZ" || part2 != "MCD" {
	// 	log.Fatal("Test failed")
	// }
	fmt.Println(day6("input.txt"))
}
