package router

import (
	"time-tracker/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

// Настройка и запуск роутера
func SetupAndRunRouter() {
	r := gin.Default()
	r.POST("/adduser", handlers.AddUserHandler)
	r.GET("/getusertime", handlers.GetUserTimeHandler)
	r.PATCH("/updateuser", handlers.UpdateUserHandler)
	r.PATCH("/starttime", handlers.StartTimeHandler)
	r.PATCH("/endtime", handlers.EndTimeHandler)
	r.DELETE("/deleteuser", handlers.DeleteUserHandler)
	r.Run()
}
