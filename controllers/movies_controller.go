package controllers

import (
	"encoding/json"
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
		var newResult []byte
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
			//addd:=Server{}
			//addd.addCommentCountToMovies(movies)
			indexed:=server.addCommentCountToMovies(movies)
			//for i, movie := range *movies {
			//	commentCount, _ := server.DB.CountComments(movie.EpisodeId)
			//	hold := models.MovieData{
			//		EpisodeId:    movie.EpisodeId,
			//		Title:         movie.Title,
			//		CommentCount: commentCount,
			//		OpeningCrawl: movie.OpeningCrawl,
			//		ReleaseDate:  movie.ReleaseDate,
			//	}
			//	(*movies)[i] = hold
			//}
			newResult, _ = json.Marshal(indexed)
			server.Cache.Set("movies", movies)
			log.Println("Movie List added to cache")
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"response": newResult,
		})
	}
}

func (s *Server) addCommentCountToMovies(movies *[]models.MovieData) models.MovieData  {
	var temp models.MovieData
	for idx, movie := range *movies {
		commentCount, _ := s.DB.CountComments(movie.EpisodeId)
		temp = models.MovieData{
			EpisodeId:    movie.EpisodeId,
			Title:         movie.Title,
			CommentCount: commentCount,
			OpeningCrawl: movie.OpeningCrawl,
			ReleaseDate:  movie.ReleaseDate,
		}
		(*movies)[idx] = temp
	}

	return temp
}
