package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	count := make(chan int)

	go PrintDetails("A", count)
	go PrintDetails("B", count)
	count <- 1
	wg.Wait()
}

func PrintDetails(s string, count chan int) {
	defer wg.Done()

	for {
		r,ok := <- count

		if !ok {
			return
		}

		fmt.Println("from: ", s, "val: ", r)

		r++
		if r == 10 {
			close(count)
			// return
		}

		count <- r

	}

}
