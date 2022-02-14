package models

import (
	"encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
 )

type MovieData struct {
	MovieId int64 `json:"movie_id"`
	Name string `json:"name"`
	ReleaseDate string `json:"release_date"`
	OpeningCrawl string `json:"opening_crawl"`
	CommentCount int64 `json:"comment_count"`
}

type SomeMovie struct {
	Results [] MovieData `json:"results"`
}
const Url = "https://swapi.dev"

func GetJson(url string, target interface{}) error{
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}

func FindMoviesFromApi(c *gin.Context) (*[]MovieData, error) {
	errList = map[string]string{}

	url:= Url+"/films/"
	response, err := http.Get(url)
	if err != nil {
		errList["No_movies"] = "No movies Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errList["Bad Request"] = "Could not read the specified body"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return nil, err
	}
	movies := &SomeMovie{}
	err = json.Unmarshal(body, movies)
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
		"response": movies,
	})

	return &movies.Results, nil

}

