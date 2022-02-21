package models

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type CharacterList struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
}
type CharacterListSlice struct {
	CharacterList []string `json:"characters"`
}

func GetCharacterListById(id int) (*[]string, error) {
	var c *gin.Context
	errList = map[string]string{}
	convId := strconv.Itoa(id)
	url := UrlAlt + "/films/" + convId
	response, err := http.Get(url)
	if err != nil {
		errList["No_movie"] = "Could not find Character"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errList["No_movie"] = "Could not read the specified body from the api"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return nil, err
	}

	var character CharacterListSlice
	err = json.Unmarshal(body, &character)
	if err != nil {
		errList["Bad Request"] = "Could not read the specified body from the api"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return nil, err
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"status":   http.StatusOK,
	//	"response": character,
	//})
	return &character.CharacterList, nil
}

func GetCharInformation(id string) (CharacterList, error) {
	url := fmt.Sprintf("%s/people/%s/", UrlAlt, id)
	var c *gin.Context
	errList = map[string]string{}
	response, err := http.Get(url)

	if err != nil {
		errList["No_character"] = "Could not get Character from link"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return CharacterList{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errList["No_character"] = "Could not get Character from link"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return CharacterList{}, err
	}

	var character CharacterList
	err = json.Unmarshal(body, &character)
	if err != nil {
		errList["Bad Request"] = "Could not read the specified body from the api"
		return CharacterList{}, err
	}
	return character, nil
}
