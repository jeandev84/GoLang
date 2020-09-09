
// Array indexes (Card or Maps)


package main

import "fmt"

func main () {

   // webSites := make(map[typeKeyData]typeValueData)
   // webSites := make(map[string]int)
   webSites := make(map[string]float64)

   webSites["itProger"] = 0.8
   webSites["yandex"] = 0.99

   fmt.Println(webSites["itProger"])
}

