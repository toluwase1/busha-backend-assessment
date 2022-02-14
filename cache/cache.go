package cache

import "github.com/toluwase1/busha-assessment/models"

type Cache interface {
	Set(key string, value *[]models.MovieData)
	Get(key string) *[]models.MovieData
	SetCharacters(key string, value []models.CharacterList)
	GetCharacters(key string) []models.CharacterList
}
