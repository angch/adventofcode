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

func isSafe(row []int, skip int) bool {
	prev, dir := -100, 0
	for k, v := range row {
		if k == skip {
			continue
		}
		if prev == -100 {
			prev = v
			continue
		}
		if dir == 0 {
			if prev > v {
				dir--
			} else {
				dir++
			}
		}
		d := (v - prev) * dir
		if d > 3 || d <= 0 {
			return false
		}
		prev = v
	}
	return true
}

func day2(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		row := []int{}
		for _, v := range strings.Fields(t) {
			i, _ := strconv.Atoi(v)
			row = append(row, i)
		}
		safe := isSafe(row, -1)
		if safe {
			part1++
			part2++
		} else {
			for k := range len(row) {
				safe := isSafe(row, k)
				if safe {
					part2++
					break
				}
			}
		}
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day2("test.txt")
	fmt.Println(part1, part2)
	if part1 != 2 || part2 != 4 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day2("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
