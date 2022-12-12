package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord [2]int

var directions []coord = []coord{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func solve(grid map[coord]int, start, end coord) int {
	dist := map[coord]int{}
	queue := []coord{start}

	for len(queue) > 0 {
		this := queue[0]
		queue = queue[1:]
		h := grid[this]
		if h == 'S' {
			h = 'a'
		}
		// visted[this] = true
		for _, d := range directions {
			neighbour := coord{this[0] + d[0], this[1] + d[1]}
			h2, ok := grid[neighbour]
			if !ok {
				continue
			}
			if h2 <= h+1 {
				x, ok2 := dist[neighbour]
				_ = x
				if !ok2 {
					queue = append(queue, neighbour)
					dist[neighbour] = dist[this] + 1
				}
			}

			if grid[neighbour] == 'E' && (h == 'z' || h == 'y') {
				return dist[this] + 1
			}
		}
	}
	return 999999
}

func day12(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	grid := map[coord]int{}

	scanner := bufio.NewScanner(f)
	y := 0
	start := coord{0, 0}
	end := coord{0, 0}
	// maxx := 0
	a := []coord{}
	b := []coord{}
	for scanner.Scan() {
		t := scanner.Text()
		for x, c := range t {
			if c == 'S' {
				start = coord{x, y}
			} else if c == 'E' {
				end = coord{x, y}
			}
			grid[coord{x, y}] = int(c)
			if c == 'a' {
				a = append(a, coord{x, y})
			} else if c == 'z' {
				b = append(b, coord{x, y})
			}
		}
		// maxx = len(t)
		y++
	}
	// fmt.Println(grid, start, end)
	fmt.Println("grid is ", y)
	// maxy := y

	part1 = solve(grid, start, end)

	// for _, c := range b {
	// 	grid[c] = 'E'
	// }
	grid[start] = 'a'
	part2 = part1
	for _, c := range a {
		grid[c] = 'S'
		p2 := solve(grid, c, end)
		if p2 < part2 {
			part2 = p2
		}
		grid[c] = 'a'
	}

	// part1 = dist[end]
	// if false {
	// 	for y := 0; y < maxy; y++ {
	// 		for x := 0; x < maxx; x++ {
	// 			if visted[coord{x, y}] {
	// 				// fmt.Print("X")
	// 				fmt.Printf("%03x", dist[coord{x, y}])
	// 			} else {
	// 				// fmt.Print(".")
	// 				fmt.Printf("%03x", dist[coord{x, y}])
	// 			}

	// 		}
	// 		fmt.Println()
	// 	}
	// }
	// fmt.Println("visted", len(visted), len(grid))
	// _ = best
	return part1, part2

}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// fmt.Println(day21("test.txt", person{HP: 10, Mana: 250}, []string{"Poison", "Magic Missile"}))
	// fmt.Println(day21("test2.txt", person{HP: 10, Mana: 250}, []string{"Recharge", "Shield", "Drain", "Poison", "Magic Missile"}))
	// fmt.Println(day12("test.txt"))
	part1, part2 := day12("test.txt")
	fmt.Println(part1, part2)
	if part1 != 31 {
		log.Fatal("Bad")
		_ = part2
	}
	// 1309 too high
	fmt.Println(day12("input.txt")) // 352 too high
}
