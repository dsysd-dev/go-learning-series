package main

import (
	"sync"
)

func main() {

	oc, ec := make(chan int), make(chan int)

	go func() {
		for {
			select {
			case <-oc:

			case <-ec:
			}
		}
	}()
}

func main2() {

	var wg sync.WaitGroup
	wg.Add(2)

	signal_odd, signal_even := make(chan int), make(chan int)

	go func() {
		// this one will print odd numbers
		for i := 1; i < 10; i += 2 {
			<-signal_odd
			println("\t", i)
			signal_even <- 1
		}
		wg.Done()
	}()

	go func() {
		// this one will print even numbers
		for i := 2; i < 10; i += 2 {
			<-signal_even
			println("\t", i)
			signal_odd <- 1
		}
		<-signal_even
		wg.Done()
	}()
	signal_odd <- 1
	wg.Wait()
}

/*

there are 2 go routines
print odd numbers and even numbers
print them in sequence


1 3 5 7 9
2 4 6 8


1 2 3 4 5 6 7 8 9

*/
