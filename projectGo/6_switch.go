// в GoLang при использовании switch не нужно дописать "BREAK";

package main

import "fmt"

func main () {
  var age = 10
  switch age {
      case 5: fmt.Println("Вам 5 лет")
      case 15: fmt.Println("Вам 15 лет")
      case 10: fmt.Println("Вам 10 лет")
      default: fmt.Println("Нам неизвестно сколько вам лет")
  }
}

$ go run project.go