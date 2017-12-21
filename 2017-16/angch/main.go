package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type order struct {
	O []byte
	m []int
}

func NewOrder(i int) order {
	s := make([]byte, i)
	m := make([]int, i)
	for i := range s {
		s[i] = 'a' + uint8(i)
		m[i] = i
	}
	return order{O: s, m: m}
}

type instr struct {
	op byte
	p1 int8
	p2 int8
}

func (o *order) do(input string, times int) {
	inputs := strings.Split(input, ",")

	prog := make([]instr, len(inputs))

	for k, v := range inputs {
		op := ' '
		p1, p2 := 0, 0
		fmt.Sscanf(v, "%c%d/%d", &op, &p1, &p2)
		switch op {
		case 's':
			s := len(o.O) - p1
			prog[k] = instr{'s', int8(s), 0}
		case 'x':
			prog[k] = instr{'x', int8(p1), int8(p2)}
		case 'p':
			var pp1, pp2 byte
			fmt.Sscanf(v, "%c%c/%c", &op, &pp1, &pp2)
			prog[k] = instr{'p', int8(pp1), int8(pp2)}
		default:
			log.Fatal(v)
		}
	}
	//fmt.Println(prog)

	ar1 := [16]byte{}
	ar2 := [16]byte{}

	for i := range o.O {
		ar1[i] = o.O[i]
	}

	seen := make(map[([16]byte)]([16]byte), 0)
	for i := 0; i < times; i++ {
		if false && i%1000000 == 0 {
			fmt.Println(i)
		}
		_, exists := seen[ar1]
		if exists {
			ar1 = seen[ar1]
			continue
		} else {
			ar3 := ar1
			for _, v := range prog {
				switch v.op {
				case 's':
					src1, dst1 := int(v.p1), 0
					for ; dst1 < len(o.O); dst1++ {
						ar2[dst1] = ar1[src1]
						src1++
						if src1 >= len(o.O) {
							src1 = 0
						}
					}
					ar1 = ar2
				case 'x':
					ar1[v.p1], ar1[v.p2] = ar1[v.p2], ar1[v.p1]
				case 'p':
					var pp1, pp2 byte
					p1, p2 := -1, -1
					pp1 = byte(v.p1)
					pp2 = byte(v.p2)

					for k, i := range ar1 {
						if i == pp1 {
							p1 = k
						}
						if i == pp2 {
							p2 = k
						}
						if p1 != -1 && p2 != -1 {
							break
						}
					}
					ar1[p1], ar1[p2] = ar1[p2], ar1[p1]
				}
			}
			seen[ar3] = ar1
		}
	}
	for i := range o.O {
		o.O[i] = ar1[i]
	}
}

func main() {
	input_, _ := ioutil.ReadFile("test.txt")
	input := string(input_)

	init := NewOrder(5)
	init.do(input, 2)
	fmt.Println(string(init.O))

	if true {
		input_, _ = ioutil.ReadFile("input.txt")
		input = string(input_)
		real := NewOrder(16)
		real.do(input, 1000000000)
		fmt.Println(string(real.O))
	}

}
