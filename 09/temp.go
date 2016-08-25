package main

import (
	"fmt"
)

func main() {
	B(1, 3, 2, 5, 6, 8)

	// 匿名函数
	a := func() {
		fmt.Println("Func a")
	}
	a()

	// 闭包
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

	fmt.Println("=======")
	// fmt.Println("a")
	// defer fmt.Println("b")
	// defer fmt.Println("c")

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	fmt.Println("=======Panic")
	pA()
	pB()
	pC()
}

func A(a, b, c int) (d, e, f int) {
	d = a + 1
	e = b + 1
	f = c + 1
	return
}

func B(a ...int) {
	fmt.Println(a)
}

func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

func pA() {
	fmt.Println("Func pA")
}

func pB() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()

	panic("Panic in B")
}

func pC() {
	fmt.Println("Func C")
}
