package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"slices"
)

type Coord struct {
	X, Y int
}

func Manhattan(a, b Coord) int {
	return max(a.X, b.X) - min(a.X, b.X) + max(a.Y, b.Y) - min(a.Y, b.Y)
}

type Span struct {
	L, R int
}
type Spans []Span

type Board struct {
	Board  map[Coord]rune
	Spans  map[int]Spans
	Spans2 []Spans
	Min    Coord
	Max    Coord
}

func Compare(a, b Span) int {
	if a.L < b.L {
		return -1
	} else if a.L == b.L {
		if a.R < b.R {
			return -1
		} else if a.R == b.R {
			return 0
		}
		return 1
	}
	return 1
}

func Less(x, y Span) bool {
	return true
}

func NewSpans() Spans {
	s := make(Spans, 0, 10)
	return s
}

func (s *Spans) Add(l, r int) Spans {
	l, r = min(l, r), max(l, r)
	l, r = max(-1000, l), min(r, 4000000)
	return append(*s, Span{l, r})
}

func (s Spans) Compress() Spans {
	// Yes, we should have done this as we're adding spans, not after.
	// Deleting elements from the right to left means we avoid allocs.
	// fmt.Println(" pp ", s)

	for i := len(s) - 1; i > 0; i-- {
		j := i - 1
		a, b := s[j], s[i]
		// fmt.Println("a b", a, b)
		if a.R >= b.L-1 {
			a.R = max(b.R, a.R)
			// fmt.Println("new a", a)
			s[j] = a

			// Delete an element:
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			// s = append(s[:i], s[i+1:]...)

			// if a.R >= b.L {
			// 	fmt.Println("repeat")
			// 	i++
			// }
			// Delete an element, badly:
			//s = append(s[:i], s[i+1:]...)
			// fmt.Println(" pp ", s)
		}
	}
	return s
}

func NewBoard(max int) *Board {
	return &Board{
		make(map[Coord]rune),
		make(map[int]Spans),
		make([]Spans, max),
		Coord{0, 0},
		Coord{0, 0},
	}
}

var maxRows = 4000000

func (b *Board) Clone(max int) *Board {
	b2 := Board{
		make(map[Coord]rune),
		make(map[int]Spans),
		make([]Spans, max),
		b.Min, b.Max,
	}
	for k, v := range b.Board {
		b2.Board[k] = v
	}
	for k, v := range b.Spans {
		b2.Spans[k] = v
	}
	return &b2
}

func (b *Board) Draw() {
	for y := b.Min.Y; y <= b.Max.Y; y++ {
		fmt.Printf("%3d ", y)
		for x := b.Min.X; x <= b.Max.X; x++ {
			i, ok := b.Board[Coord{x, y}]
			if ok {
				fmt.Print(string(i))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (b *Board) Set(c Coord, v rune) {
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

func day15(file string, countRow int) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	board := NewBoard(maxRows + 1)
	sensors := []Coord{}

	for scanner.Scan() {
		t := scanner.Text()
		var sensor, beacon Coord
		fmt.Sscanf(t, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)

		// fmt.Println(sensor, beacon)
		board.Set(sensor, 'S')
		sensors = append(sensors, sensor)
		distance := Manhattan(sensor, beacon)
		// fmt.Println(distance)

		top := sensor.Y - distance
		bottom := sensor.Y + distance
		i := 0
		for y := top; y <= bottom; y++ {
			if y == countRow {
				for x := sensor.X - i; x <= sensor.X+i; x++ {
					_, ok := board.Board[Coord{x, y}]
					if !ok {
						board.Set(Coord{x, y}, '#')
					}
				}
			}

			if y >= 0 && y <= 4000000 {
				s, ok := board.Spans[y]
				if !ok {
					s = NewSpans()
					board.Spans[y] = s
				}
				// board.Spans[y] = append(s, [2]int{sensor.X - i, sensor.X + i})
				// if y == 10 {
				// 	fmt.Println("debug", sensor.X-i, sensor.X+i)
				// }
				board.Spans[y] = s.Add(sensor.X-i, sensor.X+i)
			}

			if y < sensor.Y {
				i++
			} else {
				i--
			}
		}
		board.Set(beacon, 'B')
	}
	// board.Draw()

	// part1
	if false {
		for x := board.Min.X; x <= board.Max.X; x++ {
			c := board.Board[Coord{x, countRow}]
			if c == '#' {
				part1++
			}
		}
	}

	fmt.Println("Span counting")
	spans := board.Spans[countRow]
	counts := make(map[int]bool)
	for i := 0; i < len(spans); i++ {
		for j := spans[i].L; j <= spans[i].R; j++ {
			counts[j] = true
		}
	}
	for _, b := range sensors {
		if b.Y == countRow {
			delete(counts, b.X)
		}
	}
	// board.Draw()
	part1 = len(counts) - 1
	log.Println("xxx", len(counts), len(board.Spans))

	// part2
	minY := max(board.Min.Y, 0)
	maxY := min(board.Max.Y, 4000000)
	for y := minY; y <= maxY; y++ {
		span, ok := board.Spans[y]
		if !ok {
			continue
		}
		slices.SortFunc[Spans](span, Compare)
		span = span.Compress()
		if len(span) == 1 && span[0].L == -1000 && span[0].R == 4000000 {
			continue
		}
		if len(span) == 2 {
			for i := 0; i < len(span)-1; i++ {
				if span[i+1].L-span[i].R == 2 {
					fmt.Printf("af %d %v\n", y, span)
					fmt.Println("solution is", y, span[i].R+1, (span[i].R+1)*4000000+y)
					part2 = (span[i].R+1)*4000000 + y
				}
			}
		}
	}

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// part1, part2 := day15("test.txt", 10)
	part1, part2 := day15alt3("test.txt", 10)

	fmt.Println(part1, part2)
	if part1 != 26 || part2 != 56000011 {
		log.Fatal("Bad test")
	}
	// fmt.Println(day15("input.txt", 2000000))     // 11062575042518 too low
	fmt.Println(day15alt("input.txt", 2000000)) // 11062575042518 too low
}
