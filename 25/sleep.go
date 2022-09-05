package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

// Sleep - pause a thread on time t
func Sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	start := time.Now()
	fmt.Println("Program started")

	fmt.Println("Program will fall asleep for 3 seconds")
	Sleep(time.Second * 3)

	fmt.Println("Program fineshed with time: ", time.Since(start))
}
