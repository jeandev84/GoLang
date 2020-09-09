package main

import "fmt"

func main () {
   var i = 10
   for i > 0 {
      fmt.Println(i)
      i--
   }
}

// FOR inverse (считать в обратном порядке)
// $ go run project.go
10
9
8
7
6
5
4
3
2
1
