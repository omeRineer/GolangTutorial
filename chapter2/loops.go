package chapter2

import "fmt"

func LoopDemo() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	var sayi1 = 10
	for sayi1 <= 20 {
		fmt.Print(sayi1)
		sayi1++
	}

}
