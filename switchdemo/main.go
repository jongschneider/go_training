package main

import (
  "fmt"
)

func main() {
  var a int = 2

  switch a {
  case 0:
    fmt.Println("a = 0")
  case 1:
    fmt.Println("a = 1")
  case 2:
    fmt.Println("a = 2")
  default:
    fmt.Println("a != 0, 1 or, 2")
  }
}
