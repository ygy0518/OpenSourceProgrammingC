package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	var b bool
	var s string

	my_school_account := 5
	println(my_school_account)

	q7x1abc := 5
	println(q7x1abc)

	fmt.Println(f, b, s, i)
	fmt.Printf("%f %t %s %d\n", f, b, s, i)
	f = 2.7
	i = 3
	fmt.Print("\n\n", f <= float64(i), "\n")
}
