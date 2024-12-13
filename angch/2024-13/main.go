package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strconv"
	"time"
)

// Function to calculate the determinant of a 2x2 matrix
func det(a, b, c, d int) int {
	return a*d - b*c
}

func day13(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	regex := regexp.MustCompile(`Button (\w+): X\+([0-9]+), Y\+([0-9]+)`)
	regex2 := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)

	for scanner.Scan() {
		l1 := scanner.Text()
		scanner.Scan()

		l2 := scanner.Text()
		scanner.Scan()

		l3 := scanner.Text()
		scanner.Scan()
		buttons := [][2]int{}
		b1, c1 := 0, 0

		// Parse: Button A: X+31, Y+16

		m := regex.FindStringSubmatch(l1)
		b1, _ = strconv.Atoi(m[2])
		c1, _ = strconv.Atoi(m[3])
		buttons = append(buttons, [2]int{b1, c1})

		m = regex.FindStringSubmatch(l2)
		b1, _ = strconv.Atoi(m[2])
		c1, _ = strconv.Atoi(m[3])
		buttons = append(buttons, [2]int{b1, c1})

		m = regex2.FindStringSubmatch(l3)
		b1, _ = strconv.Atoi(m[1])
		c1, _ = strconv.Atoi(m[2])
		prize := [2]int{b1, c1}

		x1, x2 := buttons[0][0], buttons[1][0]
		y1, y2 := buttons[0][1], buttons[1][1]
		x3, y3 := prize[0], prize[1]

		// Calculate the determinant of the coefficient matrix
		D := det(x1, x2, y1, y2)
		// Calculate the determinants for a and b
		Da := det(x3, x2, y3, y2)
		Db := det(x1, x3, y1, y3)

		// Solve for a and b using Cramer's rule
		a := Da / D
		b := Db / D

		if Da%D == 0 && Db%D == 0 {
			// fmt.Printf("The solution is: a = %d, b = %d\n", a, b)
			part1 += a*3 + b
		}

		x3, y3 = x3+10000000000000, y3+10000000000000
		// Calculate the determinants for a and b
		Da = det(x3, x2, y3, y2)
		Db = det(x1, x3, y1, y3)
		// Solve for a and b using Cramer's rule
		a = Da / D
		b = Db / D

		if Da%D == 0 && Db%D == 0 {
			// fmt.Printf("The solution is: a = %d, b = %d\n", a, b)
			part2 += a*3 + b
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
	part1, part2 := day13("test.txt")
	fmt.Println(part1, part2)
	if part1 != 480 || part2 != 875318608908 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day13("input.txt"))
	if logperf {
		pprof.StopCPUProfile()
	}

	fmt.Println("Elapsed time:", time.Since(t1))
}
