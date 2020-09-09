package main

import "fmt"

func main () {

   // откладывание, мы вызываем сначало функцию two() а потом one()
   // но при выполнении код one () выполниться и потом two()
   // defer это как "за"
   defer two ()
   one()
}


func one () {
   fmt.Println("1")
}

func one () {
   fmt.Println("2")
}
