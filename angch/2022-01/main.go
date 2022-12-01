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
		_ = t
		if t == "" {
			elf = append(elf, cal)
			cal = 0
			continue
		}
		c, _ := strconv.Atoi(t)
		cal += c

		// c1, c2 := countLit(t)
		// fmt.Println(t, c2, c1)
		// count1 += c1 - c2
		// c1, c2 = countEnc(t)
		// count2 += c2 - c1
	}
	sort.Ints(elf)
	count1 = elf[len(elf)-1]
	count2 = elf[len(elf)-1]
	count2 += elf[len(elf)-2]
	count2 += elf[len(elf)-3]

	return count1, count2
}

func main() {
	fmt.Println(day1("input.txt"))
}
