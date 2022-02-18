package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/busha-assessment/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

// @Summary Adds a new comment to a post
// @Description Adds a new comment to a post with the post id
// @Accept  json
// @Produce  json
// @Param comment body models.Comments true "Comment"
// @Param movie_id path int true "MovieId"
// @Success 200 {object} models.Comments
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies/{movie_id}/comments [post]
func (server *Server) AddNewComment() gin.HandlerFunc {
	errList = map[string]string{}
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("movie_id"))
		if err != nil {
			errList["Id"] = "Id invalid"
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  errList,
			})
			return
		}

		commentRequest := &models.CommentRequestEntity{}
		err = c.ShouldBindJSON(commentRequest)
		if err != nil {
			errList["error"] = "could not decode request"
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  errList,
			})
			return
		}

		if len(commentRequest.Content) > 500 {
			errList["error"] = "comment has exceeded required limit"
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  errList,
			})
			return
		}
		comment := &models.Comments{
			MovieId:    id,
			Content:   	commentRequest.Content,
			IPAddress:   c.ClientIP(),
			CreatedAt:  time.Now(),
		}
		data, err := server.DB.AddNewCommentToDatabase(comment)
		if err != nil {
			log.Println(err)
			errList["error"] = "could not add comment"
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  errList,
			})
			return
		}

		if !server.increaseRedisCommentCount(id) {
			log.Println("movie id not found in redis")
		}
		c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully!", "data": data})
	}
}



// @Summary Endpoint Gets a list of comments
// @Description Endpoint Gets a list of comments for a movie
// @Produce  json
// @Param movie_id path int true "Movie ID"
// @Success 200 {object} models.Comments
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies/{movie_id}/comments [get]
// GetComments method returns all comments for a particular movie
func (server *Server) GetCommentList() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("movie_id"))
		if err != nil {
			log.Println(err)
			errList["error"] = "could not get comments"
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  errList,
			})
			return
		}
		data, err := server.DB.GetAllMovieComments(id)

		for i := 0; i < len(*data)/2; i++ {
			(*data)[i], (*data)[len(*data)-1-i] = (*data)[len(*data)-1-i], (*data)[i]
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"response":  data,
		})
		c.JSON(http.StatusOK, gin.H{"message": "Comments fetched successfully!", "data": data})
	}
}

func (server *Server) increaseRedisCommentCount(episodeId int) bool {
	var movies = server.Cache.GetMoviesFromCache("movies")
	if movies != nil {
		for i, movie := range *movies {
			if movie.EpisodeId == episodeId {
				hold := models.MovieData{
					EpisodeId:    movie.EpisodeId,
					Title:        movie.Title,
					CommentCount: movie.CommentCount,
					OpeningCrawl: movie.OpeningCrawl,
					ReleaseDate:  movie.ReleaseDate,
				}
				(*movies)[i] = hold
				server.Cache.Set("movies", movies)
				log.Println("comment count increased", movies)
				return true
			}
		}
	}
	return false
}