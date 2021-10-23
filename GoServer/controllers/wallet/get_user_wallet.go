package wallet

import (
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func GetUserWallet(c *gin.Context) {
	var (
		input         models.GetUserWalletRequest
		walletDetails models.GetUserWalletResponse
	)

	if err := models.DB.Raw("SELECT * FROM wallet_tab WHERE user_id = ?", input.UserID).Scan(&walletDetails).Error; err != nil {
		//check if user exists
		if walletDetails.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("user_id does not exist.")})
			log.Printf("user not found:  %v", input.GetUserID())
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during DB query: %v", err.Error())
		return
	}

	//join wallet_transaction to get list of transaction
}
