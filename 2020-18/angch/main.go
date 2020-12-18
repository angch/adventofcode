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
		num := 0

		tokbuffer := make([]string, 0)
		opstack := make([]string, 0)
		op := ""
		log.Println("input", toks)
		for _, v := range toks {
			if v == "" {
				continue
			}

			// if v == "(" {
			// 	tokbuffer = append(tokbuffer, v)
			// 	continue
			// }

			// if v == ")" {
			// 	toks := make([]string, 0)

			// 	for {
			// 		t := tokbuffer[len(tokbuffer)-1]
			// 		tokbuffer = tokbuffer[:len(tokbuffer)-1]

			// 		if t == "(" {
			// 			break
			// 		}

			// 		toks = append(toks, t)
			// 	}
			// 	log.Println(" tok", toks, num)
			// 	// num = dotok(toks)
			// } else {

			// 	tokbuffer = append(tokbuffer, v)
			// }
			// log.Print(" num", num)
			if v == "(" {
				opstack = append(opstack, v)
				log.Println(v, tokbuffer, opstack)
				continue
			}
			if v == ")" {
				for len(opstack) > 0 {
					top := opstack[len(opstack)-1]
					if top != "(" {
						tokbuffer = append(tokbuffer, top)
						opstack = opstack[:len(opstack)-1]
					} else if top == "(" {
						opstack = opstack[:len(opstack)-1]
						break
					}
				}
				log.Println(v, tokbuffer, opstack)
				continue
			}
			if v == "+" || v == "-" || v == "*" {
				if len(tokbuffer) == 0 || tokbuffer[len(tokbuffer)-1] == "(" {
					tokbuffer = append(tokbuffer, v)
				} else {
					for len(opstack) > 0 {
						top := opstack[len(opstack)-1]

						if (top == "-" || top == "+") && top != "(" {
							tokbuffer = append(tokbuffer, top)
							opstack = opstack[:len(opstack)-1]
							continue
						}

						break
					}
					opstack = append(opstack, v)
				}
				log.Println(v, tokbuffer, opstack)
				continue
			}

			tokbuffer = append(tokbuffer, v)
			if op != "" {
				tokbuffer = append(tokbuffer, op)
			}
			log.Println(v, tokbuffer, opstack)
		}

		for len(opstack) > 0 {
			top := opstack[len(opstack)-1]
			opstack = opstack[:len(opstack)-1]
			tokbuffer = append(tokbuffer, top)
		}

		log.Println(tokbuffer, opstack)

		stack := make([]int, 0)
		for _, v := range tokbuffer {
			log.Println("e ", v, stack)
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
				continue
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			stack = append(stack, num)
		}
		num = stack[0]

		// num = dotok(toks)

		log.Println(k, "num", num)
		sum += num
		// break
	}
	ret1 = sum

	//3414545442547727
	// 3881468912309
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
