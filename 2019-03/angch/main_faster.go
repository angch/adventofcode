/*
cybus:angch angch$ go build main_faster.go
cybus:angch angch$ time ./main_faster
159 610
135 410
2427 27890

real	0m0.982s
user	0m0.268s
sys	0m0.484s
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "runtime"
	"strings"
)

var max int = 40000

var directions map[byte][]int = map[byte][]int{
	'R': []int{1, 0},
	'L': []int{-1, 0},
	'U': []int{0, -1},
	'D': []int{0, 1},
}

func trace(line string, color int8, board []int8, steps []uint32) (int, int) {
	ox, oy := max/2, max/2
	points := strings.Split(line, ",")
	x, y := ox, oy
	mindist := 999999
	minsteps := 999999
	step := 0
	for _, p := range points {
		dx, dy := directions[p[0]][0], directions[p[0]][1]
		l, _ := strconv.Atoi(p[1:])

		for ; l > 0; l-- {
			if color == 2 {
				steps[y*max+x] = uint32(step)
				board[y*max+x] = color
			}
			x += dx
			y += dy
			step++

			if board[y*max+x] > color {
				// nonself intersects
				dist := 0
				if y > oy {
					dist = y - oy
				} else {
					dist = oy - y
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
					if int(steps[y*max+x])+step < minsteps {
						minsteps = int(steps[y*max+x]) + step
					}
				}
			}
		}
	}
	return mindist, minsteps
}

func partboth(one string, two string) (int, int) {
	board := make([]int8, max*max)
	steps := make([]uint32, max*max)

	// ox, oy := max/2, max/2
	trace(one, 2, board, steps)
	return trace(two, 1, board, steps)
}

func main() {

	if true {
		go func() {
			dist, step := partboth(
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			)
			fmt.Println(dist, step)
		}()

		go func() {
			dist, step := partboth(
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
			fmt.Println(dist, step)
			// log.Println("GC")
			// runtime.GC()
			// log.Println("/GC")
		}()
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
