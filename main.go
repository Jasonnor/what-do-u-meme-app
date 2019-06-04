package main

import (
	"net/http"
	"what-do-u-meme-app/app"
)

func main() {
	// add func to handle url request
	http.HandleFunc("/mock/search_by_text", app.SearchByTextMock)
	http.HandleFunc("/mock/get_trending", app.GetTrendingMock)

	// listen and serve
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
