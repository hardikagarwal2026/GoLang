package main

import "fmt"

func main() {
   //the most basic for loop
   i:= 1
   for i<= 3 {
       fmt.Println(i)
       i++
   }
   
   // initial/condition/after
   for i:=1; i<=3; i++ {
      fmt.Println(i)
   }
   
   // Do this N Times
   for i := range 3 {
     fmt.Println(i)
   }
   
   // Without Any Condition
   for{
      fmt.Println("hardik the boss!")
      break;
   }
   
   //To check Even or not
     for n := 0; n < 100; n++ {
        if n%2 == 0 {
            fmt.Printf("%d - Number is Even\n", n)
        }
    }

}