package main

import (
	"fmt"
	"time"
	"sync"
)

func produce(queue chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		queue <- i
		time.Sleep(time.Millisecond * 200)
	}
	done <- true
}

func consume(queue chan int, done chan bool) {
    for {
      msg := <-queue
      fmt.Println(msg)
   }
}

func main() {
	queue := make(chan int, 10)
	done := make(chan bool)

	go produce(queue, done)
	go consume(queue, done)

	<-done
	fmt.Println("Task done")
}