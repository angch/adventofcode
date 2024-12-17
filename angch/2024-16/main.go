package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Board struct {
	board map[[2]int]bool
	start [2]int
	end   [2]int
	maxx  int
	maxy  int
}

type Moves struct {
	move    [][2]int
	dirs    []int
	movemap map[[2]int]int
	score   int
}

var dirs = [4][2]int{
	{1, 0},  // east
	{0, 1},  // south
	{-1, 0}, // west
	{0, -1}, // north
}

var dirSymbol = [4]string{">", "v", "<", "^"}

func displayboard(b Board, m Moves) {
	for y := 0; y < b.maxx; y++ {
		for x := 0; x < b.maxy; x++ {
			if c, ok := m.movemap[[2]int{x, y}]; ok {
				fmt.Print(dirSymbol[c])
				continue
			}
			if _, ok := b.board[[2]int{x, y}]; ok {
				fmt.Print("#")
			} else if m.move[len(m.move)-1] == [2]int{x, y} {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	pause()
	return
}

func pause() {
	// fmt.Println("Press enter to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func day16(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	board := make(map[[2]int]bool)
	start := [2]int{0, 0}
	end := [2]int{0, 0}

	y := 0
	maxx := 0
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		for x, c := range t {
			if c == '#' {
				board[[2]int{x, y}] = true
			} else if c == 'S' {
				start = [2]int{x, y}
			} else if c == 'E' {
				end = [2]int{x, y}
			}
		}
		maxx = len(t)
		y++
	}
	fmt.Println("start", start, "end", end)

	b := Board{
		board: board,
		start: start,
		end:   end,
		maxx:  y,
		maxy:  maxx,
	}
	_ = b

	moves := Moves{
		move:  [][2]int{start},
		dirs:  []int{0},
		score: 0,
	}
	eval := []Moves{moves}
	evals := 0
	bestScore := -1
	t2 := time.Now()
	done := map[[2]int]int{}
	for len(eval) > 0 {
		evals++

		best := 999999999999
		bestIdx := 0
		// fmt.Print("Scores: ")
		for k := range len(eval) {
			// fmt.Print(k, eval[k].score, " ")
			if eval[k].score < best {
				best = eval[k].score
				bestIdx = k
			}
		}
		// fmt.Println()
		moves = eval[bestIdx]
		eval = append(eval[:bestIdx], eval[bestIdx+1:]...)

		// moves = eval[len(eval)-1]
		// eval = eval[:len(eval)-1]

		facing := moves.dirs[len(moves.move)-1]
		pos := moves.move[len(moves.move)-1]

		if moves.move[len(moves.move)-1] == end {
			if bestScore == -1 || moves.score < bestScore {
				bestScore = moves.score
				log.Fatal("best score", bestScore)
			}
			continue
		}
		if time.Since(t2) > 2*time.Second {
			log.Println("pos", pos, "facing", facing, moves.score, len(eval))
			t2 = time.Now()
		}

		// displayboard(b, moves)

		for dirIndex, d := range dirs {
			newPos := [2]int{d[0] + pos[0], d[1] + pos[1]}
			if _, ok := board[newPos]; ok {
				continue
			}

			if olddir, ok := done[newPos]; ok {
				if olddir&(1<<dirIndex) > 0 {
					continue
				}
			}

			if olddir, ok := moves.movemap[newPos]; ok {
				if olddir == dirIndex {
					continue
				}
				// _ = olddir
				continue
			}

			addScore := 1
			if (dirIndex+1)%4 == facing || (dirIndex+3)%4 == facing {
				addScore += 1000
			} else if (dirIndex+2)%4 == facing {
				addScore += 2000
			}
			// fmt.Println("Facing is ", facing, "dirIndex is", dirIndex, "addScore is", addScore)

			cloneMap := make(map[[2]int]int)
			for k, v := range moves.movemap {
				cloneMap[k] = v
			}
			cloneMap[newPos] = dirIndex

			moveClone := make([][2]int, len(moves.move))
			copy(moveClone, moves.move)
			moveClone = append(moveClone, newPos)

			dirsClone := make([]int, len(moves.dirs))
			copy(dirsClone, moves.dirs)
			dirsClone = append(dirsClone, dirIndex)

			done[newPos] = done[newPos] | (1 << dirIndex)

			eval = append(eval, Moves{
				move:    moveClone,
				dirs:    dirsClone,
				movemap: cloneMap,
				score:   moves.score + addScore,
			})
		}
	}
	log.Println("no more moves", evals)
	part1 = bestScore

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	// part1, part2 := day16("test.txt")
	// fmt.Println(part1, part2)
	// if part1 != 7036 || part2 != 0 {
	// 	log.Fatal("Test failed ", part1, part2)
	// }
	// 7036 too low
	fmt.Println(day16("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
