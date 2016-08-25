package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	c := make(chan bool)
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("Go Go Go!!!")
// 		c <- true

// 		fmt.Println("what ever")
// 	}()

// 	for v := range c {

// 		fmt.Println(v)
// 		if v {
// 			close(c)
// 		}
// 	}
// }

// func main() {
// 	c := make(chan bool, 1)
// 	go func() {
// 		fmt.Println("Go Go Go!!!")
// 		c <- true
// 	}()
// 	<-c
// }

func main() {
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}
