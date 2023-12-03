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

func day3(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	board := make(map[Coord]byte)
	strboard := make(map[Coord]string)
	gears := []Coord{}

	row := 0
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
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
		// log.Println(len(strboard))
	}
a:
	for k, v := range strboard {
		// log.Println(k, v)
		x1 := k.X - 1
		x2 := k.X + len(v)
		i, _ := strconv.Atoi(v)
		part1 += i

		debug := false
		if v == "617" {
			debug = true
			log.Println("yes", x1, x2, k.Y)
		}
		for x := x1; x <= x2; x++ {
			// log.Println("xx", x)
			if board[Coord{x, k.Y - 1}] != 0 {
				if debug {
					log.Println(k)
				}
				continue a
			}
			if board[Coord{x, k.Y + 1}] != 0 {
				if debug {
					log.Println(k, string(board[Coord{x, k.Y + 1}]))
				}
				continue a
			}
			if debug {
				// log.Println("y")
			}
		}
		if board[Coord{x1, k.Y}] != 0 {
			if debug {
				log.Println("1", k)
			}
			continue a
		}
		if board[Coord{x2, k.Y}] != 0 {
			if debug {
				log.Println("2", k)
			}
			continue a
		}
		if debug {
			// log.Fatal("x")
		}
		log.Println("skip", i)
		part1 -= i
	}

	for _, gearCoord := range gears {
		nums := []int{}
		for k, v := range strboard {
			dy := max(gearCoord.Y-k.Y, k.Y-gearCoord.Y)
			if dy > 1 {
				continue
			}

			if gearCoord.X+1 < k.X {
				continue
			}
			if gearCoord.X > k.X+len(v) {
				continue
			}
			i, _ := strconv.Atoi(v)
			nums = append(nums, i)
		}
		if len(nums) > 1 {
			part2 += nums[0] * nums[1]
		}
	}

	log.Println(strboard, part1)
	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day3("test.txt")
	if part1 != 4361 || part2 != 467835 {
		log.Fatal("Test failed ", part1, part2)
	}
	// 562097 too high
	fmt.Println(day3("input.txt"))
}
