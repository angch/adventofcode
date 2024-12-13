package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

type diskstruct struct {
	id  int
	len int
	pos int
}

func visualize(spans *list.List) {
	fmt.Println("\n------------------")
	for e := spans.Front(); e != nil; e = e.Next() {
		v := e.Value.(*diskstruct)
		fmt.Print("(", v.id, ":", v.len, ") ")
	}
	fmt.Println("\n")
	for e := spans.Front(); e != nil; e = e.Next() {
		v := e.Value.(*diskstruct)
		for range v.len {
			if v.id == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(v.id)
			}
		}
	}
	fmt.Println()
}

func day9(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	spans := list.New()
	spans2 := make([]*diskstruct, 0, 100)

	files := 0

	id := 0
	for scanner.Scan() {
		t := scanner.Text()
		for i := 0; i < len(t); i++ {
			l := int(t[i] - '0')
			if l > 0 {
				s := &diskstruct{
					id:  id,
					len: l,
				}
				s2 := *s
				spans.PushBack(s)
				spans2 = append(spans2, &s2)
			}
			i++

			if i < len(t) {
				l = int(t[i] - '0')
				if l > 0 {
					s := &diskstruct{
						id:  -1,
						len: l,
					}
					s2 := *s
					spans.PushBack(s)
					spans2 = append(spans2, &s2)
				}
			}
			id++
		}
		files = id
	}
	// spans2 := make([]diskstruct, spans.Len())
	// for e := spans.Front(); e != nil; e = e.Next() {
	// 	v := e.Value.(*diskstruct)
	// 	fmt.Print("(", v.id, ":", v.len, ") ")
	// }
	// fmt.Println()

	for {
		l, r := spans.Front(), spans.Back()

		i := 0
		for ; l != nil && l.Value.(*diskstruct).id != -1; l = l.Next() {
			i++
		}
		j := spans.Len()
		for ; r != nil && r.Value.(*diskstruct).id == -1; r = r.Prev() {
			j--
		}
		if i >= j {
			break
		}

		lv := l.Value.(*diskstruct)
		rv := r.Value.(*diskstruct)
		lv.id = rv.id
		if lv.len > rv.len {
			spaceleft := lv.len - rv.len
			lv.len = rv.len
			spaceleftStruct := &diskstruct{id: -1, len: spaceleft}
			spans.InsertAfter(spaceleftStruct, l)
			spans.Remove(r)
		} else if lv.len < rv.len {
			rv.len -= lv.len
		} else {
			spans.Remove(r)
		}

		last := spans.Back()
		if last.Value.(*diskstruct).id == -1 {
			spans.Remove(last)
		}

		// visualize(spans)
	}

	// fmt.Println(spans)
	blockn := 0
	for e := spans.Front(); e != nil; e = e.Next() {
		v := e.Value.(*diskstruct)
		for range v.len {
			part1 += blockn * v.id
			blockn++
		}
	}

	// idx := make(map[int]*diskstruct)
	idx := make([]*diskstruct, id)
	fsm := []*diskstruct{}
	blockn = 0
	for _, v := range spans2 {
		v.pos = blockn
		blockn += v.len
		if v.id == -1 {
			fsm = append(fsm, v)
		} else {
			v2 := v
			idx[v.id] = v2
		}
		// visualize(spans2)
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
	logperf := false
	if logperf {
		pf, _ := os.Create("default.pgo")
		err := pprof.StartCPUProfile(pf)
		if err != nil {
			log.Fatal(err)
		}
		defer pf.Close()
	}
	t1 := time.Now()
	part1, part2 := day9("test.txt")
	fmt.Println(part1, part2)
	if part1 != 1928 || part2 != 2858 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day9("input.txt"))
	if logperf {
		pprof.StopCPUProfile()
	}

	fmt.Println("Elapsed time:", time.Since(t1))
}
