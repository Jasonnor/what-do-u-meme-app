package app

import (
	"encoding/json"
	"net/http"
)

func GetTrendingMock(writer http.ResponseWriter, request *http.Request) {
	json_string, _ := json.Marshal("mock trending")
	writer.Write(json_string)
}
