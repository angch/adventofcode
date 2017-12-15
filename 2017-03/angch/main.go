package main

import (
	"log"
)

type Coord struct {
	x, y  int
	value int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func calcValue(spiral []Coord, now Coord) int {
	if len(spiral) == 0 {
		return 1
	}

	v := 0
	for _, c := range spiral {
		if abs(c.x-now.x) < 2 && abs(c.y-now.y) < 2 {
			v += c.value
		}
	}
	return v
}

func main() {
	spiral := make([]Coord, 0)

	var now, min, max, dir Coord
	dir.x++
	partTwo := true

	for i := 1; ; i++ {
		if partTwo {
			now.value = calcValue(spiral, now)
		}
		spiral = append(spiral, now)

		if partTwo && now.value > 312051 {
			log.Println(i, now.value)
			break
		} else if !partTwo && i == 312051 {
			log.Println(i, abs(now.x)+abs(now.y))
			break
		}

		now.x += dir.x
		now.y += dir.y

		if now.x > max.x {
			max.x = now.x
			dir = Coord{dir.y, -dir.x, 0}
		}
		if now.x < min.x {
			min.x = now.x
			dir = Coord{dir.y, -dir.x, 0}
		}
		if now.y > max.y {
			max.y = now.y
			dir = Coord{dir.y, -dir.x, 0}
		}
		if now.y < min.y {
			min.y = now.y
			dir = Coord{dir.y, -dir.x, 0}
		}
	}

}
