package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Golang - Concurrency")
	Counter(0, 10, 1)
	Counter(20, 30, 2)

	fmt.Printf("\nParallel Process:\n")
	// Run as parallel
	go Counter(0, 10, 1)
	Counter(20, 30, 2)

	fmt.Printf("\nWait group:\n")
	var wg sync.WaitGroup
	// wg.Add(1)
	// go CounterWaitGroup(1, 10, 100, &wg)
	// wg.Wait()

	// wg.Add(1)
	// go CounterWaitGroup(20, 30, 100, &wg)
	// wg.Wait()

	fmt.Printf("\nChannels: Send message from channel to another one, PingPong Channels\n")

	PingChannel := make(chan string)
	PongChannel := make(chan string)

	wg.Add(1)
	go SayHello(&wg, PingChannel, PongChannel)

	wg.Add(1)
	go SayWorld(&wg, PingChannel, PongChannel)
	wg.Wait()
}

func Counter(start int, end int, waitingTime int) {
	for i := start; i < end; i++ {
		fmt.Printf("%d, ", i)
		time.Sleep(time.Duration(waitingTime) * time.Millisecond)
	}
}
func CounterWaitGroup(start int, end int, waitingTime int, wg *sync.WaitGroup) {
	wg.Done()
	for i := start; i < end; i++ {
		fmt.Printf("%d, ", i)
		time.Sleep(time.Duration(waitingTime) * time.Millisecond)
	}
}

func SayHello(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Print("Hello ")
		pingChannel <- "World..."
		<-pongChannel
	}
	wg.Done()
}

func SayWorld(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
	for i := 0; i < 5; i++ {
		value := <-pingChannel
		fmt.Printf("%s\n", value)
		pongChannel <- "Pong"
	}
	wg.Done()
}
