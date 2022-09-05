package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

// Work - imitate some work for goroutine
func Work() int {
	return 1
}

func main() {

	ch := make(chan int)



	//  context
	// ctx, cancel := context.WithCancel(context.Background())

	// go func(ctx context.Context) {
	// 	for {
	// 		select {
	// 		case ch <- Work():
	// 			fmt.Println("Sending data ...")
	// 		case <-ctx.Done():
	// 			close(ch)
	// 			return
	// 		}
	// 		time.Sleep(time.Millisecond * 100)
	// 	}
	// }(ctx)

	// go func() {
	// 	time.Sleep(time.Second * 1)
	// 	cancel()
	// }()

	// for data := range ch {
		// fmt.Printf("Read data ... %d\n", data)
	// }

	//  channel
	done := make(chan struct{})

	go func() {
		for {
			select {
			case ch <- Work():
				fmt.Println("Sending data ...")
			case <-done:
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		time.Sleep(time.Second * 1)
		done <- struct{}{}
	}()

	for data := range ch {
		fmt.Printf("Read data ... %d\n", data)
	}

}
