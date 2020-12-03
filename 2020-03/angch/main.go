package main

import (
	"bufio"
	"log"
	"os"
)

func findtree(dx, dy int, inputs []string) int {
	tree := 0
	x, y := 0, 0
	for y < len(inputs) {
		x, y = x+dx, y+dy
		if y >= len(inputs) {
			break
		}
		x = x % len(inputs[0])
		if inputs[y][x] == '#' {
			tree++
		}
	}
	return tree
}

func main() {
	inputs2 := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	inputs := make([]string, 0)
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		inputs = append(inputs, l)
	}

	// x, y := 0, 0
	dx, dy := 3, 1

	log.Println(findtree(dx, dy, inputs))

	slopes := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}
	m := 1
	for _, v := range slopes {
		m *= findtree(v[0], v[1], inputs)
	}
	if false {
		log.Println(inputs2)
	}
	log.Println(m)
}
