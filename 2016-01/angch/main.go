package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Matrix struct {
	a, b, c, d int
}

//left=dir = Coord{dir.y, -dir.x, 0}
// See https://en.wikipedia.org/wiki/Rotation_matrix#Common_rotations
var directions = map[string]Matrix{
	"L": Matrix{0, -1, 1, 0},
	"R": Matrix{0, 1, -1, 0},
}

func (c *Coord) mult(m Matrix) {
	c.x, c.y = c.x*m.a+c.y*m.b, c.x*m.c+c.y*m.d
}

func (c *Coord) add(d Coord, m int) {
	c.x += d.x * m
	c.y += d.y * m
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (c *Coord) manhattan() int {
	return abs(c.x) + abs(c.y)
}

func do(input string) (int, int) {
	vector := Coord{0, 1} // North, we're using Cartesian
	pos := Coord{0, 0}

	history := make(map[Coord]bool)

	moves := strings.Split(input, ",")
	revisited := Coord{0, 0}
	done := false
	for _, m := range moves {
		move := strings.Trim(m, " ")
		d := string(move[0])
		n, _ := strconv.Atoi(move[1:])
		//fmt.Println(move, d, n)

		matrix := directions[d]
		vector.mult(matrix)

		for nn := 0; nn < n; nn++ {
			pos.add(vector, 1)
			ok := history[pos]
			if ok {
				if !done {
					revisited = pos
					done = true
				}
			} else {
				history[pos] = true
			}
		}
		//fmt.Println(pos)
	}
	return pos.manhattan(), revisited.manhattan()
}

func main() {
	fmt.Println(do("R2, L3"))
	fmt.Println(do("R2, R2, R2"))
	fmt.Println(do("R5, L5, R5, R3"))
	fmt.Println(do("R8, R4, R4, R8"))
	if true {
		input, _ := ioutil.ReadFile("input.txt")
		fmt.Println(do(string(input)))
	}
}
