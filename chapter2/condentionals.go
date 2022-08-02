package chapter2

import "fmt"

func CondentionalsDemo() {
	var s1, s2 = 45, 34
	if s1 > s2 {
		fmt.Print(s1, " sayısı ", s2, " den büyüktür")
	} else if s2 > s1 {
		fmt.Print(s2, " sayısı ", s1, " den büyüktür")
	} else {
		fmt.Print(s1, " sayısı ", s2, " sayısına eşittir")
	}
}
