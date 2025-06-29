package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	nums := []int{1, 2, 3}
	sum(nums...)

	n1 := []int{4, 5, 6}
	for a, vals := range n1 { //a returns index and vals hold value
		fmt.Println(a, ":", vals)
	}
}
