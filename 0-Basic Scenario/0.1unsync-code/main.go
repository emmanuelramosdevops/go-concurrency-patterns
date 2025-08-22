package main

import (
	"fmt"
	"time"
)

func write() {
	time.Sleep(time.Second * 2)
	fmt.Println("Write...")
}

func main() {
	go write()
	fmt.Println("Main...")
}
