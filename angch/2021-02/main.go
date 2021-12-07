package main

import (
	"bufio"
	"fmt"
	"os"
)

func day2(filepath string) {
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	x, z1, z2, aim := 0, 0, 0, 0

	for scanner.Scan() {
		dir, n := "", 0
		fmt.Sscanf(scanner.Text(), "%s %d", &dir, &n)
		switch dir {
		case "forward":
			x += n
			z2 += aim * n
		case "down":
			z1 += n
			aim += n
		case "up":
			z1 -= n
			aim -= n
		}
	}
	fmt.Println("Part 1", x*z1)
	fmt.Println("Part 2", x*z2)
}

func main() {
	day2("test.txt")
	day2("input.txt")
}
