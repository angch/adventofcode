package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var verbose = false

func day5(file string) (string, string) {
	part1, part2 := "", ""
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	stacks := make([][]rune, 0, 0)
	max := 0
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}

		if max == 0 {
			max = (len(t) + 1) / 4
			stacks = make([][]rune, max)
			for i := 0; i < max; i++ {
				stacks[i] = make([]rune, 0)
			}
		}
		for i := 0; i < max; i++ {
			if t[i*4+0] == '[' {
				c := rune(t[i*4+1])
				stacks[i] = append([]rune{c}, (stacks[i])...)
			}
		}
	}

	stacks2 := make([][]rune, max)
	for i := 0; i < max; i++ {
		stacks2[i] = make([]rune, len(stacks[i]))
		copy(stacks2[i], stacks[i])
	}
	if verbose {
		for i := 0; i < max; i++ {
			fmt.Println(i, string(stacks[i]))
		}
	}
	for scanner.Scan() {
		t := scanner.Text()
		n, from, to := 0, 0, 0
		fmt.Sscanf(t, "move %d from %d to %d", &n, &from, &to)
		from--
		to--
		for i := n; i > 0; i-- {
			p := stacks[from][len(stacks[from])-1]
			stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to] = append(stacks[to], p)
		}

		p := stacks2[from][len(stacks2[from])-n : len(stacks2[from])]
		stacks2[from] = stacks2[from][:len(stacks2[from])-n]
		stacks2[to] = append(stacks2[to], p...)

		if verbose {
			fmt.Println(t)
			for i := 0; i < max; i++ {
				fmt.Println(i, string(stacks2[i]))
			}
		}
	}

	for i := 0; i < max; i++ {
		if len(stacks[i]) > 0 {
			part1 += string(stacks[i][len(stacks[i])-1])
		}
		if len(stacks2[i]) > 0 {
			part2 += string(stacks2[i][len(stacks2[i])-1])
		}
	}
	return part1, part2 // 635 not
}

func main() {
	part1, part2 := day5("test.txt")
	fmt.Println(part1, part2)
	if part1 != "CMZ" || part2 != "MCD" {
		log.Fatal("Test failed")
	}
	fmt.Println(day5("input.txt"))
}
