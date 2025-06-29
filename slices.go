package main

import (
	"fmt"
	"slices"
)

func main(){


	var s[]string
	fmt.Println(s, s==nil, len(s)==0)

	s = make([]string, 3)
	fmt.Println(len(s))


	s[0] = "a"
	s[1] = "b"
	s[2] = "c" //setter
	fmt.Println(s[0])

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s)


	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c[2:5]) //slice operator
    
	t1 := []string {"a", "b", "c"}
	t2 := []string {"a", "b", "c"}
	if slices.Equal(t1, t2) {
		fmt.Println("Both Slices are Equal")
	}
   

	two2D := make([][]int, 3)
	for i := range 3 {
		innerLen := i+1
		two2D[i] = make([]int, innerLen)
		for j:= range innerLen {
            two2D[i][j] = i + j
		}
	}

	fmt.Println(two2D)
}