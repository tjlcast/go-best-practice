package fmt

import (
	"log"
	"io"
	"os"
)

var Logger *log.Logger

func init() {
	file, err := os.OpenFile("trace.txt", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	//Logger = log.New(os.Stdout, "Log: ", log.LstdFlags)
	Logger = log.New(io.MultiWriter(os.Stdout, file), "Log: ", log.LstdFlags)
}

func Println(args ...interface{}) {
	Logger.Println(args...)
}