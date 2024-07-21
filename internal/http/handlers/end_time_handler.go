package handlers

import (
	"Go/Go/time-tracker/internal/database"
	"Go/Go/time-tracker/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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

}
