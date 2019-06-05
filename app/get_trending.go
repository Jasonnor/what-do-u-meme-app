package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const getTrendingMockSeed = 2

// GetTrendingMock returns a list of meme mock
func GetTrendingMock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Printf("body length: %d\n", len(body))
	fmt.Printf("body content:\n%s\n", body)

	var input trendingInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Printf("Error unmarshaling search input json string")
		http.Error(w, "can't unmarshal input json", http.StatusBadRequest)
		return
	}

	jsonMockList := createJSONMockList(input.NumOfResult, getTrendingMockSeed)
	jsonString, _ := json.Marshal(jsonMockList)
	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}
