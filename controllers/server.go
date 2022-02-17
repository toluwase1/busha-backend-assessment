package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/busha-assessment/database"
	"github.com/toluwase1/busha-assessment/helpers"
	_ "gorm.io/driver/postgres"
	"log"
	"net/http"
	"os"
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
	server.Cache = helpers.NewRedisCache("localhost:6379", 1, "", 100)
	server.Router = gin.New()
	server.InitializeRoutes()
	CORSMiddleware()
	//middlewares.CORSMiddleware()

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT = ":8080"
	}
	srv := &http.Server{
		Addr:    PORT,
		Handler: server.Router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	log.Printf("Server started on %s\n", PORT)
}