package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var max int = 50000

var directions map[string][]int = map[string][]int{
	"R": []int{1, 0},
	"L": []int{-1, 0},
	"U": []int{0, -1},
	"D": []int{0, 1},
}

func trace(line string, color int8, board [][]int8, steps [][]int) (int, int) {
	ox, oy := max/2, max/2
	points := strings.Split(line, ",")
	x, y := ox, oy
	mindist := 999999
	minsteps := 999999
	step := 0
	for _, p := range points {
		dir := string(p[0])
		l := 0
		fmt.Sscanf(p[1:], "%d", &l)

		for ; l > 0; l-- {
			board[y][x] = color
			x += directions[dir][0]
			y += directions[dir][1]
			step++
			if color == 2 {
				steps[y][x] = step
			}

			if board[y][x] > color {
				// intersects
				dist := y - oy
				if dist < 0 {
					dist = -dist
				}
				if x > ox {
					dist += x - ox
				} else {
					dist += ox - x
				}
				if dist < mindist {
					mindist = dist
				}
				if color != 2 {
					if steps[y][x]+step < minsteps {
						minsteps = steps[y][x] + step
					}
				}
				// fmt.Println("i", x, y, dist)
			}
		}
	}
	return mindist, minsteps
}

func partboth(one string, two string) (int, int) {
	board := make([][]int8, max, max)
	steps := make([][]int, max, max)
	for k := range board {
		board[k] = make([]int8, max)
		steps[k] = make([]int, max)
	}

	// ox, oy := max/2, max/2
	trace(one, 2, board, steps)
	return trace(two, 1, board, steps)
}

func main() {
	if true {
		dist, step := partboth(
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
		)
		fmt.Println(dist, step)

		dist, step = partboth(
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
		fmt.Println(dist, step)
	}

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	one := scanner.Text()
	scanner.Scan()
	two := scanner.Text()
	// log.Println(one)
	// log.Println(two)
	dist, step := partboth(one, two)
	fmt.Println(dist, step)
}
