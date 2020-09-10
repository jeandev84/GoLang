package main

import "fmt"

func main () {
   var i = 1
   for i <= 10 {
      fmt.Println(i)
      i++
   }
}

$ go run project.go
1
2
3
4
5
6
7
8
9
10
