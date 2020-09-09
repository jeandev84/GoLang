package main

import "fmt"

func main () {
  var a = 29
  var b = 1
  var r int

  r = summ(a, b)
  fmt.Println(r)
}


// функция принимает 2 параметра и возврашает один параметр
func summ (num_1 int, num_2 int) int {
    var res int
    res = num_1 + num_2
    return res
}