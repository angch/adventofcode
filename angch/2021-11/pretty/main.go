package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Coord struct{ X, Y int }

var adj = []Coord{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func day11(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	board := make(map[Coord]int)
	c := Coord{0, 0}
	for scanner.Scan() {
		t := scanner.Text()
		for _, v := range t {
			board[c] = int(v - '0')
			c.X++
		}
		c = Coord{0, c.Y + 1}
	}

	part1, part2 := 0, 0
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGray)
	boomStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGray)
	levelStyle := make([]tcell.Style, 10)
	for i := 0; i < 10; i++ {
		style := tcell.StyleDefault.Foreground(tcell.ColorGray)
		g := int32(i * 128 / 10)
		style = style.Background(tcell.NewRGBColor(g, g, g))
		levelStyle[i] = style
	}
	t := fmt.Sprint("After step", 11, part1)
	for k, v := range t {
		s.SetContent(k+1, 1, v, nil, tcell.StyleDefault)
	}
	s.SetStyle(defStyle)
	s.Clear()
	old := make(map[Coord]int)
	draw := func(step int) {
		// return
		t := fmt.Sprintln("After step", step, part1)
		for k, v := range t {
			s.SetContent(k, 0, v, nil, tcell.StyleDefault)
		}
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				c := Coord{x, y}
				v := board[c]
				if v >= 10 {
					v = 0
				}
				if old[c] != v {
					if v == 0 {
						s.SetContent(x, y+1, rune(v+'0'), nil, boomStyle)

					} else {
						s.SetContent(x, y+1, rune(v+'0'), nil, levelStyle[v])
					}
					old[c] = v
				}
			}
			// fmt.Println()
		}
		s.Show()

		time.Sleep(time.Millisecond * 10)
	}
	_ = draw

	for step := 1; ; step++ {
		flashed := make(map[Coord]bool)
		eval := make([]Coord, 0)

		for c := range board {
			board[c]++
			if board[c] > 9 {
				eval = append(eval, c)
			}
		}
		delay := []int{len(eval)}
		drawn := 0

		for len(eval) > 0 {
			c, eval = eval[len(eval)-1], eval[:len(eval)-1]
			if flashed[c] {
				continue
			}
			flashed[c] = true
			board[c] = 0
			if step <= 100 {
				part1++
			}
			d2 := 0
			for _, d := range adj {
				c2 := Coord{c.X + d.X, c.Y + d.Y}

				if _, ok := board[c2]; !ok {
					continue
				}
				if !flashed[c2] {
					board[c2]++
					if board[c2] > 9 {
						eval = append(eval, c2)
						d2++
					}
				}
			}
			if d2 > 0 {
				delay = append(delay, d2)
			}
			delay[0]--
			if delay[0] <= 0 {
				draw(step)
				drawn++
				delay = delay[1:]
			}
		}

		draw(step)
		drawn++
		delay2 := 200 - drawn*10
		if delay2 > 0 {
			time.Sleep(time.Duration(delay2) * time.Millisecond)
		}
		if len(flashed) == len(board) {
			part2 = step
			break
		}
	}

	t = fmt.Sprint("Part 1", part1)
	for k, v := range t {
		s.SetContent(k, 12, v, nil, tcell.StyleDefault)
	}
	t = fmt.Sprint("Part 2", part2)
	for k, v := range t {
		s.SetContent(k, 13, v, nil, tcell.StyleDefault)
	}
	s.Fini()
}

func main() {
	// day11("test.txt")
	day11("input.txt")
}
