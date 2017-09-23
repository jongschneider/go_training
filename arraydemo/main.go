package main

import (
  "fmt"
  "math/rand"
)

func main() {
  fmt.Println("define arrays")
  var numbers[5] int
  var cities[5] string
  var matrix[3][3] int

  fmt.Println(">>>>>>>>>>insert array data")
  for i := 0; i < 5; i++ {
    numbers[i] = i
    cities[i] = fmt.Sprintf("City %d", i)
  }

  fmt.Println(">>>>>>>>>>insert matrix data")
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      matrix[i][j] = rand.Intn(100)
    }
  }

  fmt.Println(">>>>>>>>>>display array data")
  for i := 0; i < 5; i++ {
    fmt.Println(numbers[i])
    fmt.Println(cities[i])
  }

  fmt.Println(">>>>>>>>>>display matric data")
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      fmt.Println(matrix[i][j])
    }
  }
}
