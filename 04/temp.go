package main

import (
  "fmt"
)

const (

  B float64 = 1 << (iota * 10)
  KB
  MB
  GB
)

func main() {
  // fmt.Println(1 ^ 2)
  // fmt.Println(!true)

  // fmt.Println(1 << 10)

  // fmt.Println(6 & 11)

  fmt.Println(B)
  fmt.Println(KB)
  fmt.Println(MB)
  fmt.Println(GB)
}
