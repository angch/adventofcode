package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Coord struct {
	X, Y int
}

func day3(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	board := make(map[Coord]byte)
	strboard := make(map[Coord]string)
	intboard := make(map[Coord]int)
	gears := []Coord{}

	row := 0
	for scanner.Scan() {
		t := scanner.Text()
		isDigit := false
		row++
		startk := 0
		for k, v := range t {
			if v == '.' {
				v = 0
			}
			board[Coord{k, row}] = byte(v)

			if v >= '0' && v <= '9' {
				if !isDigit {
					isDigit = true
					startk = k
				}
				strboard[Coord{startk, row}] += string(v)
			} else {
				if v == '*' {
					gears = append(gears, Coord{k, row})
				}
				isDigit = false
			}
		}
	}
a:
	for k, v := range strboard {
		x1, x2 := k.X-1, k.X+len(v)
		i, _ := strconv.Atoi(v)
		intboard[k] = i
		part1 += i

		for x := x1; x <= x2; x++ {
			// Scan the top and bottom neighbours, skipping when we find nonempty
			if board[Coord{x, k.Y - 1}] != 0 || board[Coord{x, k.Y + 1}] != 0 {
				continue a
			}
		}
		// Test left and right of the current row
		if board[Coord{x1, k.Y}] != 0 || board[Coord{x2, k.Y}] != 0 {
			continue
		}
		part1 -= i
	}

	for _, gearCoord := range gears {
		nums := []int{}
		for k, v := range strboard {
			dy := max(gearCoord.Y-k.Y, k.Y-gearCoord.Y)
			if dy > 1 {
				continue
			}
			if gearCoord.X+1 < k.X || gearCoord.X > k.X+len(v) {
				continue
			}
			nums = append(nums, intboard[k])
		}
		if len(nums) > 1 {
			part2 += nums[0] * nums[1]
		}
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day3("test.txt")
	if part1 != 4361 || part2 != 467835 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day3("input.txt"))
}
