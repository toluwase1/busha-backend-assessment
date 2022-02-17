package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/busha-assessment/models"

	"log"
	"net/http"
	"sort"
)

// @Summary      Route Gets all movies
// @Description  This route Gets movies starting oldest release date to the newest either in the cache or from the Api, data from the cache takes priority
// @Produce  json
// @Success 200 {object} models.MovieData
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies [get]

func (server *Server) GetMoviesListController() gin.HandlerFunc {
	return func(c *gin.Context) {
		errList = map[string]string{}
		movies := server.Cache.GetMoviesFromCache("movies")
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
