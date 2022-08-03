package chapter4

import (
	"fmt"
	"time"
)

func GoRoutineDemo() {
	fmt.Println("Main başladı")
	go asenkronMetod()
	fmt.Println("Main bitti")
	time.Sleep(time.Second * 2)
}

func asenkronMetod() {
	fmt.Println("Asenkron metod çalıştı")
}
