package chapter3

import "fmt"

func MapDemo() {
	var A map[string]string
	A = make(map[string]string)
	A["Table"] = "Masa"
	A["Pencil"] = "Kalem"
	fmt.Println(A)

	for k, _ := range A {
		value, state := A[k]
		if state {
			fmt.Println(k, " : ", value)
		} else {
			fmt.Println("Değeri mevcut değil")
		}

	}

}
