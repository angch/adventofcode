package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	name           string
	weight         int
	children       []string
	childrenNodes  map[string]*Node
	parent         string
	childrenweight int
}

func main() {
	//log.SetOutput(ioutil.Discard)

	if false {
		test1 := advent07a("test1.txt")
		fmt.Println(test1)
		if test1 != "tknk" {
			log.Fatal("err")
		}
		o := advent07a("input.txt")
		fmt.Println(o)
	} else {
		test1 := advent07b("test1.txt")
		fmt.Println(test1)
		if test1 != 60 {
			log.Fatal("err")
		}
		o := advent07b("input.txt")
		fmt.Println(o)
	}
}

func advent07a(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	Map := make(map[string]*Node)
	for scanner.Scan() {
		t := scanner.Text()

		words := strings.Split(t, " ")

		weight := 0
		fmt.Sscanf(words[1][1:len(words[1])-1], "%d", &weight)

		c := make([]string, 0)
		if len(words) > 3 {
			for i := 3; i < len(words); i++ {
				w := words[i]
				last := len(w)
				if w[last-1] == ',' {
					last--
				}
				c = append(c, w[0:last])
			}
		}
		node := Node{
			name:     words[0],
			weight:   weight,
			children: c,
		}
		log.Printf("%#v\n", words[0])
		Map[node.name] = &node
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for name, v := range Map {
		for _, c := range v.children {
			log.Println(c)
			Map[c].parent = name
		}
	}
	//log.Println(Map)
	for k, v := range Map {
		if v.parent == "" {
			return k
		}
	}

	return " "
}

var badweightIdx string
var fixWeight int

var pong int

func (n *Node) CWeight() int {
	pong++
	if len(n.children) == 0 {
		n.childrenweight = n.weight
		return n.weight
	}
	w := n.weight

	weights := make(map[int]int)
	weightsIdx := make(map[int]string)
	for _, c := range n.childrenNodes {
		cweight := c.CWeight()
		w += cweight
		//c.childrenweight = cweight
		weights[cweight] = weights[cweight] + 1
		weightsIdx[cweight] = c.name
	}
	n.childrenweight = w

	badweight, goodweight := 0, 0
	if len(weights) > 1 {
		for k, c := range weights {
			if c == 1 {
				log.Println("badweight ", n.children, weights)
				badweightIdx = weightsIdx[k]
				badweight = k
				log.Println("badweight ", badweightIdx, k, c)
			} else {
				goodweight = k
			}
		}
		if badweight > 0 {
			//log.Println("badweight")
			fixWeight = n.childrenNodes[badweightIdx].weight - (badweight - goodweight)
			log.Println("fixweight", fixWeight)
		}
	}

	return w
}

func advent07b(fileName string) int {
	pong = 0
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	Map := make(map[string]*Node)
	badweightIdx = ""
	fixWeight = -1
	for scanner.Scan() {
		t := scanner.Text()

		words := strings.Split(t, " ")

		weight := 0
		fmt.Sscanf(words[1][1:len(words[1])-1], "%d", &weight)

		c := make([]string, 0)
		if len(words) > 3 {
			for i := 3; i < len(words); i++ {
				w := words[i]
				last := len(w)
				if w[last-1] == ',' {
					last--
				}
				c = append(c, w[0:last])
			}
		}
		node := Node{
			name:     words[0],
			weight:   weight,
			children: c,
		}
		//log.Printf("%#v\n", words[0])
		Map[node.name] = &node
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for name, v := range Map {
		v.childrenNodes = make(map[string]*Node)
		for _, c := range v.children {
			//log.Println(c)
			Map[c].parent = name
			v.childrenNodes[c] = Map[c]
		}
	}

	top := ""
	for k, v := range Map {
		if v.parent == "" {
			top = k
			break
		}
	}

	fmt.Println("parent is", top)

	//checkWeight(Map[top])
	Map[top].CWeight()

	if false {
		for k, v := range Map {
			if len(v.childrenNodes) == 0 {
				continue
			}
			fmt.Printf("%s(%d,%d) ", k, v.weight, v.childrenweight)
			for _, c := range v.childrenNodes {
				fmt.Print(c.name, "(", c.childrenweight, ") ")
			}
			fmt.Println()
		}
	}
	log.Println("Pong", pong, len(Map))
	log.Println(badweightIdx, fixWeight)
	return fixWeight
}
