package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filepath = "input.txt"

// var filepath = "test.txt"

type Fish struct {
}

func part1and2() {
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	agesStr := strings.Split(scanner.Text(), ",")

	ageCounts := [9]int{}
	total := len(agesStr)
	for _, age := range agesStr {
		a, _ := strconv.Atoi(age)
		ageCounts[a]++
	}

	index := 0
	for day := 0; day < 256; {
		zero := ageCounts[index]
		index++
		if index >= 9 {
			index = 0
		}
		ageCounts[(6+index)%9] += zero
		day++
		total += zero

		if day == 80 {
			fmt.Println("Part 1", total)
		}
	}

	fmt.Println("Part 2", total)
}

func main() {
	part1and2()
}
