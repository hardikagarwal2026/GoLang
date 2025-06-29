package main

import "fmt"

type rect struct{
	height, width int
}

//value receiver
func (r rect) area() int {
	return r.height*r.width
}

//pointer receiver
func (r *rect) perim() int{
	return 2*(r.height + r.width)
}


func main() {
   r1 := rect{height: 4, width: 5}
   fmt.Println(r1.area())
   fmt.Println(r1.perim())

   r2 := & r1
   fmt.Println(r2.area())
   fmt.Println(r2.perim())
}