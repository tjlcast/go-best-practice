package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0
	for {
		time.Sleep(time.Second)
		fmt.Printf("%d, hello world!\n", counter)
		counter+=1
	}
}