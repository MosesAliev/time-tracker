package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"time-tracker/internal/database"
	"time-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

// Начать отсчёт
func StartTimeHandler(c *gin.Context) {
	var query1, _ = c.GetQuery("passportSerie")
	var query2, _ = c.GetQuery("passportNumber")
	var fullQuery = query1 + query2
	var passportNumber, err = strconv.Atoi(fullQuery)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Неверный формат ввода"})
		return
	}

	var user = models.User{}
	var res = database.DB.Db.Where("passport_number = ?", passportNumber).First(&user)
	if res.Error != nil {
		log.Println(res.Error)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Не найден пользователь"})
		return
	}

	var task = models.Task{}
	task.StartTime = time.Now()
	task.UserPassport = user.PassportNumber
	res = database.DB.Db.Create(&task)
	if res.Error != nil {
		log.Println(res.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка"})
		return
	}

	user.TaskID = task.ID
	res = database.DB.Db.Save(&user)
	if res.Error != nil {
		log.Println(res.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Время пошло"})
}
