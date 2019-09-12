package app

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

// SearchByTextMock a func to return mock response for search_by_text api
func SearchByTextMock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println("[SearchByTextMock]: parse url query error, " + err.Error())
		http.Error(w, "can't parse url raw query", http.StatusBadRequest)
	}

	input, err := parseQueryInput(params)
	if err != nil {
		log.Println("[SearchByTextMock]: parse query input error, " + err.Error())
		http.Error(w, "can't parse query input from url queries", http.StatusBadRequest)
	}

	jsonMockList := createJSONMockList(input)
	jsonString, _ := json.Marshal(jsonMockList)
	if _, err := w.Write(jsonString); err != nil {
		log.Println("[SearchByTextMock]: write json string error, " + err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}

// SearchByText a func to return response for search_by_text api
func SearchByText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println("[SearchByText]: parse url query error, " + err.Error())
		http.Error(w, "can't parse url raw query", http.StatusBadRequest)
	}

	input, err := parseQueryInput(params)
	if err != nil {
		log.Println("[SearchByText]: parse query input error, " + err.Error())
		http.Error(w, "can't parse query input from url queries", http.StatusBadRequest)
	}

	db, err := connectDB()
	if err != nil {
		log.Println("[SearchByText]: connect db error, " + err.Error())
		http.Error(w, "connect db error", http.StatusBadRequest)
	}

	memeIds, err := getMemeIdsByKeyword(db, input)
	if err != nil {
		log.Println("[SearchByText]: get meme ids by keyword error, " + err.Error())
		http.Error(w, "get ids by keyword error", http.StatusBadRequest)
	}

	memes, err := getMemesByIds(db, memeIds)
	if err != nil {
		log.Println("[SearchByText]: get memes by ids error, " + err.Error())
		http.Error(w, "get memes by ids error", http.StatusBadRequest)
	}

	jsonString, _ := json.Marshal(memes)
	if _, err := w.Write(jsonString); err != nil {
		log.Println("[SearchByText]: write json string error, " + err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}
