package controllers

import (
	"github.com/toluwase1/busha-assessment/cache"
	"github.com/toluwase1/busha-assessment/database"
)

type Server struct {
	DB    database.DB
	Cache cache.Cache
}