package app

import (
	"fmt"
	"math/rand"
)

func createJSONMockList(numOfResults int, seed int) []memeIcon {
	if seed == 0 {
		seed = 1
	}

	jsonMockList := make([]memeIcon, numOfResults)
	for i := 0; i < numOfResults; i++ {
		var jsonMock memeIcon

		switch seed {
		case 1:
			jsonMock.ImageURL = "https://i.imgflip.com/32hp9x.jpg"
			jsonMock.Title = "Batman Slapping Robin"
			jsonMock.ItemURL = "https://imgflip.com/meme/Batman-Slapping-Robin"
		case 2:
			jsonMock.ImageURL = "https://i.imgflip.com/2/1ur9b0.jpg"
			jsonMock.Title = "Distracted Boyfriend"
			jsonMock.ItemURL = "https://imgflip.com/meme/Distracted-Boyfriend"
		case 3:
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
