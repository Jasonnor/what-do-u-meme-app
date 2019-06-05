package app

func createJSONMockList(numOfResults int, seed int) []memeItem {
	if seed == 0 {
		seed = 1
	}

	jsonMockList := make([]memeItem, numOfResults)
	for i := 0; i < numOfResults; i++ {
		var jsonMock memeItem

		switch seed {
		case 1:
			jsonMock.ImageURL = "https://i.imgflip.com/32hp9x.jpg"
			jsonMock.Title = "Batman Slapping Robin"
			jsonMock.ItemURL = "https://imgflip.com/meme/Batman-Slapping-Robin"
		case 2:
			jsonMock.ImageURL = "https://i.imgflip.com/2/1ur9b0.jpg"
			jsonMock.Title = "Distracted Boyfriend"
			jsonMock.ItemURL = "https://imgflip.com/meme/Distracted-Boyfriend"
		}

		jsonMockList[i] = jsonMock
	}

	return jsonMockList
}
