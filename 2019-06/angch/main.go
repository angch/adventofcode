package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

type edge struct {
	from string
	to   string
}

func parse(r io.Reader) []edge {
	e := make([]edge, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		split := strings.Split(t, ")")
		if len(split) > 1 {
			from, to := split[0], split[1]
			e = append(e, edge{from: from, to: to})
		}
	}
	return e
}

func part1(edges []edge) {
	nodes := make(map[string][]string)
	orbits := make(map[string]string)
	for _, e := range edges {
		_, exists := nodes[e.from]
		if !exists {
			nodes[e.from] = make([]string, 0)
		}
		nodes[e.from] = append(nodes[e.from], e.to)
		orbits[e.to] = e.from
	}
	if false {
		for k, v := range nodes {
			log.Printf("%s: %+v\n", k, v)
		}
	}
	root := ""
	for k, _ := range nodes {
		if orbits[k] == "" {
			root = k
		}
	}
	log.Println("root is", root)
	// Not needed
	indirect := 0
	for k, _ := range orbits {
		indirect1 := 0 // This one is for debugging
		o := k
		for o != root {
			o = orbits[o]
			indirect++
			indirect1++
		}
		// Debugging
		// log.Println(k, indirect1)
	}
	log.Println("Orbits", indirect)

	if orbits["YOU"] == "" || orbits["SAN"] == "" {
		// Only a test
		return
	}

	// Part 2
	you_path := make([]string, 0)
	o := orbits["YOU"]
	for o != root {
		you_path = append(you_path, o)
		o = orbits[o]
	}
	log.Println(you_path)
	o = orbits["SAN"]
	san_path := make([]string, 0)
	for o != root {
		san_path = append(san_path, o)
		o = orbits[o]
	}
	log.Println(san_path)

	// We traced from both santa to root, and you to root
	// Now we backtrace from root to a path that diverges
	// (We can codegolf this to even less lines, but it gets confusing)
	i, j := len(you_path)-1, len(san_path)-1
	for {
		if you_path[i-1] != san_path[j-1] {
			break
		}
		i--
		j--
	}
	// Distance of the two paths added:
	log.Println(i, j, i+j)
}

func main() {
	if true {
		test := `COM)B
		B)C
		C)D
		D)E
		E)F
		B)G
		G)H
		D)I
		E)J
		J)K
		K)L`
		lines := bytes.NewBufferString(test)
		part1(parse(lines))
	}
	if true {
		fileName := "input.txt"

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		part1(parse(file))
	}
}
