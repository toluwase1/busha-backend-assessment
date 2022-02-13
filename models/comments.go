package models

import "time"

type Comments struct {
	ID uint64 `json:"id"`
	MovieId int `json:"movie_id"`
	Body string `json:"body" validate:"required,min=2,max=500"`
	IPAddress string `json:"ip_address"`
	DateCreated time.Time `json:"date_created"`
}