package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"what-do-u-meme-app/app"
)

func SearchByTextMock(writer http.ResponseWriter, request *http.Request) {
	jsonString, _ := json.Marshal("mock search results")
	if _, err := writer.Write(jsonString); err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// add func to handle url request
	// http.HandleFunc("/get_trending", app.GetTrending)
	http.HandleFunc("/mock/search_by_text", SearchByTextMock)
	http.HandleFunc("/mock/get_trending", app.GetTrendingMock)

	// listen and serve
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
