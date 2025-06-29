package main

import (
	"fmt"
	"maps"
)

func main() {
   m := make(map[string]int)
   m["k1"] = 1
   m["k2"] = 2
   fmt.Println(m["k1"])
   fmt.Println(len(m))
   delete(m, "k1")
   _, prs := m["k2"]
   fmt.Println(prs)
   clear(m)
   m1 := map[int]int {1:1, 2:2}
   m2 := map[int]int {1:1, 2:2}
   if maps.Equal(m1, m2) {
	fmt.Println("Both maps are Equal")
   }





}