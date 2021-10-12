package utils

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func ValidateUserID(c *gin.Context, SellerID uint32) {
	var hold models.Account
	//check if seller exists
	if err := models.DB.Raw("SELECT * FROM acc_tab WHERE user_id = ?", SellerID).Scan(&hold).Error; err != nil {
		if hold.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("seller_id does not exist.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}
}
