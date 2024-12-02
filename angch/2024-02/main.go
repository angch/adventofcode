package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isafe(row []int) bool {
	prev := -100
	safe := true
	dir := -10
	for _, v := range row {
		if prev == -100 {
			prev = v
			continue
		}
		if dir == -10 {
			if prev > v {
				dir = 1
			} else {
				dir = -1
			}
		}
		d := (v - prev) * -dir
		if d > 3 || d <= 0 {
			safe = false
		}
		prev = v
	}
	return safe

}

func day2(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	rows := [][]int{}
	for scanner.Scan() {
		t := scanner.Text()
		row := []int{}
		strs := strings.Split(t, " ")
		for _, v := range strs {
			i, _ := strconv.Atoi(v)
			row = append(row, i)
		}
		safe := isafe(row)
		if safe {
			log.Println("safe")
			part1++
			part2++
		} else {
			for k := 0; k < len(row); k++ {
				row1 := make([]int, len(row))
				copy(row1, row)
				// remove k
				row1 = append(row1[:k], row1[k+1:]...)
				safe := isafe(row1)
				if safe {
					log.Println("part2 safe")
					part2++
					break
				}
			}
			log.Println("part2 unsafe")
		}

		rows = append(rows, row)
	}
	// log.Println(rows)
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day2("test.txt")
	log.Println(part1)
	if part1 != 2 || part2 != 4 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day2("input.txt"))
}
