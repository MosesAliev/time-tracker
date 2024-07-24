package main

import (
	"time-tracker/internal/database"
	"time-tracker/internal/http/router"
)

func main() {
	database.ConnectDB()       // Подключение к БД
	router.SetupAndRunRouter() // запуск сервера
}
