package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{Name: "AAA"}
	a.Print()
}

func (a *A) Print() {
	fmt.Println("A print function: ", a.Name)
}
