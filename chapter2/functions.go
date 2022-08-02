package chapter2

import "fmt"

func Write(word string) {
	fmt.Print(word)
}

func ReturnValue(x, y, z int) int {
	return x + y + z
}

func ReturnCollection() (a, b, c int) {
	a = 45
	b = 56

	c = a + b
	return
}

//Variadic Functions
func VariadicFunc(numbers ...int) int {
	var total = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

//Clasure Functions
func ClasureFunction(x, y int) int {
	return x + y
}

func ClasureFunction2() func(word string) {
	return func(word string) {
		fmt.Print(word)
	}
}

//Recursive Functions
func Factorial(a uint) uint {
	if a == 0 {
		return 1
	}
	return a * Factorial(a-1)
}

//Anonymous Functions
func AnonymousFunction() {
	someWrite := "Hello World!"
	func(someWord string) {
		fmt.Print(someWord)
	}(someWrite)
}

//Empty Identifier
func Empty_Function() (int, int) {
	val1 := 56
	val2 := 67

	return val1, val2
}
