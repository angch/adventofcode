package main

import (
	"fmt"
	"strconv"
	"strings"
)

func do(inputs string) (int, []int) {
	rope := make([]int, 256)
	for i := range rope {
		rope[i] = i
	}
	row := make([]int, 128)

	cur, skip := 0, 0

	lengths := make([]int, len(inputs))
	for i, c := range inputs {
		lengths[i] = int(c)
	}
	for _, c := range []int{17, 31, 73, 47, 23} {
		lengths = append(lengths, c)
		//fmt.Println(rope[0] * rope[1])
	}

	if false {
		// Part1
		lengths1 := strings.Split(inputs, ",")
		lengths = make([]int, len(lengths1))
		for i, s := range lengths1 {
			lengths[i], _ = strconv.Atoi(s)
		}
		rope, cur, skip = knot(rope, lengths, cur, skip, 1)
		fmt.Println(rope[0] * rope[1])
		return 0, row
	}

	rope, cur, skip = knot(rope, lengths, cur, skip, 64)

	hash := make([]int, 16)
	for i, r := 0, 0; i < len(rope); r++ {
		c := 0
		for r := 0; r < 16; r++ {
			c ^= rope[i]
			i++
		}
		hash[r] = c
	}

	hexOut := ""
	bits := make([]int, 0)
	for hex := 0; hex < 16; hex++ {
		//fmt.Printf("%02x", hash[hex])
		hexOut += fmt.Sprintf("%02x", hash[hex])
		bits = append(bits, (hash[hex]>>4)&15)
		bits = append(bits, hash[hex]&15)
	}
	//fmt.Println(hexOut, bits)

	count := 0

	x := 0
	show := false
	for _, h := range bits {
		//fmt.Println(h)

		for i := 0; i < 4; i++ {
			if h&8 != 0 {
				if show {
					fmt.Print("#")
				}
				row[x] = -1
				count++
			} else {
				if show {
					fmt.Print(".")
				}
			}
			h <<= 1
			x++
		}
		//fmt.Println()
	}
	if show {
		fmt.Println()
	}
	return count, row
}

func main() {
	inputs := "157,222,1,2,177,254,0,228,159,140,249,187,255,51,76,30"
	//inputs := "3,4,1,5"
	//inputs = "1,2,3"
	//inputs = "1,2,4"
	//inputs = ""
	//inputs = "AoC 2017"
	inputs = "flqrgnkx"
	inputs = "hfdlxzhv"
	c := 0
	board := make([][]int, 128)
	for i := 0; i < 128; i++ {
		inp := fmt.Sprintf("%s-%d", inputs, i)
		count, row := do(inp)
		board[i] = row
		c += count
	}
	//fmt.Println(c)
	//fmt.Println(board)

	color := 1
a:
	for {
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				if board[y][x] == -1 {
					colorB(board, x, y, color)
					if false {
						for y := 0; y < len(board); y++ {
							fmt.Println(board[y])
						}
					}
					color++
					continue a
				}
			}
		}
		break
	}

	if false {
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				fmt.Printf("%02d ", board[y][x])
			}
			fmt.Println()
		}
	}
	fmt.Println(c, color)
}

type Coord struct{ x, y int }

var deltas = []Coord{
	Coord{-1, 0},
	Coord{1, 0},
	Coord{0, -1},
	Coord{0, 1},
}

func colorB(board [][]int, x, y, c int) {
	board[y][x] = c
	for _, d := range deltas {
		xc, yc := x+d.x, y+d.y
		if yc >= 0 && yc < len(board) {
			if xc >= 0 && xc < len(board[y]) {
				if board[yc][xc] == -1 {
					colorB(board, xc, yc, c)
				}
			}
		}
	}
}

func knot(rope []int, lengths []int, cur, skip int, rounds int) ([]int, int, int) {
	for r := 0; r < rounds; r++ {
		for _, length := range lengths {
			j := length - 1
			for i := 0; i <= j; i++ {
				rope[(cur+i)%len(rope)], rope[(cur+j)%len(rope)] = rope[(cur+j)%len(rope)], rope[(cur+i)%len(rope)]
				j--
			}
			cur = (cur + length + skip) % len(rope)
			skip++
		}
	}
	return rope, cur, skip
}
