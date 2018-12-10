package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Particle struct {
	X, Y, DX, DY int
	done         bool
}

func main() {
	fileName := "input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	particles := make([]Particle, 0)
	scanner := bufio.NewScanner(file)
	maxX, maxY, minX, minY := 0, 0, 0, 0
	for scanner.Scan() {
		t := scanner.Text()

		var x, y, dx, dy int
		fmt.Sscanf(t, "position=<%d, %d> velocity=<%d,%d>", &x, &y, &dx, &dy)
		//log.Println(x, y, dx, dy)
		particles = append(particles, Particle{x, y, dx, dy, false})

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
	}

	w, h := maxX-minX, maxY-minY
	log.Println(w, h, maxX, maxY, minX, minY)
	t := 0
	board := make([][]byte, h+1)
	for {
		fmt.Println("t =", t)

		done := 0
		minX1, maxX1, minY1, maxY1 := 999999, -99999, 999999, -99999
		for k, p := range particles {
			if p.done {
				continue
			}
			done++

			particles[k].X += particles[k].DX
			particles[k].Y += particles[k].DY
			if particles[k].X < minX ||
				particles[k].Y < minY ||
				particles[k].X > maxX ||
				particles[k].Y > maxY {
				particles[k].done = true
			}
			p2 := &particles[k]
			if p2.X < minX1 {
				minX1 = p2.X
			}
			if p2.X > maxX1 {
				maxX1 = p2.X
			}
			if p2.Y < minY1 {
				minY1 = p2.Y
			}
			if p2.Y > maxY1 {
				maxY1 = p2.Y
			}
		}

		//if (maxX1-minX1)*(maxY1-minY1) < len(particles)*len(particles)/10 {
		if (maxX1-minX1)*(maxY1-minY1) < 1050 {
			log.Println("t is", t+1)
			h1 := maxY1 - minY1
			w1 := maxX1 - minX1
			//log.Println(h1, w1, maxX1, minX1, maxY1, minY1)
			//log.Fatal("adsf")
			board = make([][]byte, h1+1)
			for _, p := range particles {
				if len(board[p.Y-minY1]) == 0 {
					board[p.Y-minY1] = make([]byte, w1+1)
				}

				board[p.Y-minY1][p.X-minX1] = '#'
			}

			for _, b := range board {
				for _, c := range b {
					if c == '#' {
						fmt.Print("#")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println("")
			}
			break
		}
		t++
		if done == 0 {
			break
		}
	}

}

// position=< 9,  1> velocity=< 0,  2>
