package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/busha-assessment/database"
	"github.com/toluwase1/busha-assessment/helpers"
	middlewares "github.com/toluwase1/busha-assessment/middleware"
	_ "gorm.io/driver/postgres"
)


type Server struct {
	DB     *database.PostgresDB
	Router *gin.Engine
	Cache  helpers.Cache
}

func (server *Server) Start() {
	DB := &database.PostgresDB{}
	DB.InitializeDB()
	server.DB = DB
	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())
	server.Cache = helpers.NewRedisCache("localhost:6379", 1, "", 100)
}