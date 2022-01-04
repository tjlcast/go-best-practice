package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ChannelMessage struct {
	name string
	wg *sync.WaitGroup
	ch chan int
}

func player(msg *ChannelMessage) {
	defer msg.wg.Done()
	fmt.Println(msg.name, " play.")

	for {
		ball, ok := <-msg.ch
		if !ok {
			fmt.Println("channel is closed.")
			return
		}
		fmt.Println(msg.name, " receive the ball: ", ball)

		rn := rand.Intn(100)
		if rn % 10 == 0 {
			//close(msg.ch)
			fmt.Println(msg.name, " close the channel.")
		}

		ball++
		msg.ch<- ball
	}
}

var counter = 0
var wg = sync.WaitGroup{}
var mtx = sync.Mutex{}
func unSafeCounter() {
	defer wg.Done()
	for i:=0; i<1000; i++ {
		mtx.Lock()
		counter++
		mtx.Unlock()
	}
	fmt.Println("Done.")
}

func main() {
	// ch := make(chan int, 0)
	// wg := &sync.WaitGroup{}
	// size := 2
	// wg.Add(2)

	// for i:=0; i<size; i++ {
	// 	msg := &ChannelMessage{
	// 		"tjl"+strconv.Itoa(i),
	// 		wg,
	// 		ch,
	// 	}
	// 	go player(msg)
	// }

	// ball := 0
	// ch<- ball

	// wg.Wait()

	size := 3
	wg.Add(size)
	for i:=0; i<size; i++ {
		go unSafeCounter()
	}
	wg.Wait()
	fmt.Println(counter)
}