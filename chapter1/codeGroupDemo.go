package chapter1

import "fmt"

func CodeGroupDemo() {
	degisken := 35
	{
		degisken := 45
		fmt.Println(degisken)
	}
	fmt.Println(degisken)
}
