package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	X, Y int
}

type Board struct {
	Board map[Coord]int
	Min   Coord
	Max   Coord
}

func (b *Board) Clone() *Board {
	b2 := Board{make(map[Coord]int), b.Min, b.Max}
	for k, v := range b.Board {
		b2.Board[k] = v
	}
	return &b2
}

func (b *Board) Draw() {
	for y := b.Min.Y; y <= b.Max.Y; y++ {
		for x := b.Min.X; x <= b.Max.X; x++ {
			i, ok := b.Board[Coord{x, y}]
			if ok {
				switch i {
				case 1:
					fmt.Print("#")
				case 2:
					fmt.Print("+")
				case 3:
					fmt.Print("O")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (b *Board) Set(c Coord, v int) {
	if len(b.Board) > 0 {
		b.Min.X = min(b.Min.X, c.X)
		b.Min.Y = min(b.Min.Y, c.Y)
		b.Max.X = max(b.Max.X, c.X)
		b.Max.Y = max(b.Max.Y, c.Y)
	} else {
		b.Min = c
		b.Max = c
	}
	b.Board[c] = v
}

func day14(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	board := Board{make(map[Coord]int), Coord{0, 0}, Coord{0, 0}}

	for scanner.Scan() {
		t := scanner.Text()

		coordspairsStr := strings.Split(t, " -> ")
		// log.Println(coordspairsStr)

		var prev Coord
		for k, v := range coordspairsStr {
			c := Coord{0, 0}
			fmt.Sscanf(v, "%d,%d", &c.X, &c.Y)
			if k > 0 {
				d := Coord{0, 0}
				if prev.X < c.X {
					d.X = 1
				} else if prev.X > c.X {
					d.X = -1
				}
				if prev.Y < c.Y {
					d.Y = 1
				} else if prev.Y > c.Y {
					d.Y = -1
				}
				board.Set(c, 1)
				for prev != c {
					prev.X += d.X
					prev.Y += d.Y
					board.Set(prev, 1)
				}
			} else {
				prev = c
				board.Set(c, 1)
			}
		}

		_ = t
	}
	entry := Coord{500, 0}
	board.Set(entry, 2)
	// board.Draw()
	board2 := board.Clone()

a:
	// Keep dropping sand
	for {
		sand := entry
		// Drop a sand.
	b:
		// Falling sand
		for {
			sand.Y++
			if sand.Y > board.Max.Y {
				// until a sand escapes
				break a
			}

			_, ok := board.Board[sand]
			if !ok {
				continue
			}
			sand.X--
			_, ok = board.Board[sand]
			if !ok {
				continue
			}
			sand.X += 2
			_, ok = board.Board[sand]
			if !ok {
				continue
			}

			sand.Y--
			sand.X--
			break b
		}
		board.Set(sand, 3)
		part1++

		// fmt.Println("Dropped sand", sand)
		// board.Draw()
	}

	// Keep dropping sand
	maxY := board2.Max.Y + 1
	for {
		sand := entry
		// Drop a sand.
	b2:
		// Falling sand
		for {
			sand.Y++
			if sand.Y > maxY {
				// until a sand escapes
				sand.Y--
				break b2
			}

			_, ok := board2.Board[sand]
			if !ok {
				continue
			}
			sand.X--
			_, ok = board2.Board[sand]
			if !ok {
				continue
			}
			sand.X += 2
			_, ok = board2.Board[sand]
			if !ok {
				continue
			}

			sand.Y--
			sand.X--
			break b2
		}
		if sand == entry {
			break
		}
		board2.Set(sand, 3)
		part2++

		// fmt.Println("Dropped sand", sand)
		// board2.Draw()
	}
	part2++

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day14("test.txt")

	fmt.Println(part1, part2)
	if part1 != 24 || part2 != 93 {
		log.Fatal("Bad test")
	}
	fmt.Println(day14("input.txt"))
}
