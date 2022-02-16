package controllers

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/movies", server.GetMoviesListController())
		v1.POST("/movies/:movie_id/comments", server.AddNewComment())
		v1.GET("/movies/:movie_id/comments", server.GetCommentList())
		v1.POST("/movies/:movie_id/characters", server.GetCharacterList())
	}
}
