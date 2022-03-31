package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	resultsChan := make(chan int)
	floatChan := make(chan float64)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			resultsChan <- i * 100
			time.Sleep(time.Millisecond * time.Duration(i*1000))
			fmt.Printf("Sleep time over for: %d\n", i)
			floatChan <- float64(i) * .5
		}()
	}

	go func() {
		defer fmt.Printf("Exiting result collector")
		moreResults := true
		for moreResults {
			select {
			case num, more := <-resultsChan:
				if more {
					fmt.Printf("Got result: %d\n", num)
				} else {
					moreResults = false
					fmt.Printf("Exiting\n")
					break
				}
			case float, more := <-floatChan:
				if more {
					fmt.Printf("got float: %f\n", float)
				} else {
					moreResults = false
					fmt.Printf("Exiting from Float\n")
					break
				}
			}
		}
	}()

	wg.Wait()
	close(floatChan)
	close(resultsChan)
	time.Sleep(time.Millisecond * 1000)
}
