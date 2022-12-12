package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Note: tr is inlined by Go. (canInlineFunction(cost: 18)optimizer details)
func tr(b byte) byte {
	if b <= 'Z' {
		return b - 'A' + 1 + 26
	}
	return b - 'a' + 1
}

func day3(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	emptybool, emptygroup := [53]bool{}, [53]byte{}
	g, group := byte(1), emptygroup
	for scanner.Scan() {
		t := scanner.Text()

		// Part 2. Note "g" is a bitmask
		if g != 4 {
			for _, c := range t {
				group[tr(byte(c))] |= g
			}
			g <<= 1
		} else {
			for _, c := range t {
				pr := tr(byte(c))
				if group[pr] == g-1 {
					part2 += int(pr)
					break
				}
			}
			g, group = 1, emptygroup
		}

		// Part 1
		c := emptybool
		for i := 0; i < len(t)/2; i++ {
			c[tr(t[i])] = true
		}
		for i := len(t) / 2; i < len(t); i++ {
			pr := tr(t[i])
			if c[pr] {
				part1 += int(pr)
				break
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
