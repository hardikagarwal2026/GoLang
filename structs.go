package main

import "fmt"

type Person struct {
	name string
	age int
}

func newperson(name string) *Person {
	p := Person{name: name}
	p.age = 40
	return &p
}


func main() {
   fmt.Println(Person{"Hardik", 21})
   fmt.Println(Person{name: "Hardik", age: 22})
   fmt.Println(Person{name: "Janhvi"})

   s1 := Person{name: "Janhvi", age: 20}
   fmt.Println(s1)

   fmt.Println(newperson("Arpit"))

   s2 := &s1 // pointing to the same address as of s1, so if any change in s2 happen happen will make changes in in s1 also  
   s2.age = 21
   fmt.Println(s2)

   //creating object dynamically using new keyword
   a := new(Person)
   a.name = "Sahil"
   a.age = 23
   fmt.Println(a)


   //creating an anomymous object
   dog := struct{
	name string
	age int
   }{
	"Pilla",
	 20,
   }
   fmt.Println(dog)
}