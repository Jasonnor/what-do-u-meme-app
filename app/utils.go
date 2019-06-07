package app

import (
	"fmt"
	"math/rand"
)

func createJSONMockList(numOfResults int, input string) []memeItem {

	jsonMockList := make([]memeItem, numOfResults)
	for i := 0; i < numOfResults; i++ {
		var jsonMock memeItem

		switch input {
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
