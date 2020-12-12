package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func do(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	dx, dy := 1, 0
	x, y := 0, 0
	// board := make([][]byte, 0)
	for scanner.Scan() {
		l := scanner.Text()
		d, n := ' ', 0
		fmt.Sscanf(l, "%c%d", &d, &n)
		// fmt.Println("x", string(d), n)

		switch d {
		case 'F':
			x = x + n*dx
			y = y + n*dy
		case 'N':
			y -= n
		case 'S':
			y += n
		case 'E':
			x += n
		case 'W':
			x -= n
		case 'R':
			for n >= 90 {
				dx, dy = -dy, dx
				n -= 90
			}
		case 'L':
			for n >= 90 {
				dx, dy = dy, -dx
				n -= 90
			}
		}
		// fmt.Println(x, y, dx, dy)
	}

	// Math.abs
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func do2(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	x, y := 0, 0
	dx, dy := 10, -1
	for scanner.Scan() {
		l := scanner.Text()
		d, n := ' ', 0
		fmt.Sscanf(l, "%c%d", &d, &n)
		// fmt.Println("x", string(d), n)

		switch d {
		case 'F':
			x = x + n*dx
			y = y + n*dy
		case 'N':
			dy -= n
		case 'S':
			dy += n
		case 'E':
			dx += n
		case 'W':
			dx -= n
		case 'R':
			for n >= 90 {
				dx, dy = -dy, dx
				n -= 90
			}
		case 'L':
			for n >= 90 {
				dx, dy = dy, -dx
				n -= 90
			}
		}
		// fmt.Println(x, y, wx, wy)
	}

	// Math.abs
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func main() {
	log.Println(do("test.txt"))
	log.Println(do2("test.txt"))
	// log.Println(do("test2.txt"))
	// log.Println(do("test3.txt"))
	log.Println(do("input.txt"))
	log.Println(do2("input.txt"))
}
