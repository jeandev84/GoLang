// Array non indexes

package main

import "fmt"

func main () {

   // var arr[number_of_elements_massiv] int

   var arr[3] int

   arr[0] = 45
   arr[1] = 97
   arr[2] = 76

   fmt.Println(arr[1])

   // Method compact for declaration array

   nums := [3]float64 {4.23, 5.23, 98.1}

   // like foreach in php
   for i, value := range nums {
      fmt.Println(value, i)
   }
}

