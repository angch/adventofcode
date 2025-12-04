package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func day4(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	board := map[[2]int]bool{}

	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	_ = directions

	y := 0
	for scanner.Scan() {
		t := scanner.Text()
		_ = t

		for x, v := range t {
			if v == '@' {
				board[[2]int{x, y}] = true
			}
		}
		y++
	}

	iter := 0
	for {
		rmme := [][2]int{}
		for b := range board {
			count := 0
			for _, d := range directions {
				testcoord := [2]int{d[0] + b[0], d[1] + b[1]}
				if board[testcoord] {
					count++
				}
			}
			if count < 4 {
				rmme = append(rmme, b)
				if iter == 0 {
					part1++
				}
				part2++
			}
		}
		iter++

		for _, b := range rmme {
			delete(board, b)
		}
		if len(rmme) == 0 {
			break
		}
		// rmme = [][2]int{}
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day4("test.txt")
	fmt.Println(part1, part2)
	if part1 != 13 || part2 != 43 {
		log.Fatal("Test failed ", part1, part2)
	}

	part1, part2 = day4("input.txt")
	fmt.Println(part1, part2)
	fmt.Println("Elapsed time:", time.Since(t1))
}
