package chapter4

import (
	"fmt"
	"sync"
)

func WaitGroupDemo() {

	fmt.Println("Main is run")

	var wg sync.WaitGroup
	wg.Add(2)

	go FirstFunction(&wg)
	go SecondFunction(&wg)
	wg.Wait()

	fmt.Println("Main is end ")
}

func FirstFunction(wg1 *sync.WaitGroup) {
	fmt.Println("First function is run")
	wg1.Done()
}

func SecondFunction(wg2 *sync.WaitGroup) {
	fmt.Println("Second function is run")
	wg2.Done()
}
