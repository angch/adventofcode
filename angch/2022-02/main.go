package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day2(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		lr := strings.Split(t, " ")
		l, r := lr[0], lr[1]
		r0 := int(r[0] - 'X')
		l0 := int(l[0] - 'A')

		// Part 1
		score1 := (l0+r0)%3*3 + r0 + 1
		part1 += score1

		// Part 2
		i := r0 * 3
		part2 += i + (l0+r0)%3
	}
	return part1, part2
}

func main() {
	part1, part2 := day2("test.txt")
	fmt.Println(part1, part2)
	if part1 != 15 || part2 != 12 {
		log.Fatal("Test failed")
	}
	fmt.Println(day2("input.txt"))
}
