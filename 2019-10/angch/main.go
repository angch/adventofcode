package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

		if dir[d] {
			continue
		}
		seen++
		dir[d] = true
	}
	return seen
}

type vectorDist struct {
	x, y     int
	distance float64
}
type vectorDists []vectorDist

func (s vectorDists) Len() int      { return len(s) }
func (s vectorDists) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s vectorDists) Less(i, j int) bool {
	return s[i].distance < s[j].distance
}

func Zappy(points map[Coord]bool, point Coord) int {
	fmt.Println("Zapping", point.x, point.y)
	dir := make(map[float64]vectorDists)
	dirs := make([]float64, 0)
	debug := false
	for k := range points {
		dx := float64(point.x - k.x)
		dy := float64(point.y - k.y)
		distance := math.Sqrt(dx*dx + dy*dy)

		if dx == 0.0 && dy == 0.0 {
			continue
		}

		// clockwise now, in degrees. Hack, hack
		direction := math.Atan2(dx, dy) * 180 / math.Pi
		if direction < 0 {
			direction += 360
		}
		direction = -direction
		if direction < 0 {
			direction += 360
		}
		if k.x == 11 && k.y == 12 {
			fmt.Println("xxx", dx, dy, distance, direction)
		}

		_, exists := dir[direction]
		if !exists {
			dir[direction] = make(vectorDists, 0)
			dirs = append(dirs, direction)
		}
		dir[direction] = append(dir[direction], vectorDist{k.x, k.y, distance})
	}
	sort.Float64s(dirs)
	fmt.Println(dirs)
	for k := range dir {
		if debug {
			fmt.Printf("Before %+v\n", dir[k])
		}
		sort.Sort(vectorDists(dir[k]))
		if debug {
			fmt.Printf("After %+v\n", dir[k])
		}
	}

	currentDirIdx := 0
	for zapped := 0; zapped <= 200; zapped++ {
		currentDir := dirs[currentDirIdx]
		for {
			if len(dir[currentDir]) > 0 {
				break
			}
			// fmt.Println("skip", dir[currentDir])
			currentDirIdx = (currentDirIdx + 1) % len(dirs)
			currentDir = dirs[currentDirIdx]
		}

		if zapped == 199 {
			fmt.Println("200", dir[currentDir][0])
			break
		}
		// Zap!
		fmt.Println(zapped, currentDirIdx, currentDir, dir[currentDir][0])
		dir[currentDir] = dir[currentDir][1:]
		currentDirIdx = (currentDirIdx + 1) % len(dirs)
	}
	// 	seen--
	// }

	return 0
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

	type countVec struct {
		count int
		coord Coord
	}
	countChan := make(chan countVec)
	wg := sync.WaitGroup{}
	for k := range points {
		wg.Add(1)
		go func(k Coord) {
			count := CountSeen(points, k)
			if debug {
				grid[k.y][k.x] = count
			}
			countChan <- countVec{count, Coord{k.x, k.y}}
			wg.Done()
		}(k)
	}
	go func() {
		wg.Wait()
		close(countChan)
	}()
	best := 0
	bestCoord := Coord{-1, -1}
	for c := range countChan {
		if c.count > best {
			best = c.count
			bestCoord = c.coord
		}
	}
	fmt.Println("Best", fileName, "is", best, "at", bestCoord)

	Zappy(points, bestCoord)

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
	// dofile("test.txt")  // 8
	// dofile("test2.txt") // 33
	// dofile("test3.txt") // 41
	// dofile("test4.txt") // 210
	dofile("input.txt") // 13 19 too high
}
