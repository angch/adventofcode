package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// var part1, part2 int
	// var win []int
	x, y, z := 0, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		n := 0
		fmt.Sscanf(s[1], "%d", &n)
		fmt.Println(s[0], n)
		switch s[0] {
		case "forward":
			x += n
		case "down":
			z += n
		case "up":
			z -= n
		}

	}
	fmt.Println(x, y, z, x*z)
}

func part2() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// var part1, part2 int
	// var win []int
	x, y, z := 0, 0, 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		n := 0
		fmt.Sscanf(s[1], "%d", &n)
		fmt.Println(s[0], n)
		switch s[0] {
		case "forward":
			x += n
			z += aim * n
		case "down":
			// z += n
			aim += n
		case "up":
			// z -= n
			aim -= n
		}

	}
	fmt.Println(x, y, z, x*z)
}
func main() {
	part2()

}
