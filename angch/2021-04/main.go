package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var filepath = "input.txt"

// var filepath = "test.txt"

var debug = false

type board struct {
	nums  [][]int
	marks [][]int
}

func newBoard(n int) board {
	b := make([][]int, n)
	marks := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
		marks[i] = make([]int, n)
	}
	return board{nums: b, marks: marks}
}

func boardFromLines(lines []string) board {
	n := len(lines)
	b := newBoard(n)
	for k, line := range lines {
		row := make([]int, 0, n)
		ns := strings.Split(line, " ")
		for _, n := range ns {
			if n != "" {
				num := 0
				fmt.Sscanf(n, "%d", &num)
				row = append(row, num)
			}

		}
		b.nums[k] = row
	}
	return b
}

func (b board) mark(n int) {
	for i := 0; i < len(b.nums); i++ {
		for j := 0; j < len(b.nums); j++ {
			if b.nums[i][j] == n {
				b.marks[i][j] = 1
			}
		}
	}
}

func (b board) sum() int {
	sum := 0
	for i := 0; i < len(b.nums); i++ {
		for j := 0; j < len(b.nums); j++ {
			if b.marks[i][j] == 0 {
				sum += b.nums[i][j]
			}
		}
	}
	return sum
}

func (b board) isBingo() bool {
	for i := 0; i < len(b.nums); i++ {
		bcount, bcount2 := 0, 0
		for j := 0; j < len(b.nums); j++ {
			if b.marks[i][j] == 1 {
				bcount++
			}
			if b.marks[j][i] == 1 {
				bcount2++
			}
		}
		if bcount == len(b.nums) || bcount2 == len(b.nums) {
			return true
		}
	}
	return false
}

func day4() {
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	nums := scanner.Text()
	boards := make([]board, 0)
	scanner.Scan()

	for {
		lines := make([]string, 0)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			lines = append(lines, line)
		}
		if len(lines) == 0 {
			break
		}
		boards = append(boards, boardFromLines(lines))
	}
	bingo := strings.Split(nums, ",")

	won := 0
	wonboard := make([]bool, len(boards))
	for _, v := range bingo {
		b := -1
		fmt.Sscanf(v, "%d", &b)
		if b == -1 {
			continue
		}

		for bn, board := range boards {
			if wonboard[bn] {
				continue
			}
			board.mark(b)

			if board.isBingo() {
				won++
				wonboard[bn] = true

				if won == 1 {
					// fmt.Println(board.marks)
					fmt.Println("Part 1 Score:", board.sum()*b)
				} else if won == len(boards) {
					// fmt.Println(board.marks)
					fmt.Println("Part 2 Score:", board.sum()*b)
					return
				}
			}
		}
	}
}

func main() {
	day4()
}
