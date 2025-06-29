package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
   const s = "@@@"
   for i:=0;i<len(s);i++{    //iterating over each byte if rune, a rune can have multiple bytes
	fmt.Printf("%x\n", s[i]) //gives lower hexadecimal 
   }
   fmt.Println(utf8.RuneCountInString(s))


   // to decode each rune along with its offset
   for idx, runevalue := range s {
	fmt.Printf("%#U starts at %d\n", runevalue, idx)
   }

}