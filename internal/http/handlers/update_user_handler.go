package handlers

import (
	"log"
	"net/http"
	"time-tracker/internal/database"
	"time-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

// Внести изменения информации о пользователе
func UpdateUserHandler(c *gin.Context) {
	var user models.User
	var err = c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Неправильное тело запроса"})
		return
	}

	var res = database.DB.Db.Where("passport_number = ?", user.PassportNumber).First(&models.User{})
	if res.Error != nil {
		log.Println(res.Error)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Не найден пользователь"})
		return
	}

	res = database.DB.Db.Save(&user)
	if res.Error != nil {
		log.Println(res.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "ошибка"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "обновлено"})
}
