package main

import (
	"fmt"
	"runtime"
)

func main() {
    ch := make(chan int)

    go func() {
        val := <-ch
        fmt.Println("We received a value:", val)
    }()

    fmt.Println("Goroutines:", runtime.NumGoroutine())
}