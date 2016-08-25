package main

import (
	"fmt"
)

func main() {
	var m map[int]string = make(map[int]string)

	m[1] = "OK"
	m[2] = "OK2"
	delete(m, 1)

	fmt.Println(m[1] + " " + m[2])

	/*
			  迭代操作：
		    相当于 foreach
	*/

	sm := make([]map[int]string, 5)

	for i := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "OK"
		sm[i][2] = "OK2"
		fmt.Println(sm[i])
	}

	fmt.Println(sm)

	/*
				sort map
		    import (
		      "sort"
		    )
	*/

	/*
	   Test
	*/
	fmt.Println("===========")
	m1 := map[int]string{5: "a", 8: "b", 2: "c"}
	fmt.Println(m1)
	m2 := make(map[string]int, len(m1))

	//s1 := make([]string, 10)

	for k, v := range m1 {
		m2[v] = k
	}

	fmt.Println(m2)
}
