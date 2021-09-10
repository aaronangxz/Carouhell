package controllers

import (
	"fmt"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func GetNotificationsByUserID(c *gin.Context) {
	var UserNotifications []models.GetNotificationsByUserIDResposne
	input := c.Param("user_id")

	err := models.DB.Raw("SELECT notification_id, notification_text FROM notifications WHERE user_id = ? ORDER BY notification_id DESC", input).Scan(&UserNotifications)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"resp": err})
	}

	c.JSON(http.StatusOK, gin.H{"resp": "success", "data": UserNotifications})
}

func CreateMockNotifications(c *gin.Context) {
	// Validate input
	var (
		input models.CreateNotificationRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"resp": err.Error()})
		return
	}

	if err := models.DB.Exec("INSERT INTO notifications (user_id, notification_text) VALUES (?,?)", input.UserID, input.NotificationText).Error; err != nil {
		fmt.Println("Error:", err)
	}

	c.JSON(http.StatusOK, gin.H{"resp": "success"})
}
