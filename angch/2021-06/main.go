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
	// lines := make([]Line, 0)

	scanner.Scan()
	t := scanner.Text()
	ages := make([]int, 0)
	agesStr := strings.Split(t, ",")

	for _, age := range agesStr {
		a, _ := strconv.Atoi(age)
		ages = append(ages, a)
	}

	if false {
		// Yeah, don't do this
		for day := 0; day < 80; {
			add := 0
			for i := 0; i < len(ages); i++ {
				ages[i]--
				if ages[i] < 0 {
					ages[i] = 6
					add++
				}
			}
			for ; add > 0; add-- {
				ages = append(ages, 8)
			}
			day++
		}
	}

	ageCounts := make([]int, 9)
	for _, age := range ages {
		ageCounts[age]++
	}

	for day := 0; day < 256; {
		zero := ageCounts[0]
		ageCounts = ageCounts[1:]
		ageCounts[6] += zero
		ageCounts = append(ageCounts, zero)
		day++

		if day == 80 {
			total := 0
			for i := 0; i < len(ageCounts); i++ {
				total += ageCounts[i]
			}
			fmt.Println("Part 1", total)
		}
	}
	total := 0
	for i := 0; i < len(ageCounts); i++ {
		total += ageCounts[i]
	}

	fmt.Println("Part 2", total)
}

func main() {
	part1and2()
}
