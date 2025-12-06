package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func day6(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	cells := [][]int{}
	cells2 := [][]int{}

	// ops := map[byte]bool{
	// 	'*':true, '+':true,
	// }
	ops := []string{}
	grid := [][]byte{}
	for scanner.Scan() {
		t := scanner.Text()
		gridrow := []byte{}
		for _, v := range t {
			gridrow = append(gridrow, byte(v))
		}
		grid = append(grid, gridrow)

		for {
			ot := t
			t = strings.ReplaceAll(t, "  ", " ")
			t = strings.TrimSpace(t)
			if t == ot {
				break
			}
		}

		if t[0] == '*' || t[0] == '+' {
			ops = strings.Split(t, " ")
			break
		}

		nums_ := strings.Split(t, " ")
		row := []int{}
		for _, v := range nums_ {
			n, _ := strconv.Atoi(v)
			row = append(row, n)
		}
		cells = append(cells, row)
	}

	row2 := []int{}
	opCount := len(ops) - 1
a:
	for y := len(grid[0]) - 1; y >= 0; y-- {
		s := ""
		for x := range len(grid) - 1 {
			if grid[x][y] != ' ' {
				s += string(grid[x][y])
			}
		}
		i, _ := strconv.Atoi(s)
		if i == 0 {
			if ops[opCount] == "*" {
				out := 1
				for _, j := range row2 {
					out *= j
				}
				// fmt.Println("part2", out)
				part2 += out
			} else {
				out := 0
				for _, j := range row2 {
					out += j
				}
				// fmt.Println("part2", out)
				part2 += out
			}
			row2 = []int{}
			opCount--
			if opCount < 0 {
				break a
			}
		} else {
			row2 = append(row2, i)
		}
	}
	if ops[opCount] == "*" {
		out := 1
		for _, j := range row2 {
			out *= j
		}
		// fmt.Println("part2", out)
		part2 += out
	} else {
		out := 0
		for _, j := range row2 {
			out += j
		}
		// fmt.Println("part2", out)
		part2 += out
	}
	_ = cells2

	// log.Printf("cells %+v\n", cells)
	// log.Printf("%#v\n", grid)

	for col, op := range ops {
		a := 0
		if op == "*" {
			a = 1
		}
		for i := range cells {
			// log.Println(i, col, cells[i][col])
			if op == "*" {
				a *= cells[i][col]
			} else {
				a += cells[i][col]
			}
		}
		// fmt.Println(a)
		part1 += a
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day6("test.txt")
	fmt.Println(part1, part2)
	if part1 != 4277556 || part2 != 3263827 {
		log.Fatal("Test failed ", part1, part2)
	}

	part1, part2 = day6("input.txt")
	fmt.Println(part1, part2)
	fmt.Println("Elapsed time:", time.Since(t1))
}
