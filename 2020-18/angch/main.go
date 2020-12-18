package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func do(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	stack := make([]int, 0)
	opstack := make([]string, 0)
	_ = stack
	sum := 0
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}

		l = strings.ReplaceAll(l, "(", " ( ")
		l = strings.ReplaceAll(l, ")", " ) ")
		l = strings.ReplaceAll(l, "  ", " ")
		toks := strings.Split(l, " ")
		num := 0
		op := ""
		for _, v := range toks {
			if v == "" {
				continue
			}
			log.Println(v, num, stack, op, opstack)
			if v == "*" || v == "/" || v == "+" || v == "-" {
				op = v
				continue
			}

			if v == "(" {
				stack = append(stack, num)
				opstack = append(opstack, op)
				op = ""
				num = 0
				continue
			}
			if v == ")" {
				num1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				op = opstack[len(opstack)-1]
				opstack = opstack[:len(opstack)-1]

				switch op {
				case "+":
					num1 += num
				case "-":
					num1 -= num
				case "*":
					num1 *= num
				case "/":
					num1 /= num
				case "":
					num1 = num
				}
				// stack = append(stack, num1)
				num = num1
				continue
			}

			num1, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			if op == "" {
				// stack = append(stack, num)
				num = num1
			} else {
				// num1 := stack[len(stack)-1]
				// stack = stack[:len(stack)-1]
				switch op {
				case "+":
					num1 += num
				case "-":
					num1 -= num
				case "*":
					num1 *= num
				case "/":
					num1 /= num
				}
				// stack = append(stack, num1)
				num = num1
			}
			// log.Print(" num", num)
		}
		log.Println("num", num)
		sum += num
	}
	ret1 = sum

	return ret1, ret2
}

func dotok(toks []string) int {
	stack := make([]int, 0)
	opstack := make([]string, 0)
	op := ""
	num := 0
	for _, v := range toks {
		if v == "" {
			continue
		}
		if v == "*" || v == "/" || v == "+" || v == "-" {
			op = v
			continue
		}

		num1, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		if op == "" {
			// stack = append(stack, num)
			num = num1
		} else {
			// num1 := stack[len(stack)-1]
			// stack = stack[:len(stack)-1]
			switch op {
			case "+":
				num1 += num
			case "-":
				num1 -= num
			case "*":

				num1 *= num
			}
			// stack = append(stack, num1)
			num = num1
		}
		// log.Print(" num", num)
	}

	for {
		if len(opstack) == 0 {
			break
		}
		op = opstack[len(opstack)-1]
		opstack = opstack[:len(opstack)-1]

		num1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch op {
		case "+":
			num1 += num
		case "-":
			num1 -= num
		case "*":
			num1 *= num
		case "":
			num1 = num
		}
		// stack = append(stack, num1)
		num = num1
	}

	return num
}

func do2(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	stack := make([]int, 0)
	_ = stack
	sum := 0
	k := 0
	for scanner.Scan() {
		l := scanner.Text()
		_ = l
		k++

		if l == "" {
			break
		}

		l = strings.ReplaceAll(l, "(", " ( ")
		l = strings.ReplaceAll(l, ")", " ) ")
		l = strings.ReplaceAll(l, "  ", " ")
		toks := strings.Split(l, " ")

		tokbuffer := make(chan string)
		log.Println("Input:", toks)
		log.Println("Goroutine #1")
		go func() {
			opstack := make([]string, 0)
			for _, v := range toks {
				if v == "" {
					continue
				}
				if v == "(" {
					opstack = append(opstack, v)
					log.Println("RPN:", v, opstack)
					continue
				}
				if v == ")" {
					for len(opstack) > 0 {
						top := opstack[len(opstack)-1]
						if top != "(" {
							// tokbuffer = append(tokbuffer, top)
							log.Println("RPN: push", top)
							tokbuffer <- top
							opstack = opstack[:len(opstack)-1]
						} else if top == "(" {
							opstack = opstack[:len(opstack)-1]
							break
						}
					}
					log.Println("RPN:", v, opstack)
					continue
				}
				if v == "+" || v == "-" || v == "*" {
					for len(opstack) > 0 {
						top := opstack[len(opstack)-1]

						if (top == "-" || top == "+") && top != "(" {
							// tokbuffer = append(tokbuffer, top)
							log.Println("RPN: push", top)
							tokbuffer <- top
							opstack = opstack[:len(opstack)-1]
							continue
						}

						break
					}
					opstack = append(opstack, v)
					// }
					log.Println("RPN:", v, opstack)
					continue
				}

				// tokbuffer = append(tokbuffer, v)
				log.Println("RPN: push", v)
				tokbuffer <- v
				// log.Println("RPN:", v, tokbuffer, opstack)
			}

			for len(opstack) > 0 {
				top := opstack[len(opstack)-1]
				opstack = opstack[:len(opstack)-1]
				tokbuffer <- top
			}
			log.Println("RPN: Close channel", opstack)
			close(tokbuffer)
		}()

		stack := make([]int, 0)
		log.Println("\t\t\t\tGoroutine #2")
		for v := range tokbuffer {
			log.Println("\t\t\t\teval:", v, stack)
			if v == "*" || v == "+" || v == "-" {
				n1 := stack[len(stack)-1]
				n2 := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				if v == "+" {
					n3 := n1 + n2
					stack = append(stack, n3)
				} else if v == "-" {
					n3 := n1 - n2
					stack = append(stack, n3)
				} else if v == "*" {
					n3 := n1 * n2
					stack = append(stack, n3)
				}
				log.Println("\t\t\t\teval:", v, stack)
				continue
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			stack = append(stack, num)
			// log.Println("eval:", v, stack)
		}
		num := stack[0]

		log.Println("Result:", num)
		sum += num
	}
	ret1 = sum

	return ret1, ret2
}

func main() {
	// log.Println(do("test.txt"))
	// log.Println(do2("test2.txt"))
	// log.Println(do2("test2.txt"))
	// log.Println(do2("test3.txt"))
	// log.Println(do("test2.txt"))
	// log.Println(do("input.txt"))
	log.Println(do2("input.txt"))
}
