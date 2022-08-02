package chapter2

import "fmt"

func DeferDemo() {
	defer fmt.Print("1.Satır")
	defer fmt.Print("2.Satır")
	defer fmt.Print("3.Satır")

	fmt.Print("4.Satır")
}
