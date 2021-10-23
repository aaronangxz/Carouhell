package wallet

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func CreateUserWallet(c *gin.Context) {
	var (
		input models.CreateUserWalletRequest
	)

	wallet := models.Wallet{
		WalletID:      input.UserID,
		UserID:        input.UserID,
		WalletBalance: utils.Uint32(0),
		WalletStatus:  utils.Uint32(constant.WALLET_STATUS_ACTIVE),
		LastTopUp:     nil,
		LastUsed:      utils.Int64(time.Now().Unix()),
	}

	if err := models.DB.Table("wallet_tab").Create(&wallet).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during DB query: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully activated wallet.")})

	data, err := json.Marshal(wallet)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v", err.Error())
	}
	log.Printf("Successful: CreateUserWallet. Data: %s", data)
}
