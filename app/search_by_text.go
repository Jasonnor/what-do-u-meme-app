package app

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const searchByTextMockSeed = 3

// SearchByTextMock a func to return mock response for search_by_text api
func SearchByTextMock(w http.ResponseWriter, r *http.Request) {
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

	jsonMockList := createJSONMockList(numOfResult, searchByTextMockSeed)
	jsonString, _ := json.Marshal(jsonMockList)
	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}

// SearchByText a func to return response for search_by_text api
func SearchByText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//params, err := url.ParseQuery(r.URL.RawQuery)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, "can't parse url raw query", http.StatusBadRequest)
	//}

	// numOfResult, err := strconv.Atoi(params["n_result"][0])

	db, err := connectDB()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "connect db error", http.StatusBadRequest)
	}
	memeIds := []int{1, 2}
	memes, err := getMemesByIds(db, memeIds)
	jsonString, _ := json.Marshal(memes)

	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}
