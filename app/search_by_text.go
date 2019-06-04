package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type searchOutputItem struct {
	ImageUrl string `json:"image_url"`
	Title    string `json:"title"`
	ItemUrl  string `json:"item_url"`
}

func SearchByTextMock(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	const numOfResults = 9
	jsonMockList := createJsonMockList(numOfResults)
	jsonString, _ := json.Marshal(jsonMockList)
	if _, err := writer.Write(jsonString); err != nil {
		fmt.Println(err.Error())
	}
}

func createJsonMockList(numOfResults int) []searchOutputItem {
	jsonMockList := make([]searchOutputItem, numOfResults)
	for i := 0; i < numOfResults; i++ {
		jsonMock := searchOutputItem{
			ImageUrl: "https://i.imgflip.com/32hp9x.jpg",
			Title:    "Batman Slapping Robin",
			ItemUrl:  "https://imgflip.com/meme/Batman-Slapping-Robin",
		}
		jsonMockList[i] = jsonMock
	}

	return jsonMockList
}
