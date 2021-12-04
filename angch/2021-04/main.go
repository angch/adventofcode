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
	// log.Println(b, n)
	return board{nums: b, marks: marks}
}

func (b board) mark(n int) {
	for i := 0; i < len(b.nums); i++ {
		for j := 0; j < len(b.nums); j++ {
			// log.Println(i, j, len(b.nums), n)
			// nice, copilot
			if b.nums[i][j] == n {
				b.marks[i][j] = 1
				return
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

func (b board) isbingo() bool {
	for i := 0; i < len(b.nums); i++ {
		bcount := 0
		bcount2 := -0
		for j := 0; j < len(b.nums); j++ {
			if b.marks[i][j] == 1 {
				bcount++
			}
			if b.marks[j][i] == 1 {
				bcount2++
			}
		}
		if bcount == len(b.nums) {
			return true
		}
		if bcount2 == len(b.nums) {
			return true
		}
	}
	return false
}

func boardfromtext(lines []string) board {
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

func part1() {
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	nums := scanner.Text()
	boards := make([]board, 0)
	_ = nums
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
		boards = append(boards, boardfromtext(lines))
	}
	fmt.Println(boards)
	bingo := strings.Split(nums, ",")

	won := 0
	wonboard := make([]bool, len(boards))
	for _, v := range bingo {
		b := -1
		fmt.Sscanf(v, "%d", &b)
		if b == -1 {
			continue
		}
		// log.Println("mark", b)

		for bn, board := range boards {
			if wonboard[bn] {
				continue
			}
			board.mark(b)

			if board.isbingo() {
				won++
				wonboard[bn] = true
				fmt.Println("Bingo!", bn, won)
				fmt.Println("Sum:", board.sum())
				fmt.Println("Score:", board.sum()*b)
				if won == 1 {
					fmt.Println(board.marks)

				} else if won == len(boards)+1 {
					fmt.Println(board.marks)
					fmt.Println("Bingo2!", bn)
					fmt.Println("Sum:", board.sum())
					fmt.Println("Score:", board.sum()*b)
					return
				}
			}
		}
		// break
	}
}

func part2() {
	file, _ := os.Open(filepath)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		_ = line
	}
}

func main() {
	part1()
	// part2()
}
