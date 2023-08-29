package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func day15alt3(file string, countRow int) (int, int) {
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
				s2 := s.Add(sensor.X-i, sensor.X+i)
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
	slices.SortFunc[Spans](spans, Compare)
	spans = spans.Compress()
	board.Spans2[countRow] = spans
	counts := 0
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
		spans := board.Spans2[y]
		if spans == nil {
			continue
		}
		// log.Println("b4", spans)
		slices.SortFunc[Spans](spans, Compare)
		// log.Println("af", spans)
		spans = spans.Compress()
		// log.Println("cp", spans)
		if len(spans) == 2 {
			if spans[1].L-spans[0].R == 2 {
				part2 = (spans[0].R+1)*4000000 + y
				break p2
			}
		}
		// fmt.Printf("af %d %v\n", y, span)
	}

	return part1, part2
}
