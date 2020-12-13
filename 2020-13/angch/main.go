package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func do(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	l := scanner.Text()
	t1 := 0
	fmt.Sscanf(l, "%d", &t1)
	scanner.Scan()
	l2 := scanner.Text()
	notes := strings.Split(l2, ",")

	min := 0
	mint := 999999999999
	waitt := 0
	bus := make([]int, 0)
	for _, v := range notes {
		if v == "x" {
			bus = append(bus, -1)
			continue
		}
		bid, _ := strconv.Atoi(v)

		t2 := t1 % bid
		t3 := t1 + (bid - t2)
		log.Println("bus", bid, t2, t3)
		if t3 < mint {
			min = bid
			mint = t3
			waitt = t2
		}
		bus = append(bus, bid)
	}
	log.Println(min, mint, waitt)

	ret1 := (mint - t1) * min

	ret2 := 0
	// sort.Ints(bus)
	// log.Println(bus)
	// log.Fatal("")
a:
	for t := bus[0]; ; t += bus[0] {
		// target := bus[0]
		// if t > 1000 {
		// 	break
		// }
		for k, v := range bus {
			// if k == 0 {
			// 	continue
			// }
			// log.Println("x", t, v, t%v, v-k, t+k, (t+k)%v)
			if v == -1 {
				continue
			}
			if (t+k)%v != 0 {
				// log.Fatal("xx")
				continue a
			}
		}
		ret2 = t
		break
	}
	// 0 7 0-0 0
	// 1 13 13-1 12

	return ret1, ret2
}

func do2(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	l := scanner.Text()
	t1 := 0
	fmt.Sscanf(l, "%d", &t1)
	scanner.Scan()
	l2 := scanner.Text()
	notes := strings.Split(l2, ",")

	min := 0
	mint := 999999999999

	bus := make([]int, 0)
	for _, v := range notes {
		if v == "x" {
			bus = append(bus, -1)
			continue
		}
		bid, _ := strconv.Atoi(v)

		t2 := t1 % bid
		t3 := t1 + (bid - t2)
		// log.Println("bus", bid, t2, t3)
		if t3 < mint {
			min = bid
			mint = t3
		}
		bus = append(bus, bid)
	}
	// log.Println(min, mint, waitt)

	ret1 := (mint - t1) * min

	ret2 := 0
	// sort.Ints(bus)
	// log.Println(bus)
	// log.Fatal("")
	// gcd := make([]int, 0)
	largest := 0
	fudge := 0
	for k, v := range bus {
		if v > largest {
			largest = v
			fudge = k
		}
	}
	log.Println("lar", largest, bus)

	tt1 := time.Now()

	// starthere := 100000000000000
	starthere := (200000000000000/largest)*largest - fudge

a:
	// for t := largest - fudge; ; t += largest {
	for t := starthere; ; t += largest {
		if time.Since(tt1) > 10*time.Second {
			log.Println(t)
			tt1 = time.Now()
		}

		for k, v := range bus {
			if v == -1 {
				continue
			}
			if (t+k)%v != 0 {
				continue a
			}
		}
		ret2 = t
		break
	}
	// 493075315291
	// 100000000000000
	// 28252201718241
	// 55709012670191
	// 114458982288401
	// log.Println(gcd)
	// 0 7 0-0 0
	// 1 13 13-1 12

	return ret1, ret2
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func do3(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	l := scanner.Text()
	t1 := 0
	fmt.Sscanf(l, "%d", &t1)
	scanner.Scan()
	l2 := scanner.Text()
	notes := strings.Split(l2, ",")

	min := 0
	mint := 999999999999

	bus := make([]int, 0)
	for _, v := range notes {
		if v == "x" {
			bus = append(bus, -1)
			continue
		}
		bid, _ := strconv.Atoi(v)

		t2 := t1 % bid
		t3 := t1 + (bid - t2)
		// log.Println("bus", bid, t2, t3)
		if t3 < mint {
			min = bid
			mint = t3
		}
		bus = append(bus, bid)
	}
	// log.Println(min, mint, waitt)

	ret1 := (mint - t1) * min

	ret2 := 0
	// sort.Ints(bus)
	// log.Println(bus)
	// log.Fatal("")
	// gcd := make([]int, 0)
	largest := 0
	fudge := 0
	for k, v := range bus {
		if v > largest {
			largest = v
			fudge = k
		}
	}
	log.Println("lar", largest, bus)

	tt1 := time.Now()

	// starthere := 100000000000000
	starthere := (300000000000000/largest)*largest - fudge
	starthere = largest - fudge

	type Bus struct {
		Id    int
		Fudge int
	}
	f := make([]int, 0)
	b := make([]int, 0)
	BB := make([]Bus, 0)
	for k, v := range bus {
		if v == -1 {
			continue
		}
		BB = append(BB, Bus{Id: v, Fudge: k})
	}
	sort.Slice(BB, func(i, j int) bool { return BB[i].Id > BB[j].Id })
	lcm := 1
	lc := make([]int, 0)
	for _, v := range BB {
		b = append(b, v.Id)
		f = append(f, v.Fudge)
		lcm = LCM(lcm, v.Id)
		lc = append(lc, lcm)
	}

	b = b[1:]
	f = f[1:]
	// lc = lc[1:]
	log.Println(lc)
	// log.Println(b, f)
	// for t := largest - fudge; ; t += largest {
	t := starthere
a:
	for {
		if time.Since(tt1) > 10*time.Second {
			log.Println(t)
			tt1 = time.Now()
		}

		for k, v := range b {
			if (t+f[k])%v != 0 {
				// lcm := GCD(v, largest)
				// jump := largest
				// jump := largest
				// log.Println(v, largest, jump, lcm)
				// t += largest
				t += lc[k]
				continue a
			}
		}
		ret2 = t
		break
	}
	// 493075315291
	// 100000000000000
	// 28252201718241
	// 55709012670191
	// 114458982288401
	// log.Println(gcd)
	// 0 7 0-0 0
	// 1 13 13-1 12

	return ret1, ret2
}

func main() {
	// log.Println(do2("test.txt"))
	// log.Println(do("test2.txt"))
	// log.Println(do2("test3.txt"))
	// log.Println(do2("test4.txt"))
	// log.Println(do3("test4.txt"))
	log.Println(do3("input.txt"))
}
