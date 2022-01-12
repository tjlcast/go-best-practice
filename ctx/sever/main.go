package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler start")
	ctx := r.Context()

	complete := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		complete <- struct{}{}
	}()

	select {
	case <- complete:
		fmt.Println("Completed!")
	case <- time.After(5 * time.Second):
		fmt.Println("finish doing something")
	case <- ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("handle ends")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

