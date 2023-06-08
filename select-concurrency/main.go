package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// in select statement *case* which is ready to execute is selected. So for that reason blocking of first chanel does not happens,
	// whichever one is ready first - in this case whichever receives values from the channel will be available for execution.
	select {
	case v1 := <-ch1:
		fmt.Println("the first got executed because the first go routine sleep less is ", v1)
	case v2 := <-ch2:
		fmt.Println("the second got executed because the second go routine sleep less (sleep less work hard) which is ", v2)
	}
}
