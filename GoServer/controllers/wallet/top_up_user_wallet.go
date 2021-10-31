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

func ValidateTopUpUserWalletRequest(c *gin.Context, input *models.TopUpUserWalletRequest) error {
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
		if input.Amount == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("amount cannot be empty.")})
			errormsg := "amount cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.Amount) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("amount must be uint type.")})
			errormsg := fmt.Sprintf("amount must be uint type. input: %v", input.GetAmount())
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func ValidateTopUpUserWalletInput(c *gin.Context, input *models.TopUpUserWalletRequest) error {
	if input.GetAmount() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Amount cannot be 0.")})
		errormsg := "amount cannot be 0"
		return errors.New(errormsg)
	}
	if input.GetAmount() < 500 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Minimum top up amount is $5.")})
		errormsg := "amount cannot be < $5"
		return errors.New(errormsg)
	}
	return nil
}

func TopUpUserWallet(c *gin.Context) {
	var (
		input models.TopUpUserWalletRequest
		resp  models.TopUpUserWalletResponse
	)

	if err := ValidateTopUpUserWalletRequest(c, &input); err != nil {
		log.Printf("Error during ValidateTopUpUserWalletRequest: %v", err.Error())
		return
	}

	if err := ValidateTopUpUserWalletInput(c, &input); err != nil {
		log.Printf("Error during ValidateTopUpUserWalletInput: %v", err.Error())
		return
	}

	if err := utils.ValidateUserID(c, input.GetUserID()); err != nil {
		log.Printf("Error during ValidateTopUpUserWalletRequest - ValidateUserID: %v", err.Error())
		return
	}

	updatedWalletBalance, err := utils.StartWalletTopUpTx(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Error during wallet top up.")})
		log.Printf("Error during StartWalletTopUpTx: %v", err.Error())
		return
	}

	resp.WalletBalance = updatedWalletBalance

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully top up wallet."), "Data": resp})

	log.Printf("Successful: TopUpUserWallet. Updated balance: %v", resp)
}
