package chapter3

import "fmt"

func ArrayDemo() {
	var A [3]int
	A[0] = 12
	A[1] = 11
	A[2] = 45
	fmt.Println(A)

	var B = []int{2, 4, 5, 1}
	fmt.Println(B)
}
