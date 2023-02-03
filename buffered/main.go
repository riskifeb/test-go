package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2)

	go func() {
		for {
			i := <-messages
			fmt.Println("recieve data ", i)
		}
	}()

	for i := 1; i < 10; i++ {
		fmt.Println("Send data ", i)
		messages <- i
	}
}
