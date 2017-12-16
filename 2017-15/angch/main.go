package main

import "fmt"

type gen struct {
	state  int
	factor int
	a      int
}

func (g *gen) tick() {
	g.state = (g.state * g.factor) % 2147483647
}

func (g *gen) tick2() {
	for {
		g.state = (g.state * g.factor) % 2147483647
		if g.state&g.a == 0 {
			return
		}
	}
}

func main() {
	gena := gen{65, 16807, 3}
	genb := gen{8921, 48271, 7}

	gena.state = 512
	genb.state = 191

	// Part 1
	count := 0
	for i := 0; i < 40000000; i++ {
		gena.tick()
		genb.tick()
		if gena.state&65535 == genb.state&65535 {
			count++
		}
		//fmt.Println(gena.state, genb.state, count)
	}
	fmt.Println(count)

	// Part 2
	count = 0
	gena.state = 65
	genb.state = 8921
	for i := 0; i < 5000000; i++ {
		gena.tick2()
		genb.tick2()
		if gena.state&65535 == genb.state&65535 {
			count++
		}
		//fmt.Println(gena.state, genb.state, count)
	}
	fmt.Println(count)
}
