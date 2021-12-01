package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	var part1, part2 int
	win := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		win = append(win, n)
		if len(win) >= 2 {
			if n > win[len(win)-2] {
				part1++
			}
		}
		if len(win) >= 4 {
			a := win[len(win)-4]
			if n > a {
				part2++
			}
		}
	}
	fmt.Println(part1, part2)
}
