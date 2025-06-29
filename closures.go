package main

import "fmt"

func nextInt() func() int {
	i := 0
	return func() int {
		i= i+1
		return i
	}
}

func main() {
	a := nextInt()
    fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	
	b:= nextInt()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	
}