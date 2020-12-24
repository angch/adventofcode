package main

import (
	"bufio"
	"log"
	"os"
)

type Hex struct {
	I, J, K int
}

func Cube(i, j, k int) Hex {
	return Hex{i, j, k}
}

var dir = map[string]Hex{
	"e":  Cube(+1, -1, 0),
	"ne": Cube(+1, 0, -1),
	"nw": Cube(0, +1, -1),
	"w":  Cube(-1, +1, 0),
	"sw": Cube(-1, 0, +1),
	"se": Cube(0, -1, +1),
}

func do(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	current := Hex{}
	hexmap := make(map[Hex]bool)

	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		p := ""
		current = Hex{}
		for _, v := range l {
			if v == 'e' || v == 'w' {
				d := p + string(v)
				dh, ok := dir[d]
				if !ok {
					log.Fatal(dh)
				}
				current.I += dh.I
				current.J += dh.J
				current.K += dh.K
				p = ""
			} else {
				p = string(v)
			}
		}
		hexmap[current] = !hexmap[current]

	}

	// log.Println(hexmap)

	for _, v := range hexmap {
		if v {
			ret1++
		}
	}

	for day := 1; day <= 100; day++ {
		hexmap2 := make(map[Hex]bool)
		min, max := minmax(hexmap)
		for i := min.I - 1; i <= max.I+1; i++ {
			for j := min.J - 1; j <= max.J+1; j++ {
				for k := min.K - 1; k <= max.K+1; k++ {
					c := Hex{i, j, k}
					n := neighbour(hexmap, c)
					if hexmap[c] {
						if n == 0 || n > 2 {
							// hexmap2[c] = false
						} else {
							hexmap2[c] = true
						}
					} else {
						if n == 2 {
							hexmap2[c] = true
						} else {
							// hexmap2[c] = false
						}
					}
				}
			}
		}
		hexmap = hexmap2

		count := 0
		for _, v := range hexmap {
			if v {
				count++
			}
		}
		ret2 = count
		log.Println("Day", day, count)

	}

	return ret1, ret2
}

func minmax(hexmap map[Hex]bool) (Hex, Hex) {
	min := Hex{}
	max := Hex{}
	first := true
	for k, v := range hexmap {
		if !v {
			continue
		}
		if first {
			first = false
			min = k
			max = k
			continue
		}
		if min.I > k.I {
			min.I = k.I
		}
		if min.J > k.J {
			min.J = k.J
		}
		if min.K > k.K {
			min.K = k.K
		}
		if max.I < k.I {
			max.I = k.I
		}
		if max.J < k.J {
			max.J = k.J
		}
		if max.K < k.K {
			max.K = k.K
		}
	}
	return min, max
}

func neighbour(hexmap map[Hex]bool, c Hex) int {
	count := 0
	for _, v := range dir {
		d := c
		d.I += v.I
		d.J += v.J
		d.K += v.K

		if hexmap[d] {
			count++
		}
	}
	return count
}
func main() {
	log.Println(do("test.txt"))
	log.Println(do("input.txt"))
}
