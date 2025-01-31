package handlers

import (
	"log"
	"net/http"
	"time-tracker/internal/database"
	"time-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

// Добавить пользователя
func AddUserHandler(c *gin.Context) {
	var user models.User
	var err = c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Неправильное тело запроса"})
		return
	}

	var res = database.DB.Db.Create(&user)
	if res.Error != nil {
		log.Println(res.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "ошибка"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
