package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Current time:", time.Now())
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("OS/Arch:", runtime.GOOS, runtime.GOARCH)
	fmt.Println("CPUs available:", runtime.NumCPU())
	fmt.Println("Goroutines running:", runtime.NumGoroutine())
}
