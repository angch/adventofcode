package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	grid := make([][]bool, 1000)
	for g := range grid {
		grid[g] = make([]bool, 1000)
	}
	grid2 := make([][]int, 1000)
	for g := range grid2 {
		grid2[g] = make([]int, 1000)
	}

	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		var x1, y1, x2, y2 int
		instr := ""
		if line[0:7] == "turn on" {
			null := ""
			instr = "on"
			fmt.Sscanf(line, "%s %s %d,%d %s %d,%d", &null, &null, &x1, &y1, &null, &x2, &y2)
		}
		if line[0:6] == "toggle" {
			null := ""
			instr = "xor"
			fmt.Sscanf(line, "%s %d,%d %s %d,%d", &null, &x1, &y1, &null, &x2, &y2)
			//fmt.Println(x1, x2, y1, y2)
		}
		if line[0:8] == "turn off" {
			null := ""
			instr = "off"
			fmt.Sscanf(line, "%s %s %d,%d %s %d,%d", &null, &null, &x1, &y1, &null, &x2, &y2)
			//fmt.Println(x1, x2, y1, y2)
		}

		if x2 < x1 {
			x2, x1 = x1, y2
		}
		if y2 < y1 {
			y2, y1 = y1, y2
		}
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				switch instr {
				case "on":
					grid[x][y] = true
					grid2[x][y]++
				case "off":
					grid[x][y] = false
					grid2[x][y]--
					if grid2[x][y] <= 0 {
						grid2[x][y] = 0
					}
				case "xor":
					grid[x][y] = !grid[x][y]
					grid2[x][y] += 2
				}
			}
		}
	}
	count, count2 := 0, 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] {
				count++
			}
			count2 += grid2[x][y]
		}
	}
	fmt.Println(count, count2)
}
