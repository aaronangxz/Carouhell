package controllers

import (
	"fmt"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func GetNotificationsByUserID(c *gin.Context) {
	var (
		userNotifications []models.GetNotificationsByUserIDResponse
		input             *models.GetNotificationsByUserIDRequest
		extraCondition    string
	)

	//Check required fields
	if err := c.ShouldBindJSON(&input); err != nil {
		//user_id cannot be nil
		if input.GetUserID() == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	//Check optional fields
	//Treat nil as 0, so we can execute switch case below
	if input.GetLimit() == nil {
		input.Limit = models.DefaultNotificationResponseLimit
	}

	//Limit cannot > MaxNotificationResponseSize
	if models.ValidateLimitMax(input.GetLimit(), models.MaxNotificationResponseSize) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("limit cannot exceed 50")})
		return
	}

	//Query based on request params
	//0: when limit is nil or 0 (don't limit)
	//default: when 1 < limit < MaxNotificationResponseSize
	switch *input.GetLimit() {
	case 0:
		extraCondition = ""
	default:
		extraCondition = " LIMIT " + fmt.Sprint(*input.GetLimit())
	}

	query := "SELECT notification_id, notification_text FROM notifications WHERE user_id = ? ORDER BY notification_id DESC" + extraCondition

	if err := models.DB.Raw(query, input.GetUserID()).
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
	// if len(input.GetNotificationText()) > models.MaxNotificationTextLength {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("notification_text must be < " + fmt.Sprint(models.MaxNotificationTextLength) + " chars.")})
	// 	return
	// }

	if err := models.DB.Exec("INSERT INTO notifications (user_id, notification_text) VALUES (?,?)", input.UserID, input.NotificationText).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
