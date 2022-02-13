package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type CharacterList struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Height float64 `json:"height"`
	TotalNumberOfCharacters string `json:"totalNumberOfCharacters"`
}


func GetCharacterListById(c *gin.Context) (*[]CharacterList, error) {
	errList = map[string]string{}
	movieId := c.Param("movieId")
	url:= Url+"/films/"+movieId
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

	character := []CharacterList{}
	err = json.Unmarshal(body, &character)
	if  err != nil {
		errList["Bad Request"] = "Could not read the specified body from the api"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return nil, err
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": character,
	})
	return &character, nil
}

func GetCharInformation (c *gin.Context) (CharacterList, error){
	errList = map[string]string{}
	path := c.GetString("path")
	url:= Url+"/films/"+path
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

	character := CharacterList{}
	err = json.Unmarshal(body, &character)
	if  err != nil {
		errList["Bad Request"] = "Could not read the specified body from the api"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return CharacterList{}, err
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": character,
	})
	return character, nil
}




