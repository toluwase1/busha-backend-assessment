package models

import "time"

type Comments struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;"`
	MovieId int `json:"movie_id"`
	Content string `json:"content" validate:"required,min=2,max=500"`
	IPAddress string `json:"ip_address"`
	CreatedAt time.Time `json:"created_at"`
}

