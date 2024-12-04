package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func day4(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	board := make(map[[2]int]byte)
	coordsA := make([][2]int, 0)
	coordsX := make([][2]int, 0)
	y := 0
	for scanner.Scan() {
		t := scanner.Text()
		for x, v := range t {
			board[[2]int{x, y}] = byte(v)
			if v == 'A' {
				coordsA = append(coordsA, [2]int{x, y})
			} else if v == 'X' {
				coordsX = append(coordsX, [2]int{x, y})
			}
		}
		y++
	}
	word := "XMAS"
	dirs := [][]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {-1, 0}, {0, -1}}

	for _, coord := range coordsX {
		x, y := coord[0], coord[1]
	a:
		for _, dir := range dirs {
			for i := 1; i < len(word); i++ {
				if board[[2]int{x + i*dir[0], y + i*dir[1]}] != word[i] {
					continue a
				}
			}
			part1++
		}
	}

	// Lol, this gave so much problems with off by one errors.
	// Probably better if I figured out rotate 45 degrees instead
	// of 90 degrees and getting the deltas wrong
	word2 := "MMSS"
	crosses := [][][2]int{
		{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}, // Diagonal
		{{0, -1}, {-1, 0}, {0, 1}, {0, -1}},  // Vertical
	}

	// Rotate all of them 3 times to get all orientations
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ { // Diagonal, Vertical
			k1 := crosses[i*2+j]
			k := make([][2]int, len(k1)) // Hahah, footgun
			for j := 0; j < len(k1); j++ {
				// Rotate 90 degrees ie (x,y) -> (-y,x)
				k[j][0], k[j][1] = -k1[j][1], k1[j][0]
			}
			crosses = append(crosses, k)
		}
	}

	for _, coord := range coordsA {
		x, y := coord[0], coord[1]
	b:
		for _, c := range crosses {
			for i, c2 := range c {
				if board[[2]int{x + c2[0], y + c2[1]}] != word2[i] {
					continue b
				}
			}
			part2++
		}
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day4("test.txt")
	fmt.Println(part1, part2)
	if part1 != 18 || part2 != 9 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day4("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
