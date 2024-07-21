package router

import (
	"Go/Go/time-tracker/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAndRunRouter() {
	r := gin.Default()
	r.POST("/adduser", handlers.AddUserHandler)
	r.PATCH("/starttime", handlers.StartTimeHandler)
	r.PATCH("/endtime", handlers.EndTimeHandler)
	r.Run()
}
