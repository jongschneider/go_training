package main

import "fmt"

func main() {

  for i := 0; i < 5; i++ {
    if i ==3 {
      break
    }
    fmt.Println(i)
  }
  for j := 0; j < 11; j++ {
    if j == 7 {
      continue
    }
    fmt.Println(j)
  }
}
