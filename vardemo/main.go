package main

import "fmt"

func main()  {
  // declare variables
  var str string
  var n, m int
  var mn float32

  // assign values
  str = "Hello world"
  n = 10
  m = 50
  mn = 2.45

  fmt.Println("value of str = ", str)
  fmt.Println("value of n = ", n)
  fmt.Println("value of m = ", m)
  fmt.Println("value of mn = ", mn)

  // declare and assign
  var city string = "London"
  var x int = 100

  fmt.Println("value of city = ", city)
  fmt.Println("value of x = ", x)

  // declare variable w/o defining type
  val := 15
  fmt.Println("value of val = ", val)

  // define multiple values
  var (
    name string
    email string
    age int
  )
  name = "Jon"
  email = "J@test.com"
  age = 32

  fmt.Println("value of name = ", name)
  fmt.Println("value of email = ", email)
  fmt.Println("value of age = ", age)
}
