package app

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const getTrendingMockSeed = 2

// GetTrendingMock returns a list of meme mock
func GetTrendingMock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "can't parse url raw query", http.StatusBadRequest)
	}

	numOfResult, err := strconv.Atoi(params["n_result"][0])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "can't convert n_result from string to int", http.StatusBadRequest)
	}

	jsonMockList := createJSONMockList(numOfResult, getTrendingMockSeed)
	jsonString, _ := json.Marshal(jsonMockList)
	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}
