package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const searchByTextMockSeed = 1

// SearchByTextMock a func to return mock response for search_by_text api
func SearchByTextMock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Printf("body length: %d\n", len(body))
	fmt.Printf("body content:\n%s\n", body)

	var input searchInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Printf("Error unmarshaling search input json string")
		http.Error(w, "can't unmarshal input json", http.StatusBadRequest)
		return
	}

	jsonMockList := createJSONMockList(input.NumOfResult, searchByTextMockSeed)
	jsonString, _ := json.Marshal(jsonMockList)
	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}
