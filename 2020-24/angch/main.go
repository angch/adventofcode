package main

import (
	"bufio"
	"log"
	"os"
)

type Hex struct {
	I, J, K int16
	pad     int16
}

func Cube(i, j, k int16) Hex {
	return Hex{i, j, k, 0}
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
	neighbourcount := make(map[Hex]int)
	for c, v2 := range hexmap {
		if v2 {
			ret1++
			// neighbourcount[c] += 0

			for _, v := range dir {
				v.I += c.I
				v.J += c.J
				v.K += c.K
				neighbourcount[v]++
			}
		}
	}

	for day := 1; day <= 100; day++ {
		hexmap2 := make(map[Hex]bool)
		neighbourcount2 := make(map[Hex]int)
		for c, n := range neighbourcount {
			if hexmap[c] {
				if !(n == 0 || n > 2) {
					hexmap2[c] = true
					for _, v := range dir {
						v.I += c.I
						v.J += c.J
						v.K += c.K
						neighbourcount2[v]++
						// neighbourcount2[c] += 0
					}
				}
			} else {
				if n == 2 {
					hexmap2[c] = true
					for _, v := range dir {
						v.I += c.I
						v.J += c.J
						v.K += c.K
						neighbourcount2[v]++
						// neighbourcount2[c] += 0
					}
				}
			}
		}
		hexmap = hexmap2
		neighbourcount = neighbourcount2

		ret2 = len(hexmap)
		log.Println("Day", day, ret2, len(neighbourcount2))
	}

	return ret1, ret2
}

func main() {
	// log.Println(do("test.txt"))
	log.Println(do("input.txt"))
}
