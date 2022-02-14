package database

type DB interface {
	AddComment(comment *models.Comment) (*models.Comment, error)
	GetComments(movieId int) (*[]models.Comment, error)
	CountComments(movieId int) (int64, error)
}
