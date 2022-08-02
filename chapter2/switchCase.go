package chapter2

import "fmt"

func SwitchCaseDemo1() {
	sayi := 34

	switch sayi {
	case 10:
		fmt.Print(sayi)
	case 20:
		fmt.Print(sayi)
	case 30:
		fmt.Print(sayi)

	default:
		fmt.Print("Bu sayı mevcut değil")
	}

}

func SwitchCaseDemo2() {
	var A, B = 34, 56
	switch {
	case A > B:
		fmt.Print("A, B'den büyüktür")
	case A < B:
		fmt.Print("B, A'dan büyüktür")
	default:
		fmt.Print("A ve B birbirine eşittir")
	}
}

func SwitchCaseDemo3() {
	var A = 56

	switch {
	case A < 6:
		fmt.Print("A, 6'dan küçüktür")
		fallthrough
	case A < 7:
		fmt.Print("A, 7'den küçüktür")
	}
}
