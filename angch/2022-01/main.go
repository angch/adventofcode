package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func day1(file string) (int, int) {
	count1, count2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	elf := make([]int, 0)
	cal := 0
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			elf = append(elf, -cal)
			cal = 0
			continue
		}
		c, _ := strconv.Atoi(t)
		cal += c
	}
	sort.Ints(elf)
	count1 = elf[0]
	count2 = elf[0] + elf[1] + elf[2]
	return -count1, -count2
}

func main() {
	fmt.Println(day1("test.txt"))
	fmt.Println(day1("input.txt"))
}
