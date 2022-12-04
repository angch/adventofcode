package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day4(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Sscanf(t, "%d-%d,%d-%d", &x1, &x2, &y1, &y2)
		_ = t
		if y2-y1 > x2-x1 {
			x1, x2, y1, y2 = y1, y2, x1, x2
		}
		// log.Println(x1, x2, y1, y2)
		if y2 >= x1 && y1 <= x2 && y2 <= x2 && y1 >= x1 {
			part1++
		}
		if y2 >= x1 && y1 <= x2 {
			part2++
		}
	}
	return part1, part2 // 635 not
}

func main() {
	part1, part2 := day4("test.txt")
	fmt.Println(part1, part2)
	if part1 != 2 || part2 != 4 {
		log.Fatal("Test failed")
	}
	fmt.Println(day4("input.txt"))
}
