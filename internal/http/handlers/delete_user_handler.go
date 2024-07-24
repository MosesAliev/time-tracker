package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time-tracker/internal/database"
	"time-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

// Удалить пользователя
func DeleteUserHandler(c *gin.Context) {
	var query1, _ = c.GetQuery("passportSerie")
	var query2, _ = c.GetQuery("passportNumber")
	var fullQuery = query1 + query2
	var passportNumber, err = strconv.Atoi(fullQuery)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Неверный формат ввода"})
		return
	}

	var res = database.DB.Db.Where("passport_number = ?", passportNumber).Unscoped().Delete(&models.User{})
	if res.Error != nil {
		log.Println(res.Error)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Не найден пользователь"})
		return
	}

	res = database.DB.Db.Where("user_passport = ?", passportNumber).Unscoped().Delete(&models.Task{})
	if res.Error != nil {
		log.Println("У пользователя не было задач")
	}

	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Удалено"})
}
