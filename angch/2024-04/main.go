package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day4(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	board := make(map[[2]int]byte)
	y := 0
	maxx := 0
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		for x, v := range t {
			board[[2]int{x, y}] = byte(v)
		}
		y++
		maxx = len(t)
	}
	maxy := y
	word := "XMAS"
	dirs := [][]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {-1, 0}, {0, -1}}

	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			if board[[2]int{x, y}] != word[0] {
				continue
			}

			for _, dir := range dirs {
				found := true
				for i := 1; i < len(word); i++ {
					nx := x + i*dir[0]
					ny := y + i*dir[1]
					if nx < 0 || nx >= maxx || ny < 0 || ny >= maxy || board[[2]int{nx, ny}] != word[i] {
						found = false
						break
					}
				}
				if found {
					part1++
				}
			}

		}
	}

	// Lol, this gave so much problems with off by one errors.
	// Probably better if I figured out rotate 45 degrees instead
	// of 90 degrees and getting the deltas wrong
	crosses := [][][2]int{
		// M        M        S       S
		{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}},
		{{0, -1}, {-1, 0}, {0, 1}, {0, -1}},
	}

	for i := 0; i < 3; i++ {
		// rotate
		k1 := crosses[i*2]
		k := make([][2]int, len(k1))
		copy(k, k1)
		for j := 0; j < len(k1); j++ {
			// Rotate 90 degrees
			// (x,y) -> (-y,x)
			k[j][0], k[j][1] = -k[j][1], k[j][0]
		}
		crosses = append(crosses, k)

		k1 = crosses[i*2+1]
		k = make([][2]int, len(k1))
		copy(k, k1)
		for j := 0; j < len(k1); j++ {
			// Rotate 90 degrees
			// (x,y) -> (-y,x)
			k[j][0], k[j][1] = -k[j][1], k[j][0]
		}
		crosses = append(crosses, k)
	}
	for _, z := range crosses {
		fmt.Println(z)
	}

	for y := 1; y < maxy-1; y++ {
		for x := 1; x < maxx-1; x++ {
			if board[[2]int{x, y}] != 'A' {
				continue
			}
			// log.Println("Found A", x, y)

			for _, c := range crosses {
				// log.Println("Check", c)
				if board[[2]int{x + c[0][0], y + c[0][1]}] == 'M' &&
					board[[2]int{x + c[1][0], y + c[1][1]}] == 'M' &&
					board[[2]int{x + c[2][0], y + c[2][1]}] == 'S' &&
					board[[2]int{x + c[3][0], y + c[3][1]}] == 'S' {
					log.Println("Found X at", x, y, c)
					part2++
					continue
				}
			}

		}
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day4("test.txt")
	fmt.Println(part1, part2)
	if part1 != 18 || part2 != 9 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day4("input.txt"))
}
