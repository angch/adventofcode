package main

import (
	"fmt"
	"log"
)

func day20(top int) (int, int) {
	part1, part2 := 0, 0

	house := make([]int, 1000000)
	house2 := make([]int, 10000000)

	for i := 1; ; i++ {
		// max := i * 11
		for j, count2 := 0, 0; j < len(house); j += i {

			if count2 < 50 {
				count2++
				house2[j] += i * 11

				if house2[j] >= top && part2 == 0 && j > 0 {
					part2 = j
					// log.Println("y", j)
					return part1, part2
				}
			}

			house[j] += i * 10
			if house[j] >= top && part1 == 0 && j > 0 {
				part1 = j
				// log.Println("x", j)
				// return part1, part2
			}
			if part1 > 0 && part2 > 0 {
				return part1, part2
			}
		}
		if i >= len(house) {
			break
		}

	}
	// fmt.Println(house)
	return -1, -1

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// fmt.Println(day20(150))
	fmt.Println(day20(36000000))
	// 831600
}
