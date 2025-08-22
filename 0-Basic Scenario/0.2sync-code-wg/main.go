package main

import (
	"fmt"
	"sync"
	"time"
)

func write(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	fmt.Println("Write...")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go write(&wg)

	wg.Wait()
	fmt.Println("Main...")
}
