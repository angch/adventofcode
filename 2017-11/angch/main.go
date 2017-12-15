package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Coord struct {
	x int
	y int
	z int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func absup(a int) int {
	a = abs(a)
	if a&1 == 1 {
		return a / 2
	}
	return a / 2
}
func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dist(coord Coord) int {
	return (abs(coord.x) + abs(coord.y) + abs(coord.z)) / 2
}
func main() {
	coord := Coord{}

	input := "ne,ne,ne"
	//input = "ne,ne,sw,sw"
	input = "ne,ne,s,s"
	//input = "se,sw,se,sw,sw"
	input = ""
	input = "ne,se"
	if true {
		inputb, _ := ioutil.ReadFile("input.txt")
		input = string(inputb)
	}

	moves := strings.Split(input, ",")
	max := -1
	for _, m := range moves {
		switch m {
		case "nw":
			coord.x--
			coord.y++
		case "n":
			coord.y++
			coord.z--
		case "ne":
			coord.x++
			coord.z--
		case "sw":
			coord.x--
			coord.z++
		case "s":
			coord.y--
			coord.z++
		case "se":
			coord.x++
			coord.y--
		}
		if dist(coord) > max {
			max = dist(coord)
		}
	}
	fmt.Println(coord)
	fmt.Println(dist(coord))
	fmt.Println(max)
}
