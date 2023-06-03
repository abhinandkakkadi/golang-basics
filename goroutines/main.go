package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	count := make(chan int)
	// here same  function is called with different arguments independently
	go PrintDetails("A", count)
	go PrintDetails("B", count)
	count <- 1
	wg.Wait()
}

func PrintDetails(s string, count chan int) {
	defer wg.Done()

	for {
		// the first goroutine to reach this statement will receive the values and go on with that
		r, ok := <-count
		// if the channel is closed by other go routine this should return
		if !ok {
			return
		}

		fmt.Println("from: ", s, "val: ", r)

		r++
		if r == 10 {
			// after passing around and iterating values, close the channel when count = 10
			close(count)
			return
		}
		// send value to channel so that other goroutine which is waiting for recieving can get the value
		count <- r

	}

}
