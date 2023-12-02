package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day2(file string) (int, int) {
	part1, part2 := 0, 0
	r, g, b := 12, 13, 14
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	gameid := 0
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		gameid++

		l1, r2, _ := strings.Cut(t, ":")
		sets := strings.Split(r2, ";")

		invalid := false
		red, green, blue := 0, 0, 0
		for k, v := range sets {

			cubes := strings.Split(v, ",")
			for _, c := range cubes {
				num, col := 0, ""
				fmt.Sscanf(c, "%d %s", &num, &col)
				if false {
					fmt.Println(l1, k, num, col)
				}
				switch col {
				case "blue":
					blue = max(blue, num)
					if num > b {
						invalid = true
					}
				case "red":
					red = max(red, num)
					if num > r {
						invalid = true
					}
				case "green":
					green = max(green, num)
					if num > g {
						invalid = true
					}
				default:
					log.Fatal(col)
				}
			}
		}
		power := red * green * blue
		log.Println("Game ", gameid, invalid, power, red, green, blue)
		if !invalid {
			part1 += gameid
		}
		part2 += power

	}
	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day2("test.txt")
	fmt.Println(part1, part2)
	if part1 != 8 || part2 != 2286 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day2("input.txt"))
}
