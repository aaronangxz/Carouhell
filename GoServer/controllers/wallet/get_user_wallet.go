package wallet

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func GetUserWalletDetails(c *gin.Context) {
	var (
		input         models.GetUserWalletDetailsRequest
		walletDetails models.GetUserWalletDetailsResponse
	)

	if err := models.DB.Raw("SELECT * FROM wallet_tab WHERE user_id = ?", input.UserID).Scan(&walletDetails).Error; err != nil {
		//check if user exists
		if walletDetails.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("user_id does not exist.")})
			log.Printf("user not found:  %v", input.GetUserID())
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserWalletDetails DB query: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieved user wallet details.")})

	data, err := json.Marshal(walletDetails)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v", err.Error())
	}
	log.Printf("Successful: GetUserWalletDetails. Data: %s", data)
}
