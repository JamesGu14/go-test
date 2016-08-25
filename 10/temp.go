package main

import (
	"fmt"
	"strconv"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := &person{
		Name: "Joe",
		Age:  20,
	}
	i := 2
	fmt.Println("Before calling the function: ", strconv.Itoa(i), " ", a)
	Test(&i, a)
	fmt.Println("After calling the function: ", strconv.Itoa(i), " ", a)
	// a.Age = 18
	// a.Name = "Joe"
	fmt.Println(a)
	fmt.Println("========")
	A(a)
	fmt.Println(a)

	fmt.Println("========")
	t := teacher{
		Name:  "Teacher",
		Age:   50,
		human: human{Sex: 1},
	}
	s := student{
		Name:  "Student",
		Age:   18,
		human: human{Sex: 0},
	}
	s.Sex = 100
	fmt.Println(t, s)
}

func A(_person *person) {
	_person.Age = 28
	fmt.Println("A", _person)
}

func Test(i *int, _p *person) {
	_p.Name = "Denver"
	*i = *i + 8
	fmt.Println(*i, " ", _p)
}

/*
Go 中没有类，也没有继承
但是有 嵌入结构
*/
type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}
