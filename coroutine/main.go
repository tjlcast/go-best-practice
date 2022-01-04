package main

import (
	"fmt"
	"sync"
	"time"
)

type TaskDesc struct {
	wg *sync.WaitGroup
	no int
	msg string
}

func say(desc *TaskDesc) {
	time.Sleep(time.Second)
	fmt.Println("No. ", desc.no, " Done: ", desc.msg)
	desc.wg.Done()
}

func main() {
	go func() {fmt.Println("Start.")} ()

	wg := sync.WaitGroup{}
	size := 3
	wg.Add(size)
	for i:=0; i<size; i++ {
		desc := &TaskDesc{
			&wg,
			i,
			"hello",
		}
		say(desc)
	}
	wg.Wait()
	fmt.Print("Finish.")
}
