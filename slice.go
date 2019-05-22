package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	sl := make([]int,3)

	//fmt.Println(len(sl))

	var x string

	var n int
	var err error
	var c int

	for {

		fmt.Println("Enter the integer to add into slice or X to quit")

		_, _ = fmt.Scan(&x)


		if n , err = strconv.Atoi(x); err != nil{

			if x == "X" {

				break
			}

			continue

		}
		//fmt.Println(c)

		if c < 3 {

			sl[c] = n

		} else {
			sl = append(sl, n)
		}
		c++

		sl1:= make([]int,len(sl))
		copy(sl1,sl)

		sort.Ints(sl1)

		fmt.Println("Sorted slice after new number : ",sl1)

	}

}

