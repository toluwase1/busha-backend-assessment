package models

type CommentRequestEntity struct {
	Content string `json:"body"`
}

type ApiError struct {
	Message string `json:"message"`
}

