package main

import (
	"fmt"
	"log"
	"time"
)

type Cup struct {
	Id   int
	Next *Cup
}

var debug = false

func do(input string, games int, ncups int) (ret1 string, ret2 int) {

	// cups := make([]int, len(input))
	c := &Cup{}
	first := c

	prev := c
	// cupmap := make([]*Cup, ncups)
	cupmap := make(map[int](*Cup))
	for _, v := range input {
		c.Id = int(byte(v - '0'))
		cupmap[c.Id] = c
		prev = c
		c.Next = &Cup{}
		c = c.Next
	}

	// log.Println("start at", len(input)+1)
	for v := len(input) + 1; v <= ncups; v++ {
		c.Id = v
		cupmap[c.Id] = c

		prev = c
		c.Next = &Cup{}
		c = c.Next
	}
	// log.Println("map is", len(cupmap))

	prev.Next = first
	c.Next = first
	// prev := c
	c = first
	min := int(9999)
	max := int(0)
	if games > 1000000 {
		min = 1
		max = ncups
	}

	for _, v := range input {
		_ = v
		// log.Printf("%d %+v\n", v, *c)
		if min > c.Id {
			min = c.Id
		}
		if max < c.Id {
			max = c.Id
		}
		c = c.Next
	}

	c = first
	// log.Println("minmax", min, max)
	pickedup := make([]*Cup, 3)

	t1 := time.Now()
	for moves := 1; moves <= games; moves++ {
		if time.Since(t1) > 5*time.Second {
			log.Println("Move: ", moves)
			t1 = time.Now()
		}
		if debug {
			log.Println("Move", moves)
			c2 := first
			for {
				if c2.Id == c.Id {
					fmt.Printf("(%d) ", c2.Id)
				} else {
					fmt.Printf("%d ", c2.Id)
				}
				c2 = c2.Next
				if c2 == first {
					break
				}
			}
			fmt.Println()
		}

		prev = c
		_ = prev
		c2 := c.Next

		for i := 0; i < 3; i++ {
			pickedup[i] = c2
			if debug {
				fmt.Println("Picked up:", c2.Id)
			}
			c2 = c2.Next
		}
		prev.Next = c2

		if debug {
			log.Println("dump")
			c2 := first
			for {
				if c2.Id == c.Id {
					fmt.Printf("(%d) ", c2.Id)
				} else {
					fmt.Printf("%d ", c2.Id)
				}
				c2 = c2.Next
				if c2 == first {
					break
				}
			}
			fmt.Println()
		}

		dest := c.Id - 1
		if dest < min {
			dest = max
		}
	a:
		for i := 0; i < 3; i++ {
			// log.Println("Check", dest, pickedup[i].Id)
			if dest == pickedup[i].Id {
				dest--
				if dest < min {
					// log.Println("Roll")
					dest = max
				}
				goto a
			}
		}
		if debug {
			log.Println("Dest", dest)
		}

		c2 = c

		if false {
			for c2.Id != dest {
				// prev = c
				c2 = c2.Next
			}
		} else {
			c2 = cupmap[dest]
		}

		if c2 == nil {
			log.Println("nil", dest)
		}

		next := c2.Next
		for i := 0; i < 3; i++ {
			c2.Next = pickedup[i]
			c2 = c2.Next
		}
		c2.Next = next

		c = c.Next
		if debug {
			log.Println("dump2")
			c2 := first
			for {
				if c2.Id == c.Id {
					fmt.Printf("(%d) ", c2.Id)
				} else {
					fmt.Printf("%d ", c2.Id)
				}
				c2 = c2.Next
				if c2 == first {
					break
				}
			}
			fmt.Println()
		}
		first = c
	}

	for first.Id != 1 {
		first = first.Next
	}
	c = first.Next

	if ncups < 100 {
		for c != first {
			ret1 += fmt.Sprint(c.Id)
			c = c.Next
		}
	} else {
		// log.Println(first.Id, c.Id, c.Next.Id)
		ret2 = c.Id * c.Next.Id
	}

	return ret1, ret2
}

func main() {
	// log.Println(do("389125467", 10))
	// log.Println(do("389125467", 100))
	// log.Println(do("389125467", 10000000, 1000000))
	// log.Println(do("389125467", 100, 10))
	// log.Println(do("219347865", 100))
	log.Println(do("219347865", 100, 9))
	log.Println(do("219347865", 10000000, 1000000))
}
