package chapter3

import "fmt"

func SliceDemo() {
	var A = []int{23, 34, 45, 56, 67}
	var B = A[2:4]
	var C = A[:3]

	fmt.Println(A, B, C)

	A = make([]int, 5)
	A = append(A, 34, 23)

	fmt.Println(A)

}
