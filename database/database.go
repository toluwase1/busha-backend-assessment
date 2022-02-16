package database

import "github.com/toluwase1/busha-assessment/models"

type DB interface {
	AddNewCommentToDatabase(comment *models.Comments) (*models.Comments, error)
	GetAllMovieComments(movieId int) (*[]models.Comments, error)
	CountComments(movieId int) (int64, error)
}
