package chapter3

import "fmt"

func ForRangeDemo() {
	var array = []string{"Ömer", "Faruk", "Betül"}
	for i, name := range array {
		fmt.Println(i, name)
	}

	var slice = make([]int, 5)
	slice = append(slice, 23, 34, 45)
	for i, item := range slice {
		fmt.Println(i, item)
	}
}
