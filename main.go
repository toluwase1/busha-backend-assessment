package main

import (
	"github.com/toluwase1/busha-assessment/controllers"
)

func main() {
	postgres := &controllers.PostgresDB{}
	postgres.InitializeDB()
}




