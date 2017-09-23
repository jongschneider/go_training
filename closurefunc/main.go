package main

import "fmt"

func main() {
  closure_func("testing")
}

func closure_func(name string)  {

  hoo := func (a, b int)  {
    result := a*b
    fmt.Printf("hoo() = %d \n", result)
  }

  joo := func (a, b int) int {
    return a*b + a
  }
  fmt.Printf("closure_func(%s) was called \n", name)
  hoo(2,3)
  val := joo(5, 8)
  fmt.Printf("val from joo()= %d \n", val)
}
