package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello World")
}

func main() {
	u := User{1, "OK", 12}
	Info(&u)

	fmt.Println("============")
	m := Manager{User: User{1, "OK", 18}, title: "Director"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	fmt.Println("============")
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)

	fmt.Println(x)

}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("This is not a struct, it could be a pointer")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%10s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%10s: %v\n", m.Name, m.Type)
	}
}
