package main

import (
	"fmt"
	"time"
)

func write(c chan struct{}) {
	time.Sleep(time.Second * 2)
	fmt.Println("Write...")
	c <- struct{}{}
}

func main() {
	c := make(chan struct{})
	go write(c)

	<-c
	fmt.Println("Main...")
}
