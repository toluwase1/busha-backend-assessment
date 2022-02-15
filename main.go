package main

import "github.com/toluwase1/busha-assessment/database"

func main() {
	postgres := &database.PostgresDB{}
	postgres.InitializeDB()
}




