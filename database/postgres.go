package database

import (
	"github.com/pkg/errors"
	"github.com/toluwase1/busha-assessment/models"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// PostgresDB implements the DB interface
type PostgresDB struct {
	DB *gorm.DB
}

func (pgsql *PostgresDB) InitializeDb() {
	var errorMessages = make(map[string]string)
	var err error
	psqlInfo := os.Getenv("DATABASE_URL")
	if psqlInfo == "" {
		psqlInfo = "postgres://root:toluwase@db:5432/busha?sslmode=disable"
	}
	DBSession, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		errList["error"] = "Unable to connect"
		err = errors.New("could not open database connection")
		errorMessages["DB error"] = err.Error()
		panic(err)
		return
	}
	pgsql.DB = DBSession
	err = pgsql.DB.AutoMigrate(&models.Comments{})
	if err != nil {
		errList["error"] = "migration failed"
		err = errors.New("could not migrate or create tables")
		errorMessages["DB error"] = err.Error()
		panic(err)
		return
	}
}

func (pgsql PostgresDB) AddNewCommentToDatabase(comment *models.Comments) (*models.Comments, error) {
	var errorMessages = make(map[string]string)
	err := pgsql.DB.Create(&comment).Error
	if err != nil {
		errList["error"] = "migration failed"
		err = errors.New("could not migrate or create tables")
		errorMessages["DB error"] = err.Error()
		panic(err)
		return nil, err
	}
	return comment, nil
}

func (pgsql *PostgresDB) GetAllMovieComments(movieId int) (*[]models.Comments, error) {
	var errorMessages = make(map[string]string)
	var comments []models.Comments
	err := pgsql.DB.Where("movie_id = ?", movieId).Find(&comments).Error
	if err != nil {
		errList["error"] = "error getting comments"
		err = errors.New("could not retrieve comments from database")
		errorMessages["DB error"] = err.Error()
		panic(err)
		return nil, err
	}
	return &comments, nil
}


func (pgsql *PostgresDB) CountComments(id int) (int64, error) {
	var errorMessages = make(map[string]string)
	var count int64
	var err error
	err = pgsql.DB.Model(&models.Comments{}).Where("movie_id = ?", id).Count(&count).Error

	if err != nil {
		errList["error"] = "Unable to get count"
		err = errors.New("Unable to get count")
		errorMessages["count Error"] = err.Error()
		return 0, err
	}
	return count, nil
}