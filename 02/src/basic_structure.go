package main

import (
  "fmt"
  "math"
  "strconv"
)

const PI = 3.14

var (
  name = "gopher"
  age = "24"
)

type newType int

type gopher struct {}

type golang interface {}

var a, b, c, d = 1, 2, 3, 4

func main() {
  
  fmt.Println(math.MinInt8)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)

  // ======
  var x int = 65
  b := strconv.Itoa(x)
  fmt.Println(b)
}
