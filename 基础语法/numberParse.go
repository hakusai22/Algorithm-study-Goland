package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(a)

	b, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(b)

	c, _ := strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(c)

	d, _ := strconv.Atoi("123")
	fmt.Println(d)

	e, _ := strconv.Atoi("AAA")
	fmt.Println(e)

}
