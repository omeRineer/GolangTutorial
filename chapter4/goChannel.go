package chapter4

import (
	"fmt"
	"time"
)

func GoChannelDemo() {

	var kanal = make(chan string)

	fmt.Println("Fonksiyon başladı")

	go func() {
		time.Sleep(time.Second * 2)
		kanal <- "ds"
		fmt.Println("Asenkron metod çalıştı")
	}()
	<-kanal

	fmt.Println("Fonksiyon bitti")
}
