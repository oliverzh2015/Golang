package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("thread 1 is completed")
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("thread 2 is completed")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all thread are completed.")
}
