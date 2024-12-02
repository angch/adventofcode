package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isafe(row []int, skip int) bool {
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
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		row := []int{}
		strs := strings.Split(t, " ")
		for _, v := range strs {
			i, _ := strconv.Atoi(v)
			row = append(row, i)
		}
		safe := isafe(row, -1)
		if safe {
			part1++
			part2++
		} else {
			for k := range len(row) {
				safe := isafe(row, k)
				if safe {
					part2++
					break
				}
			}
		}
	}
	// log.Println(rows)
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day2("test.txt")
	fmt.Println(part1, part2)
	if part1 != 2 || part2 != 4 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day2("input.txt"))
}
