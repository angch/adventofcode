package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Race struct {
	Time     int
	Distance int
}

func day6(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	races := []Race{}
	scanner.Scan()
	t := scanner.Text()
	fields := strings.Fields(t)
	for i := 1; i < len(fields); i++ {
		t, _ := strconv.Atoi(fields[i])
		races = append(races, Race{t, 0})
	}
	scanner.Scan()
	t = scanner.Text()
	fields = strings.Fields(t)
	for i := 1; i < len(fields); i++ {
		t, _ := strconv.Atoi(fields[i])
		races[i-1].Distance = t
	}
	// log.Println(races)

	part1 = 1
	for _, v := range races {
		win := 0
		for elapsed := range v.Time {
			d := elapsed * (v.Time - elapsed)
			if d > v.Distance {
				win++
			}
		}
		part1 *= win
	}

	max := Race{0, 0}
	timeStr := ""
	distStr := ""
	for _, v := range races {
		timeStr += strconv.Itoa(v.Time)
		distStr += strconv.Itoa(v.Distance)
	}
	max.Time, _ = strconv.Atoi(timeStr)
	max.Distance, _ = strconv.Atoi(distStr)
	for elapsed := range max.Time {
		d := elapsed * (max.Time - elapsed)
		if d > max.Distance {
			part2++
		}
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day6("test.txt")
	// log.Println(part1, part2)
	if part1 != 288 || part2 != 71503 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day6("input.txt"))
	fmt.Println("Elapsed", time.Since(t1))
}
