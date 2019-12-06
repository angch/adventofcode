package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

type edge struct {
	from string
	to   string
}

func parse(i string) []edge {
	e := make([]edge, 0)
	lines := bytes.NewBufferString(i)
	scanner := bufio.NewScanner(lines)
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
		indirect1 := 0
		o := k
		for o != root {
			o = orbits[o]
			indirect++
			indirect1++
		}

		if nodes[k] == nil {
			// log.Println("n", k)
			// nodes[k] = make([]string, 0)
		}
		// log.Println(k, indirect1)
	}
	log.Println("Orbits", indirect)
	// direct := len(edges)

	if orbits["YOU"] == "" || orbits["SAN"] == "" {
		return
	}

	// Part 2
	you_path := make([]string, 0)
	o := orbits["YOU"]
	count := 0
	for o != root {
		// log.Println("o", o)
		you_path = append(you_path, o)
		o = orbits[o]
		count++
		// if count > 10000 {
		// 	return
		// }
	}
	log.Println(you_path)
	san_path := make([]string, 0)
	o, e := orbits["SAN"]
	// log.Println(orbits)
	if e == false {
		log.Fatal("Not exists")
	}
	count = 0
	for o != root {
		// log.Println("o", o)
		san_path = append(san_path, o)
		o = orbits[o]
		count++
		if count > 1000000 {
			return
		}
	}
	log.Println(san_path)

	i, j := len(you_path)-1, len(san_path)-1
	for {
		if you_path[i-1] == san_path[j-1] {
			i--
			j--
		} else {
			log.Println(i, j, i+j)
			return
		}
	}
}

func main() {
	if true {
		// Part 1
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

		// log.Printf("%+v\n", parse(test))
		a := parse(test)
		part1(a)

		// fmt.Println(part1(test))
	}
	if true {

		fileName := "input.txt"

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		e := make([]edge, 0)
		for scanner.Scan() {
			t := strings.TrimSpace(scanner.Text())
			split := strings.Split(t, ")")
			if len(split) > 1 {
				from, to := split[0], split[1]
				e = append(e, edge{from: from, to: to})
			}
		}
		part1(e)

		// fmt.Println("Part 1", part1(prog, 1))
		// fmt.Println("Part 2", part1(prog, 5))
	}
}
