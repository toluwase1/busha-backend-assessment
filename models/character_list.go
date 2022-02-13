package models

type CharacterList struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Height float64 `json:"height"`
	TotalNumberOfCharacters string `json:"totalNumberOfCharacters"`
}