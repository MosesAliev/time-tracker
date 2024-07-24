package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"time-tracker/internal/database"
	"time-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

// Получить трудозатраты пользователя
func GetUserTimeHandler(c *gin.Context) {
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

	var tasks []models.Task
	res = database.DB.Db.Where("user_passport = ?", passportNumber).Find(&tasks)
	if res.Error != nil {
		log.Println(res.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка"})
		return
	}

	// Перебираем все выполненные и выполняемые задачи пользователем
	var hours = 0
	for _, task := range tasks {
		if task.ID == user.TaskID {
			var now = time.Now()
			hours += now.Hour() - task.StartTime.Hour()
		} else {
			hours += task.EndTime.Hour() - task.StartTime.Hour()
		}

	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d hours", hours)})
}
