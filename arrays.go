package main

import "fmt"

func main() {
   var a[5] int
   a[0] = 1
   fmt.Println(a)

   b := [...]int {1,2,3} //compiler automaticlly determines length
   fmt.Println(b)

   twoD := [2][3]int {
	{1,2,3},
	{4,5,6},
   }
   fmt.Println(twoD)

   var two2D [2][3] int
   for i := range 2 {
	for j:= range 3 {
		two2D[i][j] = i+ j
	}
   }
   fmt.Println(two2D)


  //If you specify the index with :, the elements in between will be zeroed.
  //{1,0,0,4,5} at 3rd poition 4
  c := [...]int{1, 3: 4, 5}
  fmt.Println(c)

}