package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func day15alt2(file string, countRow int) (int, int) {
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

		wg := sync.WaitGroup{}
		for y := top; y <= bottom; y++ {
			if y >= 0 && y <= 4000000 {
				wg.Add(1)
				s := board.Spans2[y]
				if s == nil {
					s = NewSpans()
				}
				go func(s Spans, i, y, x1, x2 int) {
					board.Spans2[y] = s.AddCompress(x1, x2)
					wg.Done()
				}(s, i, y, sensor.X-i, sensor.X+i)
			}
			if y < sensor.Y {
				i++
			} else {
				i--
			}
		}
		wg.Wait()
		board.Set(beacon, 'B')
	}
	// board.Draw()

	// fmt.Println("Span counting")
	spans := board.Spans2[countRow]
	counts := 0
	for i := 0; i < len(spans); i++ {
		counts += spans[i].R - spans[i].L + 1
	}
	for _, b := range sensors {
		if b.Y == countRow {
			counts--
		}
	}

	// board.Draw()
	part1 = counts - 1
	// log.Println("xxx", len(counts), len(board.Spans))

	// part2
	minY := max(board.Min.Y, 0)
	maxY := min(board.Max.Y, 4000000)
p2:
	for y := minY; y <= maxY; y++ {
		span := board.Spans2[y]
		if span == nil {
			continue
		}
		if len(span) == 1 && span[0].L == -1000 && span[0].R == 4000000 {
			continue
		}
		if len(span) == 2 {
			for i := 0; i < len(span)-1; i++ {
				if span[i+1].L-span[i].R == 2 {
					// fmt.Printf("af %d %v\n", y, span)
					part2 = (span[i].R+1)*4000000 + y
					// fmt.Println("solution is", y, span[i].R+1, part2)
					break p2
				}
			}
		}
		// fmt.Printf("af %d %v\n", y, span)
	}

	return part1, part2
}
