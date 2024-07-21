package main

import (
	"Go/Go/time-tracker/internal/database"
	"Go/Go/time-tracker/internal/http/router"
)

func main() {
	database.ConnectDB()
	router.SetupAndRunRouter()
}
