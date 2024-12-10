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
	spans2 := make([]diskstruct, len(spans))
	copy(spans2, spans)

	for {
		l, r := 0, len(spans)-1
		for ; l < r && spans[l].id != -1; l++ {
		}
		for ; l < r && spans[r].id == -1; r-- {
		}
		if l >= r {
			break
		}

		spans[l].id = spans[r].id
		if spans[l].len > spans[r].len {
			spaceleft := spans[l].len - spans[r].len
			spans[l].len = spans[r].len
			spans = append(spans[:l+1], append([]diskstruct{{id: -1, len: spaceleft}}, spans[l+1:]...)...)
			spans = append(spans[:r+1], spans[r+2:]...)
		} else if spans[l].len < spans[r].len {
			spans[r].len -= spans[l].len
		} else {
			spans = append(spans[:r], spans[r+1:]...)
		}
		if spans[len(spans)-1].id == -1 {
			spans = spans[:len(spans)-1]
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
			if idx[i].len <= fsm[j].len && idx[i].pos > fsm[j].pos {
				idx[i].pos = fsm[j].pos
				fsm[j].pos += idx[i].len
				fsm[j].len -= idx[i].len
				break
			}
		}
	}

	for _, v := range idx {
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
