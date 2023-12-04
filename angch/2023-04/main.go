package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day4(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	card := 0
	dup := map[int]int{}
	for scanner.Scan() {
		t := scanner.Text()
		card++

		_, r1, _ := strings.Cut(t, ":")
		l, r, _ := strings.Cut(r1, " | ")

		win := map[string]bool{}
		for _, v := range strings.Fields(l) {
			win[v] = true
		}
		count := 0
		for _, v := range strings.Fields(r) {
			if win[v] {
				count++
			}
		}
		part1 += (1 << count) >> 1
		for ; count > 0; count-- {
			dup[card+count] += dup[card] + 1
		}
		part2 += dup[card] + 1
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day4("test.txt")
	if part1 != 13 || part2 != 30 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day4("input.txt"))
}
