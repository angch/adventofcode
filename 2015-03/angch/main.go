package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func addme(m map[int]map[int]int, x, y int) {
	if m == nil {
		m = make(map[int]map[int]int)
	}
	if m[x] == nil {
		m[x] = make(map[int]int)
	}
	m[x][y]++
}

func advent03(input string) (int, int) {
	done := make(map[int]map[int]int)
	x, y := 0, 0

	addme(done, x, y)

	for _, v := range input {
		switch v {
		case '^':
			y--
		case 'v':
			y++
		case '>':
			x++
		case '<':
			x--
		}
		addme(done, x, y)
	}

	count := 0
	for _, i := range done {
		count += len(i)
	}
	return count, 0
}

func advent03b(input string) (int, int) {
	done := make(map[int]map[int]int)
	x1, y1 := 0, 0
	x2, y2 := 0, 0

	addme(done, x1, y1)
	addme(done, x2, y2)

	for k, v := range input {
		if k%2 == 0 {
			switch v {
			case '^':
				y1--
			case 'v':
				y1++
			case '>':
				x1++
			case '<':
				x1--
			}
			addme(done, x1, y1)
		} else {
			switch v {
			case '^':
				y2--
			case 'v':
				y2++
			case '>':
				x2++
			case '<':
				x2--
			}
			addme(done, x2, y2)
		}
	}

	count := 0
	for _, i := range done {
		count += len(i)
	}
	return count, 0
}

func main() {
	test1, test2 := advent03(">")
	if test1 != 2 {
		log.Fatal(test1, test2)
	}
	test1, test2 = advent03("^>v<")
	if test1 != 4 {
		log.Fatal(test1, test2)
	}
	test1, test2 = advent03("^v^v^v^v^v")
	if test1 != 2 {
		log.Fatal(test1, test2)
	}

	test1, test2 = advent03b("^v")
	if test1 != 3 {
		log.Fatal(test1, test2)
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum, sum2 := 0, 0
	for scanner.Scan() {
		t := scanner.Text()
		//log.Println(advent03(t))
		one, _ := advent03(t)
		two, _ := advent03b(t)
		sum += one
		sum2 += two
	}
	fmt.Println(sum, sum2)
}
