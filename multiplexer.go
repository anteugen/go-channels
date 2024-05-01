package main

import (
	"fmt"
	"time"
)

func source1(ch chan string) {
	for {
		ch <- "From source 1"
		time.Sleep(time.Millisecond * 200)
	}
}

func source2(ch chan string) {
	for {
		ch <- "From source 2"
		time.Sleep(time.Millisecond * 200)
	}
}

func multiplexer(output, ch1, ch2 chan string) {
	for {
		select {
		case msg := <-ch1:
			output <- msg
		case msg := <-ch2:
			output <-msg
		}
	}
}

func main() {
	output := make(chan string, 10)

	ch1 := make(chan string)
	ch2 := make(chan string)

	go source1(ch1)
	go source2(ch2)
	go multiplexer(output, ch1, ch2)

	for msg := range output {
		fmt.Println(msg)
	}


}