package controllers

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func NewErrorResponse(err error) models.ResponseMeta {
	return models.ResponseMeta{
		DebugMsg:  err.Error(),
		ErrorCode: models.CONST_GET_NOTIFICATION_ERROR_NOT_FOUND,
	}
}

func NewSuccessResponse() models.ResponseMeta {
	return models.ResponseMeta{
		DebugMsg:  "",
		ErrorCode: models.CONST_NOTIFICATION_SUCCESS,
	}
}

func GetNotificationsByUserID(c *gin.Context) {
	var UserNotifications []models.GetNotificationsByUserIDResposne
	input := c.Param("user_id")

	if err := models.DB.Raw("SELECT notification_id, notification_text FROM notifications WHERE user_id = ? ORDER BY notification_id DESC", input).
		Scan(&UserNotifications).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": NewErrorResponse(err)})
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": NewSuccessResponse(), "data": UserNotifications})
}

func CreateMockNotifications(c *gin.Context) {
	// Validate input
	var (
		input models.CreateNotificationRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": NewErrorResponse(err)})
		return
	}

	if err := models.DB.Exec("INSERT INTO notifications (user_id, notification_text) VALUES (?,?)", input.UserID, input.NotificationText).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": NewErrorResponse(err)})

	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": NewSuccessResponse()})
}
