package main

import (
  "fmt"
  "strconv"
)

func main() {
  a := 1
  var p *int = &a
  fmt.Println(p)
  fmt.Println(*p)

  if b := 1; b < 2 {

  }

  /*
    For loop
  */
  
  for f := 1; f <= 3; f++ {
    fmt.Println("f value is " + strconv.Itoa(f))

  }

  /*
    break 
    continue
    goto
  */
  for m := 0; m < 10; m++ {
LABEL:
    for n := 0; n < 10; n++ {

      for i := 0; i < 10; i++ {
        if (i == 1) {
          break LABEL
        } else {
          fmt.Println("m: " + strconv.Itoa(m) + " n: " + strconv.Itoa(n) + " i: " + strconv.Itoa(i))
        }
      }
    }
  }
}
