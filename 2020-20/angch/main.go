package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type Tile struct {
	Id         int
	Board      []string
	Edges      map[int]bool
	Rot        int
	EdgesSlice []int
}

func (t *Tile) Rotate() {
	// t.Dump()
	log.Println("Rotating ", t.Id, t.Rot)
again:
	for {
		if t.Rot == 0 {
			log.Println("No rotate")
			return
		}

		if t.Rot >= 4 {
			if false {
				log.Println("Flipping")
				t.Dump()
			}
			t.Rot -= 4
			t.EdgesSlice = append(t.EdgesSlice[4:], t.EdgesSlice[0:4]...)

			for i, j := 0, len(t.Board)-1; i < j; i, j = i+1, j-1 {
				t.Board[i], t.Board[j] = t.Board[j], t.Board[i]
			}
			t.GetEdges()
			if false {
				log.Println("After flip")
				t.Dump()
			}
			continue again
		}

		// log.Println("Rot 90")
		// t.Dump()
		t.Rot--
		b := append(t.EdgesSlice[7:8], t.EdgesSlice[4:7]...)
		a := append(t.EdgesSlice[3:4], t.EdgesSlice[0:3]...)
		t.EdgesSlice = append(a, b...)

		board2 := make([][]byte, len(t.Board))
		board3 := make([]string, len(t.Board))
		for i := 0; i < len(t.Board); i++ {
			board2[i] = make([]byte, len(t.Board))
		}
		w1 := len(t.Board) - 1
		for y := 0; y < len(t.Board); y++ {
			for x := 0; x < len(t.Board); x++ {
				board2[y][x] = t.Board[w1-x][y]
			}
		}
		for i := 0; i < len(t.Board); i++ {
			board3[i] = string(board2[i])
		}
		t.Board = board3
		// log.Println("After rot90")
		// t.Dump()
		t.GetEdges()
	}

}

func (t *Tile) Dump() {
	for k, v := range t.Board {
		fmt.Println(k, v)
	}
	fmt.Println(t.Id, t.EdgesSlice)
}

func (t *Tile) GetEdges() []int {
	// edges := make([]int,0)

	edges := make([]int, 8)

	w := len(t.Board)
	for k := 0; k < w; k++ {
		for i := 0; i < 8; i++ {
			edges[i] <<= 1
		}

		if t.Board[0][k] == '#' {
			edges[0] |= 1
		}
		if t.Board[k][w-1] == '#' {
			edges[1] |= 1
		}
		if t.Board[w-1][k] == '#' {
			edges[2] |= 1
		}
		if t.Board[k][0] == '#' {
			edges[3] |= 1
		}

		if t.Board[0][w-k-1] == '#' {
			edges[4] |= 1
		}
		if t.Board[w-k-1][w-1] == '#' {
			edges[5] |= 1
		}
		if t.Board[w-1][w-k-1] == '#' {
			edges[6] |= 1
		}
		if t.Board[w-k-1][0] == '#' {
			edges[7] |= 1
		}

		// log.Println(t.Board[k])
	}

	log.Println(edges)
	t.EdgesSlice = edges
	return edges
}

func do(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	tiles := make(map[int]*Tile)
	edgecount := make(map[int]int)
	edgemap := make(map[int]map[int]bool)
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}

		tileid := 0
		fmt.Sscanf(l, "Tile %d:", &tileid)
		board := make([]string, 0)
		for scanner.Scan() {
			l := scanner.Text()
			if l == "" {
				break
			}
			board = append(board, l)
		}
		tiles[tileid] = &Tile{
			Id:    tileid,
			Board: board,
		}
		e := tiles[tileid].GetEdges()
		tiles[tileid].Edges = make(map[int]bool)
		for _, v := range e {
			tiles[tileid].Edges[v] = true
			edgecount[v]++
			if edgecount[v] == 1 {
				edgemap[v] = make(map[int]bool)
			}
			edgemap[v][tileid] = true
		}
	}
	// log.Println(edgecount)
	// log.Println("edgemap", edgemap)
	ret1 = 1
	firsttiles := make([]int, 0)
	tilesleft := make(map[int]*Tile)
	for k, v := range tiles {
		count := 0
		for k2 := range v.Edges {
			if edgecount[k2] == 1 {
				count++
			}
		}
		if count == 4 {
			ret1 *= k
			firsttiles = append(firsttiles, k)
		}
		tilesleft[k] = v
		log.Println(k, count, v.Edges)
	}
	sort.Ints(firsttiles)
	firsttile := firsttiles[0]

	boardW := int(math.Sqrt(float64(len(tiles))))
	log.Println(len(tiles), boardW, boardW*boardW)

	boardIndex := make([][]int, boardW)
	for k := range boardIndex {
		boardIndex[k] = make([]int, boardW)
	}

	// firsttile = 1951
	boardIndex[0][0] = firsttile
	delete(tilesleft, firsttile)
	log.Println("first tile", firsttile)

	if false {
		for k := range tiles[firsttile].Edges {
			if edgecount[k] == 2 {
				delete(edgecount, k)
			}
		}
	} else {
		prev := 0
		for k, v := range tiles[firsttile].EdgesSlice {
			prev = v
			if k == 0 {
				continue
			}
			if edgecount[v] == 2 && edgecount[prev] == 2 {
				tiles[firsttile].Rot = k - 1
			}
		}
	}
	log.Println("rot is", tiles[firsttile].Rot)
	tiles[firsttile].Rotate()

	// count := 1
	out := make([][]int, 0)
	for i := 0; i < boardW*2; i++ {
		for x := 0; x < boardW; x++ {
			for y := 0; y < boardW; y++ {
				if x+y == i {
					out = append(out, []int{x, y})
				}
			}
		}
	}
	// log.Println(out)

	placed := make(map[int]bool)
	placed[firsttile] = true
	for {
		topedge := tiles[firsttile].EdgesSlice[0]
		leftedge := tiles[firsttile].EdgesSlice[3]

		if edgecount[topedge] == 1 && edgecount[leftedge] == 1 {
			break
		}
		tiles[firsttile].Rot = 1
		tiles[firsttile].Rotate()
	}

	for tries := 0; tries < 2; tries++ {
		for _, v := range out {
			if boardIndex[v[1]][v[0]] > 0 {
				log.Println("Done, skipping", v[0], v[1])
				continue
			}

			lefttileId := 0
			if v[0] > 0 {
				lefttileId = boardIndex[v[1]][v[0]-1]
			}
			toptileId := 0
			if v[1] > 0 {
				toptileId = boardIndex[v[1]-1][v[0]]
			}

			log.Println("Board pos", v[0], v[1])

			candidates := make(map[int]int)
			toptileedge := 0
			lefttileedge := 0
			if lefttileId != 0 {
				rot := tiles[lefttileId].Rot
				e := tiles[lefttileId].EdgesSlice[rot+1]
				e2 := tiles[lefttileId].EdgesSlice[rot+1+4]
				lefttileedge = e
				log.Println("left tile", e, e2)
				if edgecount[e] > 0 {
					for k := range edgemap[e] {
						if !placed[k] {
							log.Println("candidate for left is", k)
							candidates[k] |= 1
						}
					}
				}
				if edgecount[e2] > 0 {
					for k := range edgemap[e2] {
						if !placed[k] {
							log.Println("candidate for left is", k)
							candidates[k] |= 1
						}
					}
				}
			}
			if toptileId != 0 {
				rot := tiles[toptileId].Rot
				e := tiles[toptileId].EdgesSlice[rot+2]
				e2 := tiles[toptileId].EdgesSlice[rot+4+2]
				toptileedge = e
				// lefttileedge = e

				tiles[toptileId].Rotate()
				log.Println("top tile id", toptileId, e, e2, tiles[toptileId].EdgesSlice)
				tiles[toptileId].Dump()
				if edgecount[e] > 0 {
					for k := range edgemap[e] {
						if !placed[k] {
							log.Println("candidate for top is", k)
							candidates[k] |= 2
						}
					}
				}
				if edgecount[e2] > 0 {
					for k := range edgemap[e2] {
						if !placed[k] {
							log.Println("candidate for top is", k)
							candidates[k] |= 2
						}
					}
				}
			}
			log.Println("candidates", candidates)
			if len(candidates) == 1 {
				c := 0
				for k := range candidates {
					c = k
				}
				log.Println("candidate", c, "edges", tiles[c].EdgesSlice)
			}
			c := 0

			matchMask := 3
			if v[0] == 0 {
				log.Println("foo")
				matchMask &= 2
			}
			if v[1] == 0 {
				log.Println("foo2")
				matchMask &= 1
			}
			log.Println("mask is", matchMask, v, boardW)

			if len(candidates) > 0 {
				for k, v := range candidates {
					if v == matchMask {
						c = k
					}
					break
				}
				// if c == 0 {
				// 	for k := range candidates {
				// 		c = k
				// 		break
				// 	}
				// }
			}
			if c > 0 {
				log.Println("Placing", c, tiles[c].EdgesSlice, "matching", toptileedge, lefttileedge)
				boardIndex[v[1]][v[0]] = c
				placed[c] = true

				// for rot := 0; rot < 8; rot++ {
				// 	if toptileedge == tiles[c].EdgesSlice[rot] {

				// 		// if rot > 4 {
				// 		// 	rot = rot - 4
				// 		// 	rot -= 2
				// 		// 	rot %= 4
				// 		// 	rot += 4
				// 		// } else {
				// 		// 	rot -= 2
				// 		// 	if rot < 2 {
				// 		// 		rot += 4
				// 		// 	}
				// 		// }
				// 		if rot == 2 {
				// 			rot = 4
				// 		}
				// 		if rot == 4 {

				// 		}

				// 		tiles[c].Dump()
				// 		log.Println("Rot is", rot)
				// 		tiles[c].Rot = rot
				// 		break
				// 	}
				// 	if lefttileedge == tiles[c].EdgesSlice[(rot+1)%8] {
				// 		log.Println("LRot is", rot)
				// 		tiles[c].Rot = rot
				// 		break
				// 	}
				// }

				if toptileedge > 0 {
					log.Println("Rotating", c, "to match toptile", toptileId, "edge", toptileedge)

					// tiles[toptileId].Dump()
					// tiles[c].Dump()

					// tmp := *tiles[c]
					// found := false
					// for r := 0; r < 4+4+4; r++ {
					// 	if tmp.EdgesSlice[0] == toptileedge {
					// 		tiles[c] = &tmp
					// 		// tiles[c].Rotate()
					// 		log.Println("Found rot = ", r)
					// 		found = true
					// 		break
					// 	}
					// 	if r == 4 {
					// 		log.Println("Flipping")
					// 		for i, j := 0, len(tmp.Board)-1; i < j; i, j = i+1, j-1 {
					// 			tmp.Board[i], tmp.Board[j] = tmp.Board[j], tmp.Board[i]
					// 		}
					// 		tmp.GetEdges()
					// 	} else if r == 8 {
					// 		log.Println("Flipping h")
					// 		for kk, vv := range tmp.Board {
					// 			l := []byte(vv)
					// 			for i, j := 0, len(l)-1; i < len(l); i, j = i+1, j-1 {
					// 				l[i], l[j] = l[j], l[i]
					// 			}
					// 			tmp.Board[kk] = string(l)
					// 		}
					// 		tmp.GetEdges()
					// 	} else {
					// 		tmp.Rot = 1
					// 		tmp.Rotate()
					// 	}
					// 	if tmp.EdgesSlice[0] == toptileedge {
					// 		tiles[c] = &tmp
					// 		// tiles[c].Rotate()
					// 		log.Println("Found rot = ", r)
					// 		found = true
					// 		break
					// 	}
					// 	log.Print("rot=", r)
					// 	tmp.Dump()
					// 	// if tmp.EdgesSlice[0] == FlipEdge(toptileedge) {
					// 	// 	for kk, vv := range tmp.Board {
					// 	// 		l := []byte(vv)
					// 	// 		for i, j := 0, len(l)-1; i < len(l); i, j = i+1, j-1 {
					// 	// 			l[i], l[j] = l[j], l[i]
					// 	// 		}
					// 	// 		tmp.Board[kk] = string(l)
					// 	// 	}
					// 	// 	tiles[c] = &tmp
					// 	// 	log.Println("Found rot, flipped = ", r)
					// 	// 	found = true
					// 	// 	break
					// 	// }
					// }
					tiles[toptileId].Dump()
					tiles[c].Dump()
					tmp := *tiles[c]
					found := false
					for r := 0; r < 8; r++ {
						if r == 4 {
							for i, j := 0, len(tmp.Board)-1; i < j; i, j = i+1, j-1 {
								tmp.Board[i], tmp.Board[j] = tmp.Board[j], tmp.Board[i]
							}
							tmp.GetEdges()
						} else {
							tmp.Rot = 1
							tmp.Rotate()
						}
						if tmp.EdgesSlice[0] == toptileedge {
							tiles[c] = &tmp
							log.Println("Found rot = ", r)
							found = true
							// tiles[c].Rot = r
							// tiles[c].Rotate()
							break
						}
					}
					if !found {
						log.Fatal("Not found")
					}
					tiles[c].Dump()
				}
				if lefttileedge > 0 {
					log.Println("Rotating", c, "to match lefttile", lefttileId, "edge", lefttileedge)
					tiles[lefttileId].Dump()
					tiles[c].Dump()
					tmp := *tiles[c]
					for r := 0; r < 8; r++ {
						if r == 4 {
							for i, j := 0, len(tmp.Board)-1; i < j; i, j = i+1, j-1 {
								tmp.Board[i], tmp.Board[j] = tmp.Board[j], tmp.Board[i]
							}
							tmp.GetEdges()
						} else {
							tmp.Rot = 1
							tmp.Rotate()
						}
						if tmp.EdgesSlice[3] == lefttileedge {
							tiles[c] = &tmp
							log.Println("Found rot = ", r)
							// tiles[c].Rot = r
							// tiles[c].Rotate()
							break
						}
					}
				}

				// tiles[c].Rotate()

				DumpBoard(tiles, boardIndex)
				fmt.Println()
			}
		}
	}

	for y := 0; y < boardW; y++ {
		for x := 0; x < boardW; x++ {
			fmt.Println(boardIndex)
		}
	}

	log.Println(boardIndex)
	DumpBoard(tiles, boardIndex)
	stitchedmap := StitchMap(tiles, boardIndex)
	log.Println("stiched map", stitchedmap)
	for k, v := range stitchedmap {
		fmt.Printf("%2d %s\n", k, v)
	}

	pattern := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	_ = pattern
	vecs := BoardToVec(pattern)
	count := 0

	r := 0
	stitchedmap2 := make([][]byte, len(stitchedmap))
	for {
		stitchedmap2 = make([][]byte, len(stitchedmap))
		for k, v := range stitchedmap {
			stitchedmap2[k] = []byte(v)
		}
		fmt.Println()
		for _, v := range stitchedmap2 {
			fmt.Println(string(v))
		}
		count = LocatePattern(stitchedmap2, vecs)
		if count > 0 {
			break
		}
		r++
		if r == 4 {
			stitchedmap = FlipStitchedY(stitchedmap)
		} else {
			stitchedmap = RotateStiched(stitchedmap)
		}
		if r > 8 {
			log.Println("Give up")
			break
		}
	}
	fmt.Println()

	for _, v := range stitchedmap2 {
		fmt.Println(string(v))
		for _, v2 := range v {
			if v2 == '#' {
				ret2++
			}
		}
	}

	return ret1, ret2
}

func LocatePattern(stitchedmap [][]byte, pattern []Vec) int {
	w := len(stitchedmap)
	// log.Println("w is", w)
	found := 0
yl:
	for y := 0; y < w; y++ {
	xl:
		for x := 0; x < w; x++ {

			for _, v := range pattern {
				x1 := x + v[0]
				y1 := y + v[1]
				if x1 >= w {
					continue yl
				}
				if y1 >= w {
					break yl
				}
				if y1 >= len(stitchedmap) {
					break yl
				}
				if x1 >= len(stitchedmap[y1]) {
					continue yl
				}
				if stitchedmap[y1][x1] != '#' {
					continue xl
				}
			}

			for _, v := range pattern {
				x1 := x + v[0]
				y1 := y + v[1]
				stitchedmap[y1][x1] = 'O'
			}
			found++
		}
	}
	return found
}

type Vec [2]int

func BoardToVec(m []string) []Vec {
	vecs := make([]Vec, 0)
	for y, v := range m {
		for x, v2 := range v {
			if v2 == '#' {
				vecs = append(vecs, Vec{x, y})
			}
		}
	}
	return vecs
}

func StitchMap(tiles map[int]*Tile, boardIndex [][]int) []string {
	out := make([]string, 0)
	boardW := len(boardIndex)
	firstTile := 0
	for k := range tiles {
		firstTile = k
		break
	}

	tileW := len(tiles[firstTile].Board[0])
	for y := 0; y < boardW; y++ {
		for y2 := 1; y2 < tileW-1; y2++ {
			o := ""
			for x := 0; x < boardW; x++ {
				t := boardIndex[y][x]
				if t == 0 {
					o += strings.Repeat(":", tileW-2)
				} else {
					o += tiles[t].Board[y2][1 : tileW-1]
				}
			}
			out = append(out, o)
		}
	}
	return out
}

func DumpBoard(tiles map[int]*Tile, boardIndex [][]int) {
	boardW := len(boardIndex)
	firstTile := 0
	for k := range tiles {
		firstTile = k
		break
	}
	tileW := len(tiles[firstTile].Board[0])

	for y := 0; y < boardW; y++ {
		for y2 := 0; y2 < tileW; y2++ {
			for x := 0; x < boardW; x++ {
				t := boardIndex[y][x]
				if t == 0 {
					fmt.Print(strings.Repeat(":", tileW))
					fmt.Print(" ")
				} else {
					fmt.Print(tiles[t].Board[y2])
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
	for y := 0; y < boardW; y++ {
		fmt.Println(boardIndex[y])
	}
}

func FlipEdge(i int) int {
	j := 0
	for k := 0; k < 10; k++ {
		j <<= 1
		if i&1 == 1 {
			j |= 1
		}
		i >>= 1
	}
	return j
}

func FlipStitchedY(stitchedmap []string) []string {
	out := make([][]byte, len(stitchedmap))

	for k, v := range stitchedmap {
		out[k] = make([]byte, len(v))
	}

	w := len(stitchedmap)
	for y, v := range stitchedmap {
		for x, v2 := range v {
			out[y][w-x-1] = byte(v2)
		}
	}

	out2 := make([]string, len(out))
	for k, v := range out {
		out2[k] = string(v)
	}
	return out2
}

func FlipStitchedX(stitchedmap []string) []string {
	out := make([][]byte, len(stitchedmap))

	for k, v := range stitchedmap {
		out[k] = make([]byte, len(v))
	}

	w := len(stitchedmap)
	for y, v := range stitchedmap {
		for x, v2 := range v {
			out[w-y-1][x] = byte(v2)
		}
	}

	out2 := make([]string, len(out))
	for k, v := range out {
		out2[k] = string(v)
	}
	return out2
}

func RotateStiched(stitchedmap []string) []string {
	out := make([][]byte, len(stitchedmap))

	for k, v := range stitchedmap {
		out[k] = make([]byte, len(v))
		// log.Println(len(stitchedmap), len(v))
	}

	for y, v := range stitchedmap {
		w := len(stitchedmap) - 1
		for x := range v {
			out[y][x] = byte(stitchedmap[w-x][y])
		}
	}
	/*
			w1 := len(t.Board) - 1
		for y := 0; y < len(t.Board); y++ {
			for x := 0; x < len(t.Board); x++ {
				board2[y][x] = t.Board[w1-x][y]
			}
		}
	*/

	out2 := make([]string, len(out))
	for k, v := range out {
		out2[k] = string(v)
	}
	return out2
}

func do2(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	stitchedmap := make([]string, 0)

	for scanner.Scan() {
		l := scanner.Text()
		_ = l
		stitchedmap = append(stitchedmap, l)
	}

	for _, v := range stitchedmap {
		fmt.Println(v)
	}

	// stitchedmap = FlipStitchedX(stitchedmap)
	stitchedmap = RotateStiched(stitchedmap)
	stitchedmap = FlipStitchedY(stitchedmap)

	pattern := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	vecs := BoardToVec(pattern)
	stitchedmap2 := make([][]byte, len(stitchedmap))
	for k, v := range stitchedmap {
		stitchedmap2[k] = []byte(v)
	}
	LocatePattern(stitchedmap2, vecs)

	fmt.Println()
	for _, v := range stitchedmap2 {
		fmt.Println(string(v))
	}

	// fmt.Println(stitchedmap)

	return 0, 0
}

func main() {
	// log.Println(do("test.txt"))
	// log.Println(do2("test2.txt"))

	log.Println(do("input.txt"))
}
