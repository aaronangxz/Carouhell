package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func ValidateGetUserWalletDetailsRequest(c *gin.Context, input *models.GetUserWalletDetailsRequest) error {
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

func GetUserWalletDetails(c *gin.Context) {
	var (
		input         models.GetUserWalletDetailsRequest
		walletDetails models.GetUserWalletDetailsResponse
	)

	if err := ValidateGetUserWalletDetailsRequest(c, &input); err != nil {
		log.Printf("Error during ValidateGetUserWalletDetailsRequest: %v", err.Error())
		return
	}

	if err := models.DB.Raw("SELECT * FROM wallet_tab WHERE wallet_id = ?", input.UserID).Scan(&walletDetails).Error; err != nil {
		//check if user exists
		if walletDetails.WalletID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("user_id does not exist.")})
			log.Printf("wallet not found:  %v", input.GetUserID())
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserWalletDetails DB query: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieved user wallet details."), "Data": walletDetails})

	data, err := json.Marshal(walletDetails)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v", err.Error())
	}
	log.Printf("Successful: GetUserWalletDetails. Data: %s", data)
}
