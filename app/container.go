package app

type memeItem struct {
	ImageURL string `json:"image_url"`
	Title    string `json:"title"`
	ItemURL  string `json:"item_url"`
}

type searchInput struct {
	Input       string `json:"input"`
	NumOfResult int    `json:"n_result"`
	Page        int    `json:"page"`
}

type trendingInput struct {
	NumOfResult int `json:"n_result"`
	Page        int `json:"page"`
}
