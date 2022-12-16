package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, mtx *sync.RWMutex) {
	mtx.Lock()
	msg = s
	mtx.Unlock()
}

func printMessage(mtx *sync.RWMutex) {
	mtx.RLock()
	fmt.Println(msg)
	mtx.RUnlock()
}

func concurrentPrintMessage(wg *sync.WaitGroup, mtx *sync.RWMutex, messages ...string) {
	wg.Add(len(messages))
	for _, m := range messages {
		go func(m string, mtx *sync.RWMutex) {
			updateMessage(m, mtx)
			printMessage(mtx)
			defer wg.Done()
		}(m, mtx)
	}
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"
	var wg sync.WaitGroup
	var mtx sync.RWMutex
	concurrentPrintMessage(&wg, &mtx, "Hello, Cosmos!", "Hello, Universe!")
	wg.Wait()
}
