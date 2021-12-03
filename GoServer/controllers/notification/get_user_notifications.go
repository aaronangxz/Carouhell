package notification

import (
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func GetUserNotifications(c *gin.Context) {
	var (
		input         models.GetUserNotificationsRequest
		notifications models.GetUserNotificationsResponse
	)
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("UserID cannot be empty.")})
			log.Println("user_id cannot be empty.")
			return
		}
		if !utils.ValidateUint(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("UserID must be uint type.")})
			log.Println("user_id must be uint type.")
			return
		}
	}

	if err := models.DB.Raw(utils.GetNotificationQuery(), input.GetUserID(), input.GetUserID(), input.GetUserID()).Scan(&notifications).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserNotifications - get notifications DB query: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieved notifications."), "Data": notifications})
	log.Printf("Successful: GetUserNotifications: %v - DB", input.GetUserID())
}
