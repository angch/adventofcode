package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//#1 @ 1,3: 4x4
	re := regexp.MustCompile(`^#(\d+)\s*@\s(\d+),(\d+):\s*(\d+)x(\d+)$`)

	w := 1000
	fabric := make([][]int, w)
	for x := 0; x < w; x++ {
		fabric[x] = make([]int, w)
	}

	oldlog := make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()

		a := re.FindAllStringSubmatch(t, -1)

		//log.Println(a[0][1:])
		id, _ := strconv.Atoi(a[0][1])
		x, _ := strconv.Atoi(a[0][2])
		y, _ := strconv.Atoi(a[0][3])
		w, _ := strconv.Atoi(a[0][4])
		h, _ := strconv.Atoi(a[0][5])
		log.Println(id, x, y, w, h)

		for y1 := y; y1 < y+h; y1++ {
			for x1 := x; x1 < x+w; x1++ {
				//log.Println(y, x)
				fabric[y1][x1]++
			}
		}
		oldlog = append(oldlog, t)
	}
	//log.Println(fabric)
	count := 0
	for y := 0; y < w; y++ {
		for x := 0; x < w; x++ {
			if fabric[y][x] > 1 {
				count++
			}
		}
	}
	log.Println(count)

f:
	for _, t := range oldlog {
		a := re.FindAllStringSubmatch(t, -1)

		id, _ := strconv.Atoi(a[0][1])
		x, _ := strconv.Atoi(a[0][2])
		y, _ := strconv.Atoi(a[0][3])
		w, _ := strconv.Atoi(a[0][4])
		h, _ := strconv.Atoi(a[0][5])
		log.Println(id, x, y, w, h)

		for y1 := y; y1 < y+h; y1++ {
			for x1 := x; x1 < x+w; x1++ {
				if fabric[y1][x1] > 1 {
					continue f
				}
				//log.Println(y, x)
				fabric[y1][x1]++
			}
		}
		fmt.Println(id)
		return
	}

}
