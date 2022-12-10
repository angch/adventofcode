package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day10(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	X, cycle, prev := []int{1}, 0, 1
	record := []int{20, 60, 100, 140, 180, 220}

	for scanner.Scan() {
		t := scanner.Text()
		words := strings.Split(t, " ")
		switch words[0] {
		case "noop":
			cycle += 1
			X = append(X, prev)
		case "addx":
			cycle += 2
			n, _ := strconv.Atoi(words[1])
			X = append(X, prev, prev+n)
			prev += n
		}
	}
	for _, v := range record {
		part1 += v * X[v-1]
	}

	for i := 0; i < 240; i++ {
		x := i % 40
		diff := X[i] - x
		if diff > -2 && diff < 2 {
			fmt.Print("##")
		} else {
			fmt.Print("..")
		}
		if x == 39 {
			fmt.Println()
		}
	}
	return part1, part2
}

func main() {
	part1, part2 := day10("test2.txt")
	fmt.Println(part1, part2)
	fmt.Println(day10("input.txt"))
}
