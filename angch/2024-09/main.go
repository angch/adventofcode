package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type diskstruct struct {
	id  int
	len int
	pos int
}

func day9(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	spans := []diskstruct{}
	files := 0
	for scanner.Scan() {
		t := scanner.Text()
		id := 0
		for i := 0; i < len(t); i++ {
			s := diskstruct{
				id:  id,
				len: int(t[i] - '0'),
			}
			spans = append(spans, s)
			i++
			if i < len(t) {
				s := diskstruct{
					id:  -1,
					len: int(t[i] - '0'),
				}
				spans = append(spans, s)
			}
			id++
		}
		files = id
	}
	// fmt.Println(spans)
	spans2 := make([]diskstruct, len(spans))
	copy(spans2, spans)

	for {
		l, r := 0, len(spans)-1
		for ; l < r && spans[l].id != -1; l++ {
			// Find empty slot
		}
		for ; l < r && spans[r].id == -1; r-- {
		}
		if l >= r {
			break
		}

		spans[l].id = spans[r].id
		if spans[l].len > spans[r].len {
			// fmt.Println("lr1", l, r)
			spaceleft := spans[l].len - spans[r].len
			spans[l].len = spans[r].len
			spans = append(spans[:l+1], append([]diskstruct{{id: -1, len: spaceleft}}, spans[l+1:]...)...)
			spans = append(spans[:r+1], spans[r+2:]...)
		} else if spans[l].len < spans[r].len {
			// fmt.Println("lr2", l, r)
			spans[r].len -= spans[l].len
			// spans[l].len = spans[r].len
		} else {
			// fmt.Println("lr3", l, r, spans[l], spans[r])
			spans = append(spans[:r], spans[r+1:]...)
		}
		// fmt.Println(spans)
		if spans[len(spans)-1].id == -1 {
			// fmt.Println("rmtail")
			spans = spans[:len(spans)-1]
			// fmt.Println(spans)
		}

	}

	// fmt.Println(spans)
	blockn := 0
	for _, v := range spans {
		for range v.len {
			part1 += blockn * v.id
			blockn++
		}
	}

	idx := make(map[int]*diskstruct)
	fsm := []diskstruct{}
	blockn = 0
	for _, v := range spans2 {
		v.pos = blockn
		blockn += v.len
		if v.id == -1 {
			fsm = append(fsm, v)
		} else {
			v2 := v
			idx[v.id] = &v2
		}
	}

	for i := files - 1; i >= 0; i-- {
		for j := 0; j < len(fsm); j++ {
			// if i == 2 {
			// 	log.Println("p2,", i, j, fsm[j].pos, idx[i])
			// 	log.Println(fsm, idx)
			// }
			if idx[i].len <= fsm[j].len && idx[i].pos > fsm[j].pos {
				idx[i].pos = fsm[j].pos
				fsm[j].pos += idx[i].len
				fsm[j].len -= idx[i].len
				// log.Println(fsm, idx)
				break
			}
		}
	}
	// log.Println(idx)

	for _, v := range idx {
		// log.Println(*v)
		pos := v.pos
		for range v.len {
			part2 += pos * v.id
			pos++
		}
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day9("test.txt")
	fmt.Println(part1, part2)
	if part1 != 1928 || part2 != 2858 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day9("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
