package wallet

import (
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
		transactions  []models.WalletTransactionsWithListing
		walletDetails models.Wallet
		resp          models.GetUserWalletDetailsResponse
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
		log.Printf("Error during GetUserWalletDetails DB - get wallet info query: %v", err.Error())
		return
	}

	if err := models.DB.Raw(utils.GetWalletTransactionQuery(), input.GetUserID(), input.GetUserID(), input.GetUserID()).Scan(&transactions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserWalletDetails - get wallet transactions DB query: %v", err.Error())
		return
	}

	resp.WalletInfo = walletDetails
	resp.TransactionsCount = utils.Uint32(uint32(len(transactions)))
	resp.Transactions = transactions

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieved user wallet details and transactions."), "Data": resp})

	log.Println("Successful: GetUserWalletDetails.")
}
