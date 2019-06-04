package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchByTextMock(writer http.ResponseWriter, request *http.Request) {
	jsonString, _ := json.Marshal("mock search results")
	if _, err := writer.Write(jsonString); err != nil {
		fmt.Println(err.Error())
	}
}
