package main

import (
	"fmt"
	"math"
)


type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

func(r rect) area() float64 {
	return r.height* r.width
}

func(r rect) perim() float64 {
	return 2*(r.height+r.width)
}

func(c circle) perim() float64 {
	return 2*math.Pi*c.radius
}

func(c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}


func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g,g.perim())
}

func detectcircle(g geometry){
	_, ok := g.(circle)
	if ok {
		fmt.Println("yes it is a circle")
	}else {
		fmt.Println("Not a circle")
	}
}

func main() {
   c:= circle{radius: 7.0}
   r:= rect{height: 5.0, width: 5.0}
   fmt.Println(c.area())
   fmt.Println(c.perim())
   fmt.Println(r.area())
   fmt.Println(r.perim())
   
   measure(c)
   measure(r)
   detectcircle(c)
   detectcircle(r)
}