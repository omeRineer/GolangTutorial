package chapter1

import "fmt"

func ConvertTypeDemo() {
	var sayi = 45
	var sayi2 = int(sayi)
	var sayi3 = float32(sayi2)

	fmt.Println(sayi3)
}
