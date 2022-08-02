package chapter3

import "fmt"

type Dynamic interface{}

func DynamicType() {
	var variable Dynamic = 0
	fmt.Println(variable)
	variable = "asd"
	fmt.Println(variable)
	variable = 12
	fmt.Println(variable)
}
