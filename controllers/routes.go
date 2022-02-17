package controllers

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func (server *Server) InitializeRoutes() {

	v1 := server.Router.Group("/api/v1")
	{
		v1.GET("/movies", server.GetMoviesListController())
		v1.POST("/movies/:movie_id/comments", server.AddNewComment())
		v1.GET("/movies/:movie_id/comments", server.GetCommentList())
		v1.POST("/movies/:movie_id/characters", server.GetCharacterList())
	}
	server.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
