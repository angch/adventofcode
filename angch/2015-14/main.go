package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rdeer struct {
	name         string
	speed        int
	flyduration  int
	restduration int

	state        int
	durationleft int
	distance     int
	score        int
}

func day14(file string, duration int) (int, int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rdeers := make([]*rdeer, 0)
	for scanner.Scan() {
		t := scanner.Text()
		r := rdeer{}
		fmt.Sscanf(t,
			"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &r.name, &r.speed, &r.flyduration, &r.restduration,
		)
		rdeers = append(rdeers, &r)
	}
	for i := 0; i < duration; i++ {
		for _, deer := range rdeers {
			switch deer.state {
			case 0: // start
				deer.state = 1 // flying
				deer.durationleft = deer.flyduration
				// deer.distance += deer.speed
			case 1:
				deer.durationleft--
				deer.distance += deer.speed
				if deer.durationleft == 0 {
					deer.state = 2 // resting
					deer.durationleft = deer.restduration
				}
			case 2:
				deer.durationleft--
				if deer.durationleft == 0 {
					deer.state = 1 // flying
					deer.durationleft = deer.flyduration
				}
			}
		}
		maxdistance := 0
		for _, deer := range rdeers {
			if deer.distance > maxdistance {
				maxdistance = deer.distance
			}
		}
		for _, deer := range rdeers {
			if deer.distance == maxdistance {
				deer.score++
			}
		}

	}
	part1, part2 := 0, 0
	for _, deer := range rdeers {
		if deer.distance > part1 {
			part1 = deer.distance
		}
		if deer.score > part2 {
			part2 = deer.score - 1 // offbyone error
		}
	}

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(day14("test.txt", 1000))
	fmt.Println(day14("input.txt", 2503))
}
