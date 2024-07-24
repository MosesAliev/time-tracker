package handlers

import (
	"net/http"
	"strconv"
	"time"
	"time-tracker/internal/database"
	"time-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

// Задача завершена
func EndTimeHandler(c *gin.Context) {
	var query1, _ = c.GetQuery("passportSerie")
	var query2, _ = c.GetQuery("passportNumber")
	var fullQuery = query1 + query2
	var passportNumber, err = strconv.Atoi(fullQuery)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Неверный формат ввода"})
		return
	}

	var user models.User
	var res = database.DB.Db.Where("passport_number = ?", passportNumber).First(&user)
	if res.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Не найден пользователь"})
		return
	}

	var task = models.Task{}
	res = database.DB.Db.Where("id = ?", user.TaskID).First(&task)
	if res.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Не найдена задача"})
		return
	}

	task.EndTime = time.Now()
	res = database.DB.Db.Save(&task)
	if res.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка"})
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Конец задачи"})
}
