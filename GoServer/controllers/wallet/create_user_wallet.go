package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func ValidateCreateUserWalletRequest(c *gin.Context, input *models.CreateUserWalletRequest) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			errormsg := "user_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be uint type.")})
			errormsg := fmt.Sprintf("user_id must be uint type. input: %v", input.GetUserID())
			return errors.New(errormsg)
		}

		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func CreateUserWallet(c *gin.Context) {
	var (
		input models.CreateUserWalletRequest
		hold  models.Account
	)

	if err := ValidateCreateUserWalletRequest(c, &input); err != nil {
		log.Printf("Error during ValidateCreateUserWalletRequest: %v", err.Error())
		return
	}

	if err := models.DB.Raw("SELECT * FROM acc_tab WHERE user_id = ?", input.GetUserID()).Scan(&hold).Error; err != nil {
		//check if user exists
		if hold.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("user_id does not exist.")})
			log.Printf("user not found:  %v", input.GetUserID())
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during DB query: %v", err.Error())
		return
	}

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
		log.Printf("Error during CreateUserWallet DB query: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully activated wallet.")})

	data, err := json.Marshal(wallet)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v", err.Error())
	}
	log.Printf("Successful: CreateUserWallet. Data: %s", data)
}
