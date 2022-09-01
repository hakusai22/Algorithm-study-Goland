package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	t1 := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute())
	fmt.Println(t1.Format("2006-01-02 15:04:05"))
	diff := t2.Sub(t1)
	fmt.Println(diff)
	fmt.Println(diff.Minutes(), diff.Seconds())
	fmt.Println(now.Unix())
}
