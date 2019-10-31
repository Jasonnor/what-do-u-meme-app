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
		log.Println(err.Error())
		http.Error(w, "can't parse url raw query", http.StatusBadRequest)
	}

	input, err := parseQueryInput(params)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "can't parse query input from url queries", http.StatusBadRequest)
	}

	jsonMockList := createJSONMockList(input)
	jsonString, _ := json.Marshal(jsonMockList)
	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}
