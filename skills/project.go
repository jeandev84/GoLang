package main

import "fmt"

func main () {

   bob := Cats {"Bob", 7, 0.87}
   fmt.Println("Bot function is", bob.test())
}


// Object
type Cats struct {
   name string
   age int
   happiness float64
}


// можно передать параметры функции test()
// test() возвращает float64
// float64(cat.age) : convert value age in to "float"
func (cat *Cats) test () float64 {
    return float64(cat.age) * cat.happiness
}
