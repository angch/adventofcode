package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Dst   int
	Src   int
	Range int
}
type Records []Range
type MapRecord struct {
	From    string
	To      string
	Records Records
}

func (m *MapRecord) Translate(i int) int {
	for _, v := range m.Records {
		// log.Println("Check", v.Src, v.Src+v.Range, i)
		if i >= v.Src && i < v.Src+v.Range {
			return v.Dst + i - v.Src
		}
	}
	return i
}

func findLocation(maps map[string]MapRecord, v int) int {
	curType := "seed"
	for curType != "location" {
		// fmt.Printf("%s %d, ", curType, v)
		r := maps[curType]
		v = r.Translate(v)
		curType = maps[curType].To
	}
	return v
	// fmt.Printf("%s %d, ", curType, v)
	// fmt.Println()
}

func day5(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	seeds := []int{}
	maps := make(map[string]MapRecord)
	scanner.Scan()
	t := scanner.Text()
	for _, v := range strings.Fields(t) {
		if v != "seeds:" {
			a, err := strconv.Atoi(v)
			if err == nil {
				seeds = append(seeds, a)
			}
		}
	}

	for scanner.Scan() {
		t := scanner.Text()
		if strings.HasSuffix(t, " map:") {
			t = strings.TrimSuffix(t, " map:")
			from, to, ok := strings.Cut(t, "-to-")
			if !ok {
				continue
			}
			r := Records{}
			for scanner.Scan() {
				t := scanner.Text()
				if t == "" {
					break
				}
				rng := Range{}
				fmt.Sscanf(t, "%d %d %d", &rng.Dst, &rng.Src, &rng.Range)
				r = append(r, rng)
			}
			// log.Printf("%+v\n", r)
			maps[from] = MapRecord{
				From:    from,
				To:      to,
				Records: r,
			}
		}
	}
	_ = maps

	lowest := -1
	for _, v := range seeds {
		v = findLocation(maps, v)
		if lowest == -1 || v < lowest {
			lowest = v
		}
		// log.Println(lowest)
	}
	part1 = lowest

	lowest = -1
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		rng := seeds[i+1]
		for j := start; j < start+rng; j++ {
			v := findLocation(maps, j)
			if lowest == -1 || v < lowest {
				lowest = v
			}
		}
	}
	part2 = lowest

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day5("test.txt")
	if part1 != 35 || part2 != 46 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day5("input.txt"))
}
