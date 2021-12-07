package main

import (
	"bufio"
	"fmt"
	"os"
)

const p1, p2 = 2, 4

func day1(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var part1, part2 int
	var win []int
	for scanner.Scan() {
		var n int
		fmt.Sscanf(scanner.Text(), "%d", &n)
		win = append(win, n)
		if len(win) >= p1 {
			if n > win[len(win)-p1] {
				part1++
			}
		}
		if len(win) >= p2 {
			if n > win[len(win)-p2] {
				part2++
			}
			win = win[1:]
		}
	}
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	day1("test.txt")
	day1("input.txt")
}
