package main
import "fmt"

func bhaishabcopywale(a string){
      a = "hardik"
}


func bhaishabreferencevale(a *string){
	  *a = "hardik"
}

func main() {
	 a := "agarwal"
	 fmt.Println(a)
	 bhaishabcopywale(a)
     fmt.Println(a)
	 bhaishabreferencevale(&a)
	 fmt.Println(a)
	 fmt.Println(&a)
}