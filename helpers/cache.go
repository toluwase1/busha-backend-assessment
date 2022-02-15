package helpers

import "github.com/toluwase1/busha-assessment/models"

type Cache interface {
	Set(key string, value *[]models.MovieData)
	GetMoviesFromCache(key string) *[]models.MovieData
	SetCharToCache(key string, value []models.CharacterList)
	GetCharactersFromCache(key string) []models.CharacterList
}
