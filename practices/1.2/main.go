package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)

	go func() {
		ticker := time.NewTicker(1 * time.Second)

		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt ...")
			default:
				fmt.Printf("sending message %d\n", <-messages)

			}
		}
	}()

	for i := 0; i < 20; i++ {
		messages <- i
		time.Sleep(1 * time.Second)
	}

	close(done)
	fmt.Println("Main process exit")
}
