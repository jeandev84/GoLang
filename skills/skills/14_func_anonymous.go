package main

import "fmt"

func main () {
   var num = 3

   // Function anonymous
   multiple := func() int {
       num *= 2
       return num
   }

   fmt.Println(multiple())
}

