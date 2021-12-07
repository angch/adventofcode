package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	nums   map[int][2]int
	sum    int
	count1 []int
	count2 []int
	bingo  bool
}

func boardFromLines(lines []string) *board {
	n := len(lines)
	b := board{
		nums:   make(map[int][2]int),
		count1: make([]int, n),
		count2: make([]int, n),
	}
	sum := 0

	for i, line := range lines {
		ns := strings.Split(line, " ")
		j := 0
		for _, n := range ns {
			if n != "" {
				num := 0
				fmt.Sscanf(n, "%d", &num)
				b.nums[num] = [2]int{i, j}
				j++
				sum += num
			}
		}
		b.count1[i] = n
		b.count2[i] = n
	}
	b.sum = sum
	return &b
}

func (b *board) mark(n int) {
	coord, ok := b.nums[n]
	if ok {
		b.count1[coord[0]]--
		b.count2[coord[1]]--
		b.sum -= n
		if b.count1[coord[0]] == 0 || b.count2[coord[1]] == 0 {
			b.bingo = true
			return
		}
	}
}

func day4(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	nums := scanner.Text()
	boards := make([]*board, 0)
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
	drawnNums := strings.Split(nums, ",")

	won := 0
	for _, v := range drawnNums {
		b, _ := strconv.Atoi(v)

		for _, board := range boards {
			if board.bingo {
				continue
			}
			board.mark(b)

			if board.bingo {
				won++
				if won == 1 {
					fmt.Println("Part 1", board.sum*b)
				} else if won == len(boards) {
					fmt.Println("Part 2", board.sum*b)
					return
				}
			}
		}
	}
}

func main() {
	day4("test.txt")
	day4("input.txt")
}
