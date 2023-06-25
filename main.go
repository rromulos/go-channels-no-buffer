package main

import (
	"fmt"
	"time"
)

// done chan<- bool = can only write in the channel
// done <- chan bool = can only read the channel

func numbers(done chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}

	done <- true
}

func letters(done chan<- bool) {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}

	done <- true
}

func main() {
	channelNumber := make(chan bool)
	channelLetters := make(chan bool)

	go numbers(channelNumber)
	go letters(channelLetters)

	// waiting for
	<-channelNumber
	<-channelLetters

	fmt.Println("End of execution !")
}
