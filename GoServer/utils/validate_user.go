package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func ValidateUserID(c *gin.Context, UserID uint32) error {
	var hold models.Account
	//check if seller exists
	if err := models.DB.Raw("SELECT * FROM acc_tab WHERE a_user_id = ?", UserID).Scan(&hold).Error; err != nil {
		if hold.AUserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("user_id does not exist.")})
			errormsg := fmt.Sprintf("user_id does not exist. input: %v", UserID)
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		errormsg := fmt.Sprint("DB error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}
