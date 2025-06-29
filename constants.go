package main

import(
 "fmt"
 "math"
)

const a string = "John the don!"

func main(){
   fmt.Println(a)
   
   const a = 100000
   const b = a/ 3e2
   fmt.Println(b)
   const c = 1.57
   fmt.Println(math.Sin(c)) //no need to explicitly do math.Sin(float64(c)), as go automatically detects it

}