package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func doPrint(wg *sync.WaitGroup, message string) {
// 	defer wg.Done()
// 	fmt.Println(message)
// }

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counter) Value() (x int) {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var mtx sync.Mutex
	var meter counter

	// for i := 0; i < 5; i++ {
	// 	data := fmt.Sprintf("data %d", i)

	// 	wg.Add(1)
	// 	go doPrint(&wg, data)
	// }
	// wg.Wait()

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for i := 0; i < 1000; i++ {

				mtx.Lock()
				meter.Add(i)
				mtx.Unlock()

			}

			wg.Done()

		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())

}
