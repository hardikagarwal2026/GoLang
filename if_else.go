package main

import "fmt"

func main() {
  //basic
  if 2 % 2 == 0 {
	fmt.Println("Even Number!")
  }else {
	fmt.Println("Number is Odd")
  }
  

  str1 := "hardik"
  str2 := "Janhvi"
  if str1 == str2 || 2 % 2==0 {
	fmt.Println("Great")
  }  

  // variable declare in this statement are available in the current and all the subseqquent branches
  if num := 9; num < 0 {
	fmt.Println("Number is Negative")
  }else if num<10{
	fmt.Println("Numnber has one digit!")
  }else{
	fmt.Println("Number has two digits")
  }


}