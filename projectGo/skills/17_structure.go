package main

import "fmt"

func main () {

   bob := Cats {"Bob", 7, 0.87}
   fmt.Println("Bot age is", bob.age)
}


// Object
type Cats struct {
   name string
   age int
   happiness float64
}

$ go run project.go
