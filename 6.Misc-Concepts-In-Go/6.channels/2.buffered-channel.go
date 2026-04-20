package main

import (
	"fmt"
	"time"
)

func emailSenderQueue(emailChan chan string, processDoneChan chan bool){
	defer func() {
		processDoneChan <- true
	}()

	for email := range emailChan{
		fmt.Println("sending email to ", email)
		time.Sleep(time.Second)
	}
}

func main() {
	// ch := make(chan string, 2)

	// ch <- "Hello"
	// ch <- "World"

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	emailChan := make(chan string, 100) // size of the buffer is 100
	processChan := make(chan bool)

	go emailSenderQueue(emailChan, processChan)

	fmt.Println("Channels passed to queue...")

	for idx:=1; idx < 100; i++{
		emailChan <- fmt.Sprintf("%d@gmail.com", idx)
	}

	fmt.Println("done sending")

	// close the channel
	// close(emailChan)
	<-processChan


}