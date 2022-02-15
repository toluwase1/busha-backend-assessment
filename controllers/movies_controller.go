package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/busha-assessment/models"

	"log"
	"net/http"
	"sort"
)

//gETMOVIESListController
//CacheListWithRedis

func (server *Server) GetMoviesListController() gin.HandlerFunc {
	return func(c *gin.Context) {
		errList = map[string]string{}
		movies := server.Cache.Get("movies")
		if movies == nil {
			data, err := models.FindMoviesFromApi(c)
			if err != nil {
				errList["No_movies"] = "No movies Found"
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": http.StatusInternalServerError,
					"error":  errList,
				})
				return
			}
			result := *data
			sort.Slice(result, func(i, j int) bool {
				return result[i].ReleaseDate > result[j].ReleaseDate
			})
			movies = &result
			//adding comments for uncached movies
			for i, movie := range *movies {
				commentCount, _ := server.DB.CountComments(int(movie.MovieId))
				hold := models.MovieData{
					MovieId:    movie.MovieId,
					Name:         movie.Name,
					CommentCount: commentCount,
					OpeningCrawl: movie.OpeningCrawl,
					ReleaseDate:  movie.ReleaseDate,
				}
				(*movies)[i] = hold
			}


			server.Cache.Set("movies", movies)
			log.Println("Movie List added to cache")
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"response": movies,
		})
	}
}
