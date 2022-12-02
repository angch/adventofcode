package main

import (
	"bufio"
	"fmt"
	"log"
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
	if cal > 0 {
		elf = append(elf, -cal)
	}
	sort.Ints(elf)
	count1 = elf[0]
	count2 = elf[0] + elf[1] + elf[2]
	return -count1, -count2
}

func main() {
	count1, count2 := day1("test.txt")
	if count1 != 24000 || count2 != 45000 {
		log.Fatal("Test failed")
	}
	fmt.Println(count1, count2)
	fmt.Println(day1("input.txt"))
}
