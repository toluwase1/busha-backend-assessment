package models

type CommentRequestEntity struct {
	Content string `json:"content"`
}

type ApiError struct {
	Message string `json:"message"`
}

