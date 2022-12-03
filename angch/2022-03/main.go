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
	emptygroup := [53]byte{}
	emptybool := [53]bool{}
	group := emptygroup
	g := 0
a:
	for scanner.Scan() {
		t := scanner.Text()
		if g != 2 {
			mask := byte(1 << g)
			for _, c := range t {
				group[tr(byte(c))] |= mask
			}
			g++
		} else {
			for _, c := range t {
				pr := tr(byte(c))
				if group[pr] == 3 {
					part2 += int(pr)
					break
				}
			}
			group = emptygroup
			g = 0
		}

		c := emptybool
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
