package main

import "fmt"

func main () {
  var x = 0
  pointer(x)
  fmt.Println(x)
}


func pointer (x int) {
    x = 2
}

$ go run project.go
0

==================================================

package main

import "fmt"

func main () {
  var x = 0
  pointer(&x)
  fmt.Println(x)
}


// ссылка указывается в следующем образом "*"
// will be change 0 to 2 (in our case)
func pointer (x *int) {
    *x = 2
}

$ go run project.go
2