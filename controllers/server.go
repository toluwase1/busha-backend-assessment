package controllers

import (
	"github.com/toluwase1/busha-assessment/database"
	"github.com/toluwase1/busha-assessment/helpers"
)

type Server struct {
	DB    database.DB
	Cache helpers.Cache
}