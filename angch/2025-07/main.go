package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func day7(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	t := scanner.Bytes()
	beams := map[int]int{}
	for x, c := range t {
		if c == 'S' {
			beams[x] = 1
		}
	}

	part2 = 1
	for scanner.Scan() {
		t := scanner.Bytes()
		rest := 0
		for x, c := range t {
			if c == '^' && beams[x] > 0 {
				split := 0
				if x > 0 {
					beams[x-1] += beams[x]
					split += beams[x]
				}
				if x < len(t) {
					beams[x+1] += beams[x]
					split += beams[x]
				}
				delete(beams, x)
				rest++
				part1++
			}
		}
		if rest == 0 {
			sum := 0
			for _, i := range beams {
				sum += i
			}
			// fmt.Println(beams, len(beams), sum)
			part2 = sum
		}
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day7("test.txt")
	fmt.Println(part1, part2)
	if part1 != 21 || part2 != 40 {
		log.Fatal("Test failed ", part1, part2)
	}

	part1, part2 = day7("input.txt")
	fmt.Println(part1, part2)
	fmt.Println("Elapsed time:", time.Since(t1))
}
