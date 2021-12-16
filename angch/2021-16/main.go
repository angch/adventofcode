package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type packet struct {
	Version int
	Id      int
	Num     int
}

type BitStream struct {
	shift  int
	buffer []byte
}

func (bs *BitStream) Skip() int {
	if bs.shift > 0 {
		fmt.Println("skip", 4-bs.shift)
		bs.shift = 0
		bs.buffer = bs.buffer[1:]
		return 4 - bs.shift
	}
	fmt.Println("no skip")
	return 0
}

func (bs *BitStream) GetBits(n int) int {
	out := 0

	for ; n > 0; n-- {
		if bs.shift >= 4 {
			bs.shift = 0
			bs.buffer = bs.buffer[1:]
		}
		hex := 0
		if len(bs.buffer) == 0 {
			return out
		}
		fmt.Sscanf(string(bs.buffer[0]), "%x", &hex)
		// fmt.Printf("Hex is %x\n", hex)
		out <<= 1
		if hex&8 > 0 {
			out |= 1
			// fmt.Println("1")
		} else {
			// fmt.Println("0")
		}
		hex &= 0x7
		hex <<= 1
		bs.buffer[0] = byte(fmt.Sprintf("%01x", hex)[0])
		bs.shift++
	}
	return out
}

func (bs *BitStream) ReadPacket() (packet, int, int) {
	consumed := 0

	p := packet{}
	p.Version = bs.GetBits(3)
	p.Id = bs.GetBits(3)
	consumed += 6
	versions := 0

	// fmt.Printf("%+v", p)

	if p.Id == 4 {
		// the binary number is padded with leading zeroes until its length is a multiple of four bits,
		more := true
		num := 0

		for more {
			more = bs.GetBits(1) == 1
			num <<= 4
			num |= bs.GetBits(4)
			consumed += 5
		}
		// bs.Skip()
		p.Num = num
		// fmt.Println("xxxlite", num)
	} else {
		fmt.Printf("%+v\n", p)
		lengthId := bs.GetBits(1)
		consumed += 1
		fmt.Println("lengthid", lengthId)
		// fmt.Printf("%+v\n", bs)
		subpackets, length := 0, 0
		if lengthId == 0 {
			length = bs.GetBits(15)
			consumed += 15
		} else {
			subpackets = bs.GetBits(11)
			consumed += 11
		}
		ssub2 := []packet{}
		if length > 0 {
			for length > 0 {
				fmt.Println("subpackets length!", length)
				sub2, v, consumedsub := bs.ReadPacket()
				length -= consumedsub
				consumed += consumedsub
				versions += v
				ssub2 = append(ssub2, sub2)
			}
		} else {
			fmt.Println("subpackets!", subpackets)
			for i := 0; i < subpackets; i++ {
				sub2, v, consumedsub := bs.ReadPacket()
				consumed += consumedsub
				fmt.Printf(" %d consumed %d\n", i, consumedsub)
				versions += v
				ssub2 = append(ssub2, sub2)
			}
		}

		if p.Id == 0 {
			fmt.Println("sum")
			sum := 0
			for _, sub := range ssub2 {
				fmt.Println(" ", sub.Num)
				sum += sub.Num
			}
			p.Num = sum
		} else if p.Id == 1 {
			fmt.Println("product")
			mult := 1
			for _, sub := range ssub2 {
				mult *= sub.Num
			}
			p.Num = mult
		} else if p.Id == 2 {
			fmt.Println("min")
			min := -99999
			for _, sub := range ssub2 {
				if min == -99999 {
					min = sub.Num
				} else if min > sub.Num {
					min = sub.Num
				}
			}
			p.Num = min
		} else if p.Id == 3 {
			fmt.Println("max")
			max := -99999
			for _, sub := range ssub2 {
				if max == -99999 {
					max = sub.Num
				} else if max < sub.Num {
					max = sub.Num
				}
			}
			p.Num = max
		} else if p.Id == 5 {
			// greater
			if ssub2[0].Num > ssub2[1].Num {
				p.Num = 1
			} else {
				p.Num = 0
			}
		} else if p.Id == 6 {
			// less
			if ssub2[0].Num < ssub2[1].Num {
				p.Num = 1
			} else {
				p.Num = 0
			}
		} else if p.Id == 7 {
			// eq
			if ssub2[0].Num == ssub2[1].Num {
				p.Num = 1
			} else {
				p.Num = 0
			}
		}
	}
	versions += p.Version
	return p, versions, consumed
}

func day16(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	bs := BitStream{}
	for scanner.Scan() {
		t := scanner.Text()
		bs.shift = 0
		bs.buffer = []byte(t)
	}

	p, versions, bits := bs.ReadPacket()
	fmt.Printf("Consumed %d bits\n", bits)

	part1, part2 := 0, 0
	part1 = versions
	part2 = p.Num
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	// day14("test.txt")
	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	day16("input.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
