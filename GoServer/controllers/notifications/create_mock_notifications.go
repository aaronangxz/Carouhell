package notifications

import (
	"fmt"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func CreateMockNotifications(c *gin.Context) {
	// Validate input
	var (
		input models.CreateNotificationRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be uint type.")})
			return
		}
		if input.NotificationText == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("notification_text cannot be empty.")})
			return
		}
		if !utils.ValidateString(input.NotificationText) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("notification_text must be string type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetUserID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be > 0.")})
		return
	}

	if input.GetNotificationText() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("notification_text cannot be empty.")})
		return
	}

	//Check req params
	if len(input.GetNotificationText()) > int(models.MaxNotificationTextLength) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("notification_text must be < " + fmt.Sprint(models.MaxNotificationTextLength) + " chars.")})
		return
	}

	if err := models.DB.Exec("INSERT INTO notifications (user_id, notification_text) VALUES (?,?)", input.UserID, input.NotificationText).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
