package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if true {
		fileName := "input.txt"
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		prog := ParseIntCode(scanner.Text())

		fmt.Println("Part 1:")
		PaintRunner(prog, 0)
		fmt.Println("Part 2:")
		PaintRunner(prog, 1)
	}
}
