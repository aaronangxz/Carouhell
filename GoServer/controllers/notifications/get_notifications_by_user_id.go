package notifications

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
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be uint type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetUserID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be > 0.")})
		return
	}

	//Check optional fields
	//Treat nil as 0, so we can execute switch case below
	if input.Limit == nil {
		input.Limit = models.SetDefaultNotificationResponseLimit()
	}

	//Limit cannot > MaxNotificationResponseSize
	if utils.ValidateLimitMax(input.GetLimit(), models.MaxNotificationResponseSize) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("limit cannot exceed 50.")})
		return
	}

	//Query based on request params
	//0: when limit is nil or 0 (don't limit)
	//default: when 1 < limit < MaxNotificationResponseSize
	switch input.GetLimit() {
	case 0:
		extraCondition = ""
	default:
		extraCondition = " LIMIT " + fmt.Sprint(input.GetLimit())
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