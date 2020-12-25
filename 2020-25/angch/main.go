package main

import (
	"log"
)

func do(a, b int) (ret1 int, ret2 int) {

	// 20201227
	sub1, sub2 := 7, 7
	target := a
	loop1, loop2 := 0, 0
	for loop1 = 1; sub1 != target; loop1++ {
		sub1 *= 7
		sub1 = sub1 % 20201227
		// log.Println(ret1, sub1)
	}

	for loop2 = 1; sub2 != b; loop2++ {
		sub2 *= 7
		sub2 = sub2 % 20201227
	}

	e1, e2 := b, a
	sub1, sub2 = b, a
	log.Println("xform", sub1, loop1)
	for i := 1; i < loop1; i++ {
		sub1 *= e1
		sub1 = sub1 % 20201227
	}
	for i := 1; i < loop2; i++ {
		sub2 *= e2
		sub2 = sub2 % 20201227
	}
	log.Println(sub1, sub2)

	return ret1, ret2
}

func main() {
	log.Println(do(5764801, 17807724))
	//5764801
	log.Println(do(8184785, 5293040))
}
