package database

import (
	"github.com/pkg/errors"
	"github.com/toluwase1/busha-assessment/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type PostgresDB struct{
	DB *gorm.DB
}

type DB interface {
	AddNewCommentToDatabase(comment *models.Comments) (*models.Comments, error)
	GetAllMovieComments(movieId int) (*[]models.Comments, error)
	CountComments(movieId int) (int64, error)
}


func (psql *PostgresDB) InitializeDB() {
	var errorMessages = make(map[string]string)
	var err error
	psqlInfo := os.Getenv("DATABASE_URL")
	if psqlInfo == "" {
		psqlInfo = "postgres://root:password@db:5432/postgres?sslmode=disable"
	}
	DBSession, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		errList["error"] = "Unable to connect"
		err = errors.New("could not open database connection")
		errorMessages["DB error"] = err.Error()
		panic(err)
	}
	err = DBSession.Debug().AutoMigrate(&models.Comments{})
	if err != nil {
		errList["error"] = "migration failed"
		err = errors.New("could not migrate or create tables")
		errorMessages["DB error"] = err.Error()
		panic(err)
	}
	psql.DB = DBSession
}

func (psql *PostgresDB) AddNewCommentToDatabase(comment *models.Comments) (*models.Comments, error) {
	var errorMessages = make(map[string]string)
	err := psql.DB.Create(&comment).Error
	if err != nil {
		errList["error"] = "migration failed"
		err = errors.New("could not migrate or create tables")
		errorMessages["DB error"] = err.Error()
		panic(err)
		return nil, err
	}
	return comment, nil
}

func (psql *PostgresDB) GetAllMovieComments(movieId int) (*[]models.Comments, error) {
	var errorMessages = make(map[string]string)
	var comments []models.Comments
	err := psql.DB.Where("movie_id = ?", movieId).Find(&comments).Error
	if err != nil {
		errList["error"] = "error getting comments"
		err = errors.New("could not retrieve comments from database")
		errorMessages["DB error"] = err.Error()
		panic(err)
		return nil, err
	}
	return &comments, nil
}


func (psql *PostgresDB) CountComments(id int) (int64, error) {
	var errorMessages = make(map[string]string)
	var count int64
	var err error
	err = psql.DB.Model(&models.Comments{}).Where("movie_id = ?", id).Count(&count).Error

	if err != nil {
		errList["error"] = "Unable to get count"
		err = errors.New("Unable to get count")
		errorMessages["count Error"] = err.Error()
		return 0, err
	}
	return count, nil
}