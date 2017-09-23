package main

import "fmt"

func main() {
  var a int = 4

  // inc
  fmt.Printf("a = %d \n", a)
  a = a + 1
  fmt.Printf("a + 1 = %d \n", a)
  a++
  fmt.Printf("a++ = %d \n", a)

  // dec
  fmt.Printf("a = %d \n", a)
  a = a - 1
  fmt.Printf("a - 1 = %d \n", a)
  a--
  fmt.Printf("a-- = %d \n", a)
}
