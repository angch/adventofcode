package main

import "fmt"

func do(step int, times int) {
	buf := make([]int, 0)
	//buf2 := make([]int, 0)
	buf = append(buf, 0)

	cur := 0
	i := 0

	for {
		i++
		cur = (cur + step) % len(buf)
		//fmt.Print(i, buf, cur)

		buf2, buf3 := buf[:cur+1], buf[cur+1:]
		//fmt.Print(" left:", buf[:cur+1], "right:", buf[cur+1:])

		// gotcha
		if true {
			buf4 := make([]int, 0)
			buf4 = append(buf4, buf2...)
			buf4 = append(buf4, i)
			for _, j := range buf3 {
				buf4 = append(buf4, j)
			}
			buf = buf4
		} else {
			buf4 := append(buf2, i)
			buf4 = append(buf4, buf3...)
			buf = buf4
		}
		cur = (cur + 1) % len(buf)
		//fmt.Println(buf, cur)
		if i >= times {
			break
		}
	}
	//fmt.Println(buf)
	fmt.Println(i, buf[(cur+1)%len(buf)])
}

func do2(step int, times int) {
	buf := make([]int, 0)
	//buf2 := make([]int, 0)
	buf = append(buf, 0)
	pos := make([]int, times+1)
	for i := range pos {
		pos[i] = -1
	}
	pos[0] = 0

	cur := 0
	i := 0

	for {
		i++
		cur = (cur + step) % i
		//fmt.Print(i, buf, cur)

		//fmt.Print(" left:", buf[:cur+1], "right:", buf[cur+1:])

		pos[i] = cur
		//fmt.Println(i, cur, len(buf), i+1)
		cur = (cur + 1) % (i + 1)
		//fmt.Println(buf, cur)
		if i >= times {
			//fmt.Println(pos)
			lastPos := 0
			for k, v := range pos {
				if v == 0 {
					//fmt.Println(buf[k+1])
					lastPos = k
				}
			}
			fmt.Println(lastPos)
			break
		}
	}
	//fmt.Println(buf)
	//fmt.Println(i, buf[(cur+1)%len(buf)])
}

func main() {
	do(3, 10)
	do(356, 2017)
	do2(356, 50000000)
	//do(356, 2017)
}
