package main

import "fmt"

func main() {
  var x int

  x = 10

  fmt.Println(x)
  fmt.Println(&x)

  var num *int
  val := new(int)

  num = new(int)
  *num = x

  val = &x

  fmt.Println("== Pointer Numbers ==")
  fmt.Println(*num) // 10
  fmt.Println(num)  // address
  fmt.Println("== Pointer Values ==")
  fmt.Println(*val) // 10
  fmt.Println(val)  // address
}
