package app

type memeItem struct {
	ImageURL string `json:"image_url"`
	Title    string `json:"title"`
	ItemURL  string `json:"item_url"`
}

type queryInput struct {
	Input       string `json:"input,omitempty"`
	NumOfResult int    `json:"n_result"`
	Page        int    `json:"page"`
}
