package main

import (
	"fmt"
	"runtime"
)

func getAverages(numbers []int, ch chan float64) {
	var sum = 0
	for _, n := range numbers {
		sum += n
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]

	for _, n := range numbers {
		if max < n {
			max = n
		}
	}
	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 5, 7, 2, 8, 9, 3, 4, 9, 5, 2, 13}
	fmt.Println("numbers ", numbers)

	var ch1 = make(chan float64)
	go getAverages(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t : %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t : %d \n", max)
		}
	}
}
