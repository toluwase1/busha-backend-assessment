package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/busha-assessment/models"
	"log"
	"net/http"
	"strconv"
	"time"
)
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
		//errs := s.decode(c, commentRequest)
		//if errs != nil {
		//	errList["error"] = "could not decode request"
		//	log.Println(err)
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"status": http.StatusBadRequest,
		//		"error":  errList,
		//	})
		//	return
		//}

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
			Body:   	commentRequest.Content,
			IPAddress:   c.ClientIP(),
			DateCreated: time.Now(),
		}
		data, err := server.DB.AddComment(comment)
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
		data, err := server.DB.GetComments(id)

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

func (server *Server) increaseRedisCommentCount(movieID int) bool {
	var movies = server.Cache.Get("movies")
	if movies != nil {
		for i, movie := range *movies {
			if movie.MovieId == movieID {
				hold := models.MovieData{
					MovieId:    movie.MovieId,
					Name:        movie.Name,
					CommentCount: movie.CommentCount + 1,
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