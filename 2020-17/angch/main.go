package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Vector struct {
	X, Y, Z, W int
}

func neighbour(board map[Vector]bool, x, y, z int) int {
	count := 0
	for z1 := z - 1; z1 <= z+1; z1++ {
		for y1 := y - 1; y1 <= y+1; y1++ {
			for x1 := x - 1; x1 <= x+1; x1++ {
				if x1 == x && y1 == y && z1 == z {
					continue
				}
				if board[Vector{x1, y1, z1, 0}] {
					count++
				}
			}
		}
	}
	return count
}
func neighbour2(board map[Vector]bool, x, y, z, w int) int {
	count := 0
	for w1 := w - 1; w1 <= w+1; w1++ {
		for z1 := z - 1; z1 <= z+1; z1++ {
			for y1 := y - 1; y1 <= y+1; y1++ {
				for x1 := x - 1; x1 <= x+1; x1++ {
					if x1 == x && y1 == y && z1 == z && w1 == w {
						continue
					}
					if board[Vector{x1, y1, z1, w1}] {
						count++
					}
				}
			}
		}
	}
	return count
}

func do(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	board := make(map[Vector]bool)
	row := 0
	maxX, maxY, maxZ := 0, 0, 0
	minX, minY, minZ := 0, 0, 0
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		// log.Println(l)
		for k, v := range l {
			if v == '#' {
				board[Vector{k, row, 0, 0}] = true
			}
		}
		maxX = len(l)
		maxY = row
		row++
	}

	for steps := 1; steps < 7; steps++ {
		board2 := make(map[Vector]bool)
		minX2, minY2, minZ2 := 0, 0, 0
		maxX2, maxY2, maxZ2 := 0, 0, 0
		count2 := 0
		for z := minZ - 1; z <= maxZ+1; z++ {
			// fmt.Println("z", z)
			for y := minY - 1; y <= maxY+1; y++ {
				for x := minX - 1; x <= maxX+1; x++ {
					if board[Vector{x, y, z, 0}] {
						n := neighbour(board, x, y, z)
						if n == 2 || n == 3 {
							board2[Vector{x, y, z, 0}] = true
							count2++
						}
					} else {
						n := neighbour(board, x, y, z)
						if n == 3 {
							board2[Vector{x, y, z, 0}] = true
							count2++
						}
					}
					if board2[Vector{x, y, z, 0}] {
						if x < minX2 {
							minX2 = x
						}
						if x > maxX2 {
							maxX2 = x
						}
						if y < minY2 {
							minY2 = y
						}
						if y > maxY2 {
							maxY2 = y
						}
						if z < minZ2 {
							minZ2 = z
						}
						if z > maxZ2 {
							maxZ2 = z
						}
					} else {
						// fmt.Print(".")
					}
				}
				// fmt.Println()
			}
		}
		// minZ--
		if false {
			for z := minZ2; z <= maxZ2; z++ {
				fmt.Println("\nz ", z)
				for y := minY2; y <= maxY2; y++ {
					for x := minX2; x <= maxX2; x++ {
						if board2[Vector{x, y, z, 0}] {
							fmt.Print("#")
						} else {
							fmt.Print(".")
						}
					}
					fmt.Println()
				}
			}
			log.Println(board)
		}
		board = board2
		minX = minX2
		maxX = maxX2
		minY = minY2
		maxY = maxY2
		minZ = minZ2
		maxZ = maxZ2
		ret1 = count2
	}

	return ret1, ret2
}

func do2(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	board := make(map[Vector]bool)
	row := 0
	maxX, maxY, maxZ, maxW := 0, 0, 0, 0
	minX, minY, minZ, minW := 0, 0, 0, 0
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		// log.Println(l)
		for k, v := range l {
			if v == '#' {
				board[Vector{k, row, 0, 0}] = true
			}
		}
		maxX = len(l)
		maxY = row
		row++
	}

	for steps := 1; steps < 7; steps++ {
		board2 := make(map[Vector]bool)
		minX2, minY2, minZ2, minW2 := 0, 0, 0, 0
		maxX2, maxY2, maxZ2, maxW2 := 0, 0, 0, 0
		count2 := 0
		for w := minW - 1; w <= maxW+1; w++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				// fmt.Println("z", z)
				for y := minY - 1; y <= maxY+1; y++ {
					for x := minX - 1; x <= maxX+1; x++ {
						if board[Vector{x, y, z, w}] {
							n := neighbour2(board, x, y, z, w)
							if n == 2 || n == 3 {
								board2[Vector{x, y, z, w}] = true
								count2++
							}
						} else {
							n := neighbour2(board, x, y, z, w)
							if n == 3 {
								board2[Vector{x, y, z, w}] = true
								count2++
							}
						}
						if board2[Vector{x, y, z, w}] {
							if x < minX2 {
								minX2 = x
							}
							if x > maxX2 {
								maxX2 = x
							}
							if y < minY2 {
								minY2 = y
							}
							if y > maxY2 {
								maxY2 = y
							}
							if z < minZ2 {
								minZ2 = z
							}
							if z > maxZ2 {
								maxZ2 = z
							}
							if w < minW2 {
								minW2 = w
							}
							if w > maxW2 {
								maxW2 = w
							}
						} else {
							// fmt.Print(".")
						}
					}
					// fmt.Println()
				}
			}
		}
		// minZ--
		if false {
			for z := minZ2; z <= maxZ2; z++ {
				fmt.Println("\nz ", z)
				for y := minY2; y <= maxY2; y++ {
					for x := minX2; x <= maxX2; x++ {
						if board2[Vector{x, y, z, 0}] {
							fmt.Print("#")
						} else {
							fmt.Print(".")
						}
					}
					fmt.Println()
				}
			}
			log.Println(board)
		}
		board = board2
		minX = minX2
		maxX = maxX2
		minY = minY2
		maxY = maxY2
		minZ = minZ2
		maxZ = maxZ2
		minW = minW2
		maxW = maxW2
		ret2 = count2
	}

	return ret1, ret2
}

func main() {
	log.Println(do("test.txt"))
	log.Println(do2("test.txt"))
	// log.Println(do("test2.txt"))
	log.Println(do("input.txt"))
	log.Println(do2("input.txt"))
}
