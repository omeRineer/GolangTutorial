package chapter3

import "fmt"

type Customer struct {
	id      int
	name    string
	surname string
}

func StructDemo() {
	var customer1 = Customer{12, "Ömer Faruk", "Sandıkçı"}
	fmt.Println(customer1)

	var customer2 = Customer{id: 21, name: "Hatice Betül", surname: "Sandıkçı"}
	customer2.id = 34
	customer2.name = "asd"
	customer2.surname = "asd"
	fmt.Println(customer2)
}

func AnonymousStruct() {
	//Anonymous Struct
	var people = struct {
		ad, soyad string
	}{ad: "Zehra", soyad: "Sandıkçı"}
	fmt.Println(people)
}

func (c Customer) StructFunction() {
	fmt.Print(c.name, c.surname)
}
