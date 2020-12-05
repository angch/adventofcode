package main

import (
	"bufio"
	"log"
	"os"
)

type Range struct {
	Min int
	Max int
}

func do(input string) int {
	row := Range{0, 128}
	col := Range{0, 8}
	// log.Println(row)
	for _, v := range input {
		// log.Println(row, col, k, v)
		switch v {
		case 'F':
			row.Max -= (row.Max - row.Min) / 2
		case 'B':
			row.Min += (row.Max - row.Min) / 2
		case 'L':
			col.Max -= (col.Max - col.Min) / 2
		case 'R':
			col.Min += (col.Max - col.Min) / 2
		}
	}
	// log.Println(row.Min, col.Min)
	return row.Min*8 + col.Min
}
func main() {
	// do("FBFBBFFRLR")
	// do("BFFFBBFRRR")
	// do("FFFBBBFRRR")
	// do("BBFFBBFRLL")

	max := 0
	// inputs := make([]string, 0)
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ids := make([]bool, 1000)
	min := 999999
	for scanner.Scan() {
		l := scanner.Text()
		id := do(l)
		if id > max {
			max = id
		}
		if id < min {
			min = id
		}
		ids[id] = true
		// log.Println(id)
	}
	log.Println(max, min)

	for i := min; i < max-1; i++ {
		if ids[i-1] && !ids[i] && ids[i+1] {
			log.Println("id is", i)
			break
		}
	}
}
