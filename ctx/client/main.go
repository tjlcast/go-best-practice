package main

import (
	"time"
	"context"
	"io"
	"log"
	"net/http"
	"fmt"
)

func main() {
	ctx := context.Background() 
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req = req.WithContext(ctx)

	// http.Get("http://www.baidu.com")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	fmt.Println(string(respBytes))
}