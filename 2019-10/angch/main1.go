package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
)

type Coord struct {
	x, y int
}

func CountSeen(points map[Coord]bool, point Coord) int {
	seen := 0
	dir := make(map[float64]bool)
	for k := range points {
		dx := float64(k.x - point.x)
		dy := float64(k.y - point.y)

		if dx == 0.0 && dy == 0.0 {
			continue
		}
		d := math.Atan2(dy, dx)
		// length :=ath.Sqrt(float64(dx*dx + dy*dy))

		if dir[d] {
			continue
		}
		seen++
		dir[d] = true
	}
	return seen
}

func dofile(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := make(map[Coord]bool)
	y := 0
	grid := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		for x, cell := range line {
			if cell == '#' {
				points[Coord{x, y}] = true
			}
		}
		row := make([]int, len(line))
		grid = append(grid, row)
		y++
	}

	debug := false
	if debug {
		fmt.Printf("%+v\n", points)
	}

	countChan := make(chan int)
	wg := sync.WaitGroup{}
	for k := range points {
		wg.Add(1)
		go func(k Coord) {
			count := CountSeen(points, k)
			// fmt.Println(k, count)
			if debug {
				grid[k.y][k.x] = count
			}
			countChan <- count
			wg.Done()
		}(k)
	}
	go func() {
		wg.Wait()
		close(countChan)
	}()
	best := 0
	for c := range countChan {
		if c > best {
			best = c
		}
	}
	fmt.Println("Best", fileName, "is", best)

	if debug {
		for _, row := range grid {
			for _, c := range row {
				fmt.Print(c)
			}
			fmt.Println()
		}
	}

}

func main() {
	dofile("test.txt")  // 8
	dofile("test2.txt") // 33
	dofile("test3.txt") // 41
	dofile("test4.txt") // 210
	dofile("input.txt") //
}
