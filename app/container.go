package app

type memeIcon struct {
	ImageURL string `json:"image_url"`
	Title    string `json:"title"`
	ItemURL  string `json:"item_url"`
}

type queryInput struct {
	Input       string `json:"input,omitempty"`
	NumOfResult int    `json:"n_result"`
	Page        int    `json:"page"`
}

type trendingInput struct {
	NumOfResult int `json:"n_result"`
	Page        int `json:"page"`
}

type memeDetail struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	ImageURL string   `json:"image_url"`
	About    string   `json:"about"`
	Tags     []string `json:"tags,omitempty"`
}
