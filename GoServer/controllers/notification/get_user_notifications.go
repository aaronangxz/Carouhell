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
		input models.GetUserNotificationsRequest
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
}
