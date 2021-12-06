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

	for day := 0; day < 256; day++ {
		if day == 80 {
			fmt.Println("Part 1", total)
		}
		total += ageCounts[day%9]
		ageCounts[(7+day)%9] += ageCounts[day%9]
	}
	fmt.Println("Part 2", total)
}

func main() {
	part1and2()
}
