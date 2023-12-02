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
	limit := map[string]int{"red": 12, "green": 13, "blue": 14}
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	gameid := 0
	for scanner.Scan() {
		t := scanner.Text()
		gameid++

		_, r2, _ := strings.Cut(t, ":")
		sets := strings.Split(r2, ";")

		invalid := false
		l2 := make(map[string]int)
		for _, v := range sets {
			cubes := strings.Split(v, ",")
			for _, c := range cubes {
				num, col := 0, ""
				fmt.Sscanf(c, "%d %s", &num, &col)
				l2[col] = max(l2[col], num)
				if num > limit[col] {
					invalid = true
				}
			}
		}
		if !invalid {
			part1 += gameid
		}
		part2 += l2["red"] * l2["green"] * l2["blue"]
	}
	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day2("test.txt")
	if part1 != 8 || part2 != 2286 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day2("input.txt"))
}
