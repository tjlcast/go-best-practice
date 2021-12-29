package downloader

import (
	"net/http"
	"io"
	"encoding/json"
	Logger "get-bili/fmt" 
)

type InfoRequest struct {
	Bvids	[]string
}

type VideoInfo struct {
	Code int `json: "code"`
	Data struct {
		Bvid string `json: "bvid"`
		Title string `json: title`
		Desc string `json: "desc"`
	} `json: "data"`
}

type InfoResponse struct {
	Infos []VideoInfo
}

func BatchDownloadVideoInfo(request InfoRequest) (InfoResponse, error) {
	var response InfoResponse

	for _, bvid := range request.Bvids {
		var videoInfo VideoInfo
		url := "https://api.bilibili.com/x/web-interface/view?bvid=" + bvid
		Logger.Info("Get --> ", url)
		resp, err := http.Get(url)
		if err != nil {
			return InfoResponse{}, err
		}

		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return InfoResponse{}, err
		}

		if err := json.Unmarshal(respBytes, &videoInfo); err != nil {
			return InfoResponse{}, err
		}

		if err := resp.Body.Close(); err != nil {
			return InfoResponse{}, err
		}

		response.Infos = append(response.Infos, videoInfo)
	}

	return response, nil
} 