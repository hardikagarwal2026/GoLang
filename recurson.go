=package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	a := fact(5)
	fmt.Println(a)

	var fibbonaci func(n int) int
	fibbonaci = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fibbonaci(n-1) + fibbonaci(n-2)
		}
	}

	b := fibbonaci(2)
	fmt.Println(b)
}
