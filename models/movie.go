package models

type MovieData struct {
	MovieId int64 `json:"movie_id"`
	Name string `json:"name"`
	ReleaseDate string `json:"release_date"`
	OpeningCrawl string `json:"opening_crawl"`
	CommentCount int `json:"comment_count"`
}

