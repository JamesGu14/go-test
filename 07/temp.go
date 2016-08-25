package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n'}

	sa := a[2:5]

	sb := sa[1:]

	fmt.Println(strconv.Itoa(len(a)))
	fmt.Println(strconv.Itoa(len(sa)) + " " + strconv.Itoa(cap(sa)))
	fmt.Println(strconv.Itoa(len(sb)) + " " + strconv.Itoa(cap(sb)))
	fmt.Println(string(sa))
	fmt.Println(string(sb))

	fmt.Println("======================")
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
}
