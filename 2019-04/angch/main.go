package main

import (
	"fmt"
)

func part1(from, to int) (int, int) {
	count1, count2 := 0, 0
a:
	for i := from; i <= to; i++ {
		prev, adj, adj2, adjcount := 10, 0, 0, 0
		for k := i; k > 0; k /= 10 {
			j := k % 10
			if j == prev {
				adj = 1
				adjcount++
			} else {
				if adjcount == 1 {
					adj2 = 1
				}
				adjcount = 0
			}

			if j > prev {
				continue a
			}
			prev = j
		}

		if adjcount == 1 {
			adj2 = 1
		}
		count1, count2 = count1+adj, count2+adj2
	}
	return count1, count2
}

func main() {
	//
	fmt.Println(part1(111111, 111111))
	fmt.Println(part1(223450, 223450))
	fmt.Println(part1(123789, 123789))

	fmt.Println(part1(112233, 112233))
	fmt.Println(part1(123444, 123444))
	fmt.Println(part1(111122, 111122))

	fmt.Println(part1(367479, 893698))
	// not 335
}
