package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var neighbours = [][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func dump(grid [][]byte) {
	for _, v := range grid {
		for _, v2 := range v {
			if v2 == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func step(grid [][]byte, grid2 [][]byte) {
	for y, v := range grid {
		for x, v2 := range v {
			count := 0
			for _, n := range neighbours {
				x1, y1 := x+n[0], y+n[1]
				if x1 >= 0 && x1 < len(grid) && y1 >= 0 && y1 < len(grid) {
					count += int(grid[y1][x1])
				}
			}

			if v2 == 1 {
				if count == 2 || count == 3 {
					grid2[y][x] = 1
				} else {
					grid2[y][x] = 0
				}
			} else {
				if count == 3 {
					grid2[y][x] = 1
				} else {
					grid2[y][x] = 0
				}
			}
		}
	}
	for k := range grid {
		copy(grid[k], grid2[k])
	}
}

func day18(file string, steps int) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]byte, 0)
	for scanner.Scan() {
		t := scanner.Text()
		t2 := make([]byte, len(t))
		for k, v := range []byte(t) {
			if v == '#' {
				t2[k] = 1
			} else {
				t2[k] = 0
			}
		}
		grid = append(grid, t2)
	}
	// dump(grid)

	grid2 := make([][]byte, len(grid))
	grid3 := make([][]byte, len(grid))
	for k, v := range grid {
		grid2[k] = make([]byte, len(v))
		grid3[k] = make([]byte, len(v))
		copy(grid3[k], grid[k])
	}
	for i := 0; i < steps; i++ {
		step(grid, grid2)
		// dump(grid)
	}

	for _, v := range grid {
		for _, v2 := range v {
			part1 += int(v2)
		}
	}

	// for k := range grid {
	// 	copy(grid3[k], grid[k])
	// }
	for i := 0; i < steps; i++ {
		grid3[0][0] = 1
		grid3[0][len(grid)-1] = 1
		grid3[len(grid)-1][0] = 1
		grid3[len(grid)-1][len(grid)-1] = 1
		step(grid3, grid2)
		grid3[0][0] = 1
		grid3[0][len(grid)-1] = 1
		grid3[len(grid)-1][0] = 1
		grid3[len(grid)-1][len(grid)-1] = 1
		// dump(grid)
	}
	for _, v := range grid3 {
		for _, v2 := range v {
			part2 += int(v2)
		}
	}

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(day18("test.txt", 5))
	fmt.Println(day18("input.txt", 100))
}
