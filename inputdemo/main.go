package main

import (
  "fmt"
  "math"
)
func main() {
  fmt.Println("Circle Area Calculation")
  fmt.Println("Enter a Radius Value: ")
  var radius float64
  fmt.Scanf("%f", &radius)

  area := math.Pi * math.Pow(radius,2)
  fmt.Printf("Area of a circle with a radius of %.2f = %.2f \n", radius, area)
}
