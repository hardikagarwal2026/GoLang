package main

import (
	"fmt"
	"time"
)


func whatAmI(i interface{}) {
   switch i.(type) {
   case bool:
      fmt.Println("I am boolean")
   case int:
      fmt.Println("I am Integer")
   default:
      fmt.Println("I dont know my type")
   }
}
  


func main(){
   i:= 1
    switch i {
       case 1:
          fmt.Println("one")
       case 2:
          fmt.Println("two")
    }

   j := time.Now()
   switch {
      case j.Hour() < 12:
         fmt.Println("Morining")
      default:
         fmt.Println("Afternoon") 
   }

   whatAmI("HArdik")
   whatAmI(true)
   whatAmI(7)




}