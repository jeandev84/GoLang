package main

import "fmt"

func main () {

  // ------------- FLOAT ------------
  var num float64 = 4.3254357

  // %f   - означает что мы выводим Float
  // fmt.Printf("%f \n", num)

  // %.2f - означает что мы хотим вывести 2 симбола после запитой
  // fmt.Printf("%.2f \n", num)

  // %T   - узнать тип переменной
  fmt.Printf("%T \n", num)


  // -------- BOOLEAN --------
  var isDone bool = true

  // %t - используем маленькую букву "t" чтобы вывести true / false
  fmt.Printf("%t \n", isDone)
}