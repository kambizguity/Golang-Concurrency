package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// fmt.Println("Golang - Concurrency")
	// Counter(0, 10, 1)
	// Counter(20, 30, 2)

	// fmt.Printf("\nParallel Process:\n")
	// // Run as parallel
	// go Counter(0, 10, 1)
	// Counter(20, 30, 2)

	// fmt.Printf("\nWait group:\n")
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go CounterWaitGroup(1, 10, 100, &wg)
	// wg.Wait()

	// wg.Add(1)
	// go CounterWaitGroup(20, 30, 100, &wg)
	// wg.Wait()

	// fmt.Print("Channels: Send message from channel to another one, PingPong Channels\n")

	// PingChannel := make(chan string)
	// PongChannel := make(chan string)

	// wg.Add(1)
	// go SayHello(&wg, PingChannel, PongChannel)

	// wg.Add(1)
	// go SayWorld(&wg, PingChannel, PongChannel)
	// wg.Wait()

	// //Delay to complete tasks
	// time.Sleep(4 * time.Second)

	// //Send and recieve data are blocking in channels so
	// //Read data from any port is available with channel but i don't want to block any port after recieve the data
	// //I need to recieve data from port and proccess it as well as possible, So i need to put data into buffer channel.

	// fmt.Println("Buffer Channel")
	// logChannel := make(chan string, 10)
	// go printLog(logChannel)

	// for i := 0; i < 100; i++ {
	// 	fmt.Println("Push into channel", i)
	// 	logChannel <- fmt.Sprintf("Counter is %d", i)
	// }

	// fmt.Println("Channel Closed")

	//Range & Close Channel

	count := 10
	newChannel := make(chan string, count)

	go printLog(newChannel)

	for i := 0; i < 5; i++ {
		content := fmt.Sprintf("Value: %v", i)
		newChannel <- content
		fmt.Println("Push: ", i)
	}

	time.Sleep(6 * time.Second)
	close(newChannel)
}

func printLog(ch chan string) {
	for value, ok := <-ch; ok; {
		fmt.Println("Pop Channel:", value)
	}
	fmt.Println("Channel Closed!")
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
