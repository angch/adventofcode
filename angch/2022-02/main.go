package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day2(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		lr := strings.Split(t, " ")
		l, r := int(lr[0][0]-'A'), int(lr[1][0]-'X')
		part1 += (l+r)%3*3 + r + 1
		part2 += r*3 + (l+r)%3
	}
	return part1, part2
}

func main() {
	part1, part2 := day2("test.txt")
	fmt.Println(part1, part2)
	if part1 != 15 || part2 != 12 {
		log.Fatal("Test failed")
	}
	fmt.Println(day2("input.txt"))
}
