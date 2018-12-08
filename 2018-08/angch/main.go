package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Node struct {
	ID       rune
	Children []*Node
	MetaData []int
	Value    int
}

func readNode(scanner *bufio.Scanner, startRune rune) (*Node, int, int) {
	n := Node{
		ID:       startRune,
		Children: make([]*Node, 0),
		MetaData: make([]int, 0),
		Value:    0,
	}

	scanner.Scan()
	numChildren, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	numMetaData, _ := strconv.Atoi(scanner.Text())
	startRune++
	totalAdded := 1
	metaDataSum := 0

	nodeValue := 0

	//log.Println("rn", n, numChildren, numMetaData)
	for ; numChildren > 0; numChildren-- {
		nodesAdded, added, moreMeta := readNode(scanner, startRune)
		n.Children = append(n.Children, nodesAdded)
		startRune += rune(added)
		totalAdded += added
		metaDataSum += moreMeta
		//nodeValue +=
	}
	for ; numMetaData > 0; numMetaData-- {
		scanner.Scan()
		md, _ := strconv.Atoi(scanner.Text())
		n.MetaData = append(n.MetaData, md)
		metaDataSum += md

		//log.Println(" r3", string(n.ID), "nc", numChildren)
		if len(n.Children) > 0 {
			if md <= len(n.Children) && md > 0 {
				//log.Println("add ", md-1)
				nodeValue += n.Children[md-1].Value
			}
		} else {
			nodeValue += md
		}
	}

	n.Value = nodeValue
	//log.Println("r2", string(n.ID), nodeValue, n.MetaData)

	return &n, totalAdded, metaDataSum
}

func main() {
	fileName := "input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//re := regexp.MustCompile((`(\d+)`))
	n, _, sum := readNode(scanner, 'A')
	log.Println(sum, n.Value)
}
