package app

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func createJSONMockList(input queryInput) []memeItem {

	jsonMockList := make([]memeItem, input.NumOfResult)
	for i := 0; i < input.NumOfResult; i++ {
		var jsonMock memeItem

		switch input.Input {
		case "batman slapping robin":
			jsonMock.ImageURL = "https://i.imgflip.com/32hp9x.jpg"
			jsonMock.Title = "Batman Slapping Robin"
			jsonMock.ItemURL = "https://imgflip.com/meme/Batman-Slapping-Robin"
		case "distracted boyfriend":
			jsonMock.ImageURL = "https://i.imgflip.com/2/1ur9b0.jpg"
			jsonMock.Title = "Distracted Boyfriend"
			jsonMock.ItemURL = "https://imgflip.com/meme/Distracted-Boyfriend"
		default:
			width := 150 + 50*rand.Intn(8)
			height := 150 + 50*rand.Intn(8)
			jsonMock.ImageURL = fmt.Sprintf("http://placecorgi.com/%d/%d", width, height)
			jsonMock.Title = fmt.Sprintf("Corgi %d x %d (w x h)", width, height)
			jsonMock.ItemURL = "http://placecorgi.com/"
		}

		jsonMockList[i] = jsonMock
	}

	return jsonMockList
}

func parseQueryInput(params map[string][]string) (queryInput, error) {
	var input queryInput

	queries, exists := params["input"]
	var query string
	if exists {
		query = queries[0]
	}

	numOfResult, err := strconv.Atoi(params["n_result"][0])
	if err != nil {
		log.Println(err.Error())
	}

	pages, exists := params["page"]
	page := 1
	if exists {
		page, err = strconv.Atoi(pages[0])
		if err != nil {
			log.Println(err.Error())
			return input, errors.New("fail to convert page to int")
		}
	}

	input.Input = query
	input.NumOfResult = numOfResult
	input.Page = page

	return input, nil
}
