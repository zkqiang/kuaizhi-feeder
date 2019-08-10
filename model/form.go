package model

type Card struct {
	Images []string `json:"images"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Url    string   `json:"url"`
	Video  string   `json:"video"`
}

type FeedForm struct {
	JobId string  `json:"job_id"`
	Cards *[]Card `json:"cards"`
}
