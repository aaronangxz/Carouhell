package account

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ValidateAuthenticateUser(c *gin.Context, input *models.AuthenticateUser) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserName == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_name cannot be empty.")})
			errormsg := "user_name cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.UserName) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_name must be string type.")})
			errormsg := fmt.Sprintf("user_name must be string type. input: %v\n", input.GetUserName())
			return errors.New(errormsg)
		}
		if input.UserPassword == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_password cannot be empty.")})
			errormsg := "user_password cannot be empty"
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprintf("JSON error: %v\n", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func ValidateAuthenticateUserInput(c *gin.Context, input *models.AuthenticateUser) error {
	if input.GetUserName() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_name cannot be empty.")})
		return errors.New("user_name cannot be empty")
	}
	if !utils.ValidateMaxStringLength(input.GetUserName()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		errormsg := fmt.Sprintf("item_name length cannot exceed %v. input :%v\n", models.MaxStringLength, len(input.GetUserName()))
		return errors.New(errormsg)
	}
	if input.GetUserPassword() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_password cannot be empty.")})
		return errors.New("user_password cannot be empty")
	}
	if len(input.GetUserPassword()) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_password cannot be shorter than 6 chars.")})
		return errors.New("user_password cannot be shorter than 6 chars")
	}
	return nil
}

func AuthenticateUser(c *gin.Context) {
	var (
		input           models.AuthenticateUser
		resp            models.AuthenticateUserResponse
		hold            models.Account
		holdCredentials models.AccountCredentials
	)

	if err := ValidateAuthenticateUser(c, &input); err != nil {
		log.Printf("Error during ValidateAuthenticateUser: %v\n", err.Error())
		return
	}

	if err := ValidateAuthenticateUserInput(c, &input); err != nil {
		log.Printf("Error during ValidateAuthenticateUserInput: %v\n", err.Error())
		return
	}

	//check if user exists
	result := models.DB.Raw("SELECT * FROM acc_tab WHERE user_name = ?", input.UserName).Scan(&hold)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(result.Error)})
			log.Printf("Error during check_user_exists DB query: %v\n", result.Error.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_name does not exist.")})
			log.Printf("user_name does not exist: %v\n", input.GetUserName())
			return
		}
	}

	//retrieve credentials
	resultCredentials := models.DB.Raw("SELECT * FROM acc_credentials_tab WHERE user_id = ?", hold.GetUserID()).Scan(&holdCredentials)

	if resultCredentials.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(resultCredentials.Error)})
			log.Printf("Error during retrieve_credentials DB query: %v\n", resultCredentials.Error.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id does not exist.")})
			log.Printf("user_id does not exist: %v\n", hold.GetUserID())
			return
		}
	}

	//check password matches input
	if input.GetUserPassword() != holdCredentials.GetUserPassword() {
		c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewParamErrorsResponse("Incorrect password.")})
		log.Println("Incorrect password.")
		return
	}

	resp.UserID = hold.GetUserID()

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully logged in."), "Data": resp})
	log.Printf("Successful: AuthenticateUser. user_id: %v", hold.GetUserID())
}
