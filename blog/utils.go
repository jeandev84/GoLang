package main


import (
  "crypto/rand"
  "fmt"
)


// Метод для генерации массив bytes
func GenerateId() string {

   // сгенируем наш random числа 16 bytes
   b := make([]byte, 16)

   // считаем наш random
   rand.Read(b)

   // выводим
   return fmt.Sprintf("%x", b)
}