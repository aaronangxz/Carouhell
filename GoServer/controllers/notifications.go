package controllers

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func GetNotificationsByUserID(c *gin.Context) {
	var (
		UserNotifications []models.GetNotificationsByUserIDResposne
		response          models.ResponseMeta
		input             = c.Param("user_id")
	)

	if err := models.DB.Raw("SELECT notification_id, notification_text FROM notifications WHERE user_id = ? ORDER BY notification_id DESC", input).
		Scan(&UserNotifications).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewErrorResponse(err)})
		return
	}

	//Build response
	if len(UserNotifications) == 0 {
		response = models.NewNotFoundResponse()
	} else {
		response = models.NewSuccessResponse()
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": response, "data": UserNotifications})
}

func CreateMockNotifications(c *gin.Context) {
	// Validate input
	var (
		input models.CreateNotificationRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewErrorResponse(err)})
		return
	}

	if err := models.DB.Exec("INSERT INTO notifications (user_id, notification_text) VALUES (?,?)", input.UserID, input.NotificationText).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewErrorResponse(err)})

	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
