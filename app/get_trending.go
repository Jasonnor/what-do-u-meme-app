package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTrendingMock(writer http.ResponseWriter, request *http.Request) {
	jsonString, _ := json.Marshal("mock trending")
	if _, err := writer.Write(jsonString); err != nil {
		fmt.Println(err.Error())
	}
}
