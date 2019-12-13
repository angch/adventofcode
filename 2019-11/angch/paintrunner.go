package main

import (
	"fmt"
	"log"
)

type Coord struct {
	x, y int
}

func PaintRunner(prog map[int]int, init int) {
	output := make(chan int)
	input := make(chan int, 1)

	go IntCodeVM(prog, input, output)

	grid := make(map[Coord]int)
	curloc := Coord{0, 0}
	curdir := Coord{0, -1} // Up

	globalpainted := 0

	minloc, maxloc := Coord{0, 0}, Coord{0, 0}

	if init > 0 {
		grid[curloc] = init
	}
	for {
		paintold, painted := grid[curloc]
		input <- paintold
		paintnew, more := <-output
		if !more {
			break
		}

		grid[curloc] = paintnew
		if !painted {
			globalpainted++
		}

		turn := <-output
		if turn == 0 {
			curdir.x, curdir.y = curdir.y, -curdir.x
		} else {
			curdir.x, curdir.y = -curdir.y, curdir.x
		}
		curloc.x, curloc.y = curloc.x+curdir.x, curloc.y+curdir.y
		if minloc.x > curloc.x {
			minloc.x = curloc.x
		}
		if minloc.y > curloc.y {
			minloc.y = curloc.y
		}
		if maxloc.x < curloc.x {
			maxloc.x = curloc.x
		}
		if maxloc.y < curloc.y {
			maxloc.y = curloc.y
		}
	}
	log.Println("Count", globalpainted, minloc, maxloc) // 1236 too low

	for y := minloc.y; y <= maxloc.y; y++ {
		for x := minloc.x; x <= maxloc.x; x++ {
			if grid[Coord{x, y}] == 1 {
				fmt.Print("##")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
