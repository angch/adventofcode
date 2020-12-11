package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func dump(board [][]byte) {
	for _, v := range board {
		fmt.Println(string(v))
	}
}

func occupied(board [][]byte, x, y int) int {
	count := 0
	for yy := y - 1; yy <= y+1; yy++ {
		if yy < 0 {
			continue
		}
		if yy >= len(board) {
			continue
		}
		for xx := x - 1; xx <= x+1; xx++ {
			if xx < 0 {
				continue
			}
			if xx >= len(board[0]) {
				continue
			}
			if xx == x && yy == y {
				continue
			}
			if board[yy][xx] == '#' {
				count++
			}
		}
	}
	return count
}

func occupied2(board [][]byte, x, y int) int {
	count := 0
	dx, dy := 0, 1

	for rep := 0; rep < 4; rep++ {
		x1, y1 := x, y
		for {
			x1 += dx
			y1 += dy
			if x1 >= len(board[0]) {
				break
			}
			if x1 < 0 {
				break
			}
			if y1 >= len(board) {
				break
			}
			if y1 < 0 {
				break
			}
			if board[y1][x1] == '.' {
				continue
			}
			if board[y1][x1] == 'L' {
				break
			}
			if board[y1][x1] == '#' {
				count++
				break
			}
		}

		dx, dy = -dy, dx
	}
	dx, dy = 1, 1

	for rep := 0; rep < 4; rep++ {
		x1, y1 := x, y
		for {
			x1 += dx
			y1 += dy
			if x1 >= len(board[0]) {
				break
			}
			if x1 < 0 {
				break
			}
			if y1 >= len(board) {
				break
			}
			if y1 < 0 {
				break
			}
			if board[y1][x1] == '.' {
				continue
			}
			if board[y1][x1] == 'L' {
				break
			}
			if board[y1][x1] == '#' {
				count++
				break
			}
		}

		dx, dy = -dy, dx
	}

	return count
}

func deepcopy(b1, b2 [][]byte) {
	for k, v := range b2 {
		b1[k] = make([]byte, len(v))
		copy(b1[k], v)
	}
}

func do1(board [][]byte) int {
	for rep := 0; rep < 66666; rep++ {
		// log.Println("rep", rep)
		// dump(board)
		board2 := make([][]byte, len(board))
		deepcopy(board2, board)
		for y, b := range board {
			for x, c := range b {
				count := occupied(board, x, y)
				if c == 'L' {
					if count == 0 {
						c = '#'
					}
				} else if c == '#' {
					if count >= 4 {
						c = 'L'
					}
				}
				board2[y][x] = c
			}
		}
		if reflect.DeepEqual(board, board2) {
			log.Println("Stable at", rep)
			break
		}
		board = board2
		// fmt.Println()
	}

	count := 0
	for _, v := range board {
		for _, v2 := range v {
			if v2 == '#' {
				count++
			}
		}
	}
	return count
}

func do2(board [][]byte) int {
	for rep := 0; rep < 66666; rep++ {
		// log.Println("rep", rep)
		// dump(board)
		board2 := make([][]byte, len(board))
		deepcopy(board2, board)
		for y, b := range board {
			for x, c := range b {
				count := occupied2(board, x, y)
				if c == 'L' {
					if count == 0 {
						c = '#'
					}
				} else if c == '#' {
					if count >= 5 {
						c = 'L'
					}
				}
				board2[y][x] = c
			}
		}
		if reflect.DeepEqual(board, board2) {
			log.Println("Stable at", rep)
			break
		}
		board = board2
		// fmt.Println()
	}

	count := 0
	for _, v := range board {
		for _, v2 := range v {
			if v2 == '#' {
				count++
			}
		}
	}
	return count
}

func do(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	board := make([][]byte, 0)
	for scanner.Scan() {
		l := scanner.Text()
		board = append(board, []byte(l))
	}

	count := do1(board)
	count2 := do2(board)
	return count, count2
}

func main() {
	log.Println(do("test.txt"))
	// log.Println(do("test2.txt"))
	// log.Println(do("test3.txt"))
	log.Println(do("input.txt"))
}
