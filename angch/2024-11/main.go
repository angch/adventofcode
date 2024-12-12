package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func dostone(stones map[int]int) (map[int]int, int) {
	out := make(map[int]int)
	total := 0
	for stone, count := range stones {
		total += count

		if stone == 0 {
			out[1] += count
			continue
		}

		ten := fmt.Sprintf("%d", stone)
		isEvenTen := len(ten)%2 == 0
		if isEvenTen {
			l, _ := strconv.Atoi(ten[:len(ten)/2])
			r, _ := strconv.Atoi(ten[len(ten)/2:])
			out[l] += count
			out[r] += count
			total += count
			continue
		}
		out[stone*2024] += count
	}
	return out, total
}

func day11(file string, parttwo bool) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	stones := make(map[int]int)
	for scanner.Scan() {
		for _, v := range strings.Fields(scanner.Text()) {
			i, _ := strconv.Atoi(v)
			stones[i]++
		}
	}

	count := 0
	for times := range 25 {
		stones, count = dostone(stones)
		if times == 24 {
			part1 = count
		}
	}

	for times := range 75 - 25 {
		stones, count = dostone(stones)
		if times == 75-25-1 {
			part2 = count
		}
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day11("test.txt", false)
	fmt.Println(part1, part2)
	if part1 != 55312 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day11("input.txt", true))
	fmt.Println("Elapsed time:", time.Since(t1))
}
