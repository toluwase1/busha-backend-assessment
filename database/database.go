package database

import "github.com/toluwase1/busha-assessment/models"

type DB interface {
	AddComment(comment *models.Comments) (*models.Comments, error)
	GetComments(movieId int) (*[]models.Comments, error)
	CountComments(movieId int) (int64, error)
}
