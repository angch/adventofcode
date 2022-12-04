package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day17(file string, target int) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	weights := []int{}
	for scanner.Scan() {
		t := scanner.Text()
		a, _ := strconv.Atoi(t)
		weights = append(weights, a)
	}

	max := 1 << len(weights)
	bestcount := 9999

	for i := 0; i < max; i++ {
		sum, count := 0, 0
		for j, mask := 0, 1; mask < max; mask <<= 1 {
			if (mask & i) != 0 {
				sum += weights[j]
				count++
			}
			j++
		}
		if sum == target {
			part1++
			if bestcount > count {
				bestcount = count
				part2 = 1
			} else if bestcount == count {
				part2++
			}
		}
	}

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(day17("test.txt", 25))
	fmt.Println(day17("input.txt", 150))
}
