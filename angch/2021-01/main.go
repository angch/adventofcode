package main

import (
	"bufio"
	"fmt"
	"os"
)

const p1, p2 = 2, 4

func main() {
	file, _ := os.Open("test.txt")
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
	fmt.Println(part1, part2)
}
