package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func (s Spans) Compress3(start int) Spans {
	// Yes, we should have done this as we're adding spans, not after.
	// Deleting elements from the right to left means we avoid allocs.
	for i := start; i > 0; i-- {
		j := i - 1
		a, b := s[j], s[i]
		if a.R >= b.L-1 {
			a.R = max(b.R, a.R)

			c := 1
			i2 := i + 1
			for ; i2 < len(s); i2, c = i2+1, c+1 {
				if a.R < s[i2].R {
					if a.R >= s[i2].L {
						a.R = s[i2].R
						continue
					}
					break
				}
			}
			s[j] = a

			copy(s[i:], s[i2:])
			s = s[:len(s)-c]
		}
	}
	return s
}

func (s Spans) AddCompress(l, r int) Spans {
	l, r = min(l, r), max(l, r)
	// l, r = max(-10, l), min(r, 4000000)

	if len(s) == 0 {
		return append(s, Span{l, r})
	}

	var start int
	// Insertion sort
	for i := 0; i < len(s); i++ {
		if s[i].L > l {
			// Check for special case where we are within the previous span
			if i > 0 && s[i-1].R >= r {
				// It is within the previous span, let's just go home
				return s
			}
			if i > 0 && s[i-1].R >= l && r < s[i].L {
				// New span only extends to left span, not touching the next span's left
				s[i-1].R = r
				return s
			}
			// if s[i].L <= r && r <= s[i].R {
			// 	// New span only extends to left span
			// 	s[i].L = l
			// 	start = i // FIXME: We might have an off by one error
			// 	goto compress
			// 	// return s
			// }

			s = append(s, Span{})
			copy(s[i+1:], s[i:])
			s[i] = Span{l, r}

			// Slower, as it forces an alloc
			// The above avoids an alloc
			// *s = append((*s)[:i], append([]Span{{l, r}}, (*s)[i:]...)...)
			start = i + 1 // FIXME: We might have an off by one error
			goto compress
			// return s.Compress3(i)
		}
	}

	start = len(s)
	s = append(s, Span{l, r})
	// return s.Compress()

compress:
	for i := start; i > 0; i-- {
		j := i - 1
		a, b := s[j], s[i]
		if a.R >= b.L-1 {
			a.R = max(b.R, a.R)

			c := 1
			i2 := i + 1
			for ; i2 < len(s); i2, c = i2+1, c+1 {
				if a.R < s[i2].R {
					if a.R >= s[i2].L {
						a.R = s[i2].R
						continue
					}
					break
				}
			}
			s[j] = a

			copy(s[i:], s[i2:])
			s = s[:len(s)-c]
		}
	}
	return s
}

func day15alt(file string, countRow int) (int, int) {
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
			if y >= 0 && y <= 4000000 {
				s := board.Spans2[y]
				if s == nil {
					s = NewSpans()
				}
				s2 := s.AddCompress(sensor.X-i, sensor.X+i)
				board.Spans2[y] = s2
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

	// fmt.Println("Span counting")
	spans := board.Spans2[countRow]
	counts := 0
	// fmt.Println("spans are", spans)
	for i := 0; i < len(spans); i++ {
		counts += spans[i].R - spans[i].L + 1
	}
	for _, b := range sensors {
		if b.Y == countRow {
			counts--
		}
	}
	part1 = counts - 1
	// log.Println("xxx", len(counts), len(board.Spans))

	// part2
	minY := max(board.Min.Y, 0)
	maxY := min(board.Max.Y, 4000000)
p2:
	for y := minY; y <= maxY; y++ {
		span := board.Spans2[y]
		if len(span) == 2 {
			if span[1].L-span[0].R == 2 {
				part2 = (span[0].R+1)*4000000 + y
				break p2
			}
		}
		// fmt.Printf("af %d %v\n", y, span)
	}

	return part1, part2
}
