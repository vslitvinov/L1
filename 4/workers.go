package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
//

func GenWorker(n int, wg *sync.WaitGroup) (chan  interface{}) {
	ch := make(chan interface{},n)

	fmt.Println("worker gen")
	for i:=0;i<n;i++ {
		wg.Add(1)
		go func(id int ){
			defer wg.Done()
			fmt.Printf("Worker id %v starting \n", id)
			for {
				select {
				case data, ok := <- ch:
					if ok {
						fmt.Printf("Worker id :%b, read data from the channel : %d \n",id,data)
					} else {
						fmt.Printf("Worker id: %v stop \n",id)
						return 
					}
				}
			}
		}(i)
	}
	return ch

}


func main(){

	// worker count
	var n int	
	var err error
	if len(os.Args) == 2 {
		n, err = strconv.Atoi(os.Args[1])
		if err != nil || n == 0 {
			fmt.Println("You need to write a number starting from 1")
			return
		}
	} else {
		fmt.Println("You need to specify 1 argument - amount of workers")
		return
	}

	i:=0
	wg := sync.WaitGroup{}
	fmt.Println(n)
	ch := GenWorker(n,&wg)
	

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	for {
		i++
		select {
		case signal := <- sigs:
			close(ch)
			wg.Wait()
			fmt.Printf("Close app signal %s \n",signal)
			return 
		default:
			ch <- i
		}

		time.Sleep(time.Millisecond * 100)
	}




}


