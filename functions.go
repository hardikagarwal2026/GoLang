package main

import "fmt"

func plus(a int,b int) int{
	return a+b
}

func plusplus(a, b, c int) int{
	return a+b+c
}


func main() {
    a := plus(1, 2)
	b := plusplus(1,2,3)
	fmt.Println(a)
	fmt.Println(b)
}