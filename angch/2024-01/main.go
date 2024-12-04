package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func day1(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	one := []int{}
	two := []int{}
	count := map[int]int{}

	for scanner.Scan() {
		t := scanner.Text()
		n1, n2 := 0, 0
		_, err := fmt.Sscanf(t, "%d %d", &n1, &n2)
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
		one = append(one, n1)
		two = append(two, n2)
		count[n2]++
	}
	sort.Ints(one)
	sort.Ints(two)
	for k, v := range one {
		d := two[k] - v
		if d < 0 {
			d = -d
		}
		part1 += d
		part2 += v * count[v]
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day1("test.txt")
	if part1 != 11 || part2 != 31 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day1("input.txt"))
}
