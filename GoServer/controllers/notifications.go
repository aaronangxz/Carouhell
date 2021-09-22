package controllers

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func GetNotificationsByUserID(c *gin.Context) {
	var (
		userNotifications []models.GetNotificationsByUserIDResposne
		input             models.GetNotificationsByUserIDRequest
	)

	if err := models.DB.Raw(
		"SELECT notification_id, notification_text FROM"+
			"notifications WHERE user_id = ?"+
			"ORDER BY notification_id DESC LIMIT ?", input.UserID, input.Limit).
		Scan(&userNotifications).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	//Build response
	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetNotificationsByUserIDResult(userNotifications), "Data": userNotifications})
}

func CreateMockNotifications(c *gin.Context) {
	// Validate input
	var (
		input models.CreateNotificationRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	//Check req params
	if len(input.NotificationText) > models.MaxNotificationTextLength {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse()})
		return
	}

	if err := models.DB.Exec("INSERT INTO notifications (user_id, notification_text) VALUES (?,?)", input.UserID, input.NotificationText).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
