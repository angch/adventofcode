package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day6(filepath string) {
	file, _ := ioutil.ReadFile(filepath)
	agesStr := strings.Split(string(file), ",")

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
	day6("test.txt")
	day6("input.txt")
}
