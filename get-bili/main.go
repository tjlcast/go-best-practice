package main

import (
	fmt "get-bili/fmt" 
	Logger "get-bili/fmt" 
	"get-bili/downloader"
	"encoding/json"
)

func main() {
	fmt.Logger.Println("Hello world.")
	fmt.Println("Hello world pro.")

	request := downloader.InfoRequest{
		Bvids: []string{"BV1Ff4y187q9", "BV1Hb4y1Z7nw"},
	}
	response, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}

	fmt.Info("Success number: ", len(response.Infos))
	for _, info := range response.Infos {
		bs, _ := json.Marshal(info)
		fmt.Info(string(bs))
	}
	Logger.Info("Finish!")
}
