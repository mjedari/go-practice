package app

import (
	"sync"
)

// receives slice of channels and iterate over them
// for each channel it is created a new go routine and do the process inside it
// then writes the output to the destination channel which is unique

func Funnel(sources ...<-chan string) <-chan string {
	dest := make(chan string) // make sure to close the channel you've made
	var wg sync.WaitGroup

	wg.Add(len(sources))

	for _, ch := range sources {
		go func(c <-chan string) {
			defer wg.Done()

			for i := range c {
				dest <- i
			}
		}(ch)
	}

	// Note: If you do not use another go routine, we would have deadlock
	// the main goroutine waits to all source channels Done
	// but the problem is source channels never Done because the wait to receive data from dest channel
	// and because this function never returns dest channel (until all goroutines are Done),
	// there is no receiver in the program and go runtime decide that the situation is deadlock

	// Start a goroutine to close dest after all sources close
	go func() { // <---- Note
		wg.Wait()
		defer close(dest)
	}()

	return dest
}
