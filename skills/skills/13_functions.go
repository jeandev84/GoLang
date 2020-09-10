package main

import "fmt"

func main () {
  var a = 3
  var b = 2

  a, b = summ(a, b)
  fmt.Println(a, b)
}


// функция принимает 2 параметра и возврашает 2 параметра
func summ (num_1 int, num_2 int) (int, int) {
    var res int
    res = num_1 + num_2
    return res, num_1 * num_2
}