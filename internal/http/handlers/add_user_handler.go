package handlers

import (
	"Go/Go/time-tracker/internal/database"
	"Go/Go/time-tracker/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserHandler(c *gin.Context) {
	var user models.User
	var err = c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Неправильное тело запроса"})
		return
	}

	database.DB.Db.Create(&user)
	c.IndentedJSON(http.StatusOK, user)
}
