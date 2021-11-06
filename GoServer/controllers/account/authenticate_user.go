package account

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/auth"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

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
			log.Printf("Error during AuthenticateUser - check_user_exists DB query: %v\n", result.Error.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_name does not exist.")})
			log.Printf("user_name does not exist: %v\n", input.GetUserName())
			return
		}
	}

	//retrieve credentials
	resultCredentials := models.DB.Raw("SELECT * FROM acc_credentials_tab WHERE c_user_id = ?", hold.GetUserID()).Scan(&holdCredentials)

	if resultCredentials.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(resultCredentials.Error)})
			log.Printf("Error during AuthenticateUser - retrieve_credentials DB query: %v\n", resultCredentials.Error.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id does not exist.")})
			log.Printf("user_id does not exist: %v\n", hold.GetUserID())
			return
		}
	}

	//check password matches input
	isMatch, err := CheckPasswordHash(input.GetUserPassword(), holdCredentials.GetUserPassword())
	if !isMatch || err != nil {
		c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewParamErrorsResponse("Incorrect password.")})
		log.Println("Incorrect password.")
		return
	}
	log.Println("Password match.")

	updateLastLoginQuery := fmt.Sprintf("UPDATE acc_tab SET user_last_login = %v WHERE a_user_id = %v", time.Now().Unix(), hold.GetUserID())
	updateLastLoginresult := models.DB.Exec(updateLastLoginQuery)
	if updateLastLoginresult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(updateLastLoginresult.Error)})
		log.Printf("Error during AuthenticateUser - update user_last_login DB query: %v\n", updateLastLoginresult.Error.Error())
		return
	}

	resp.UserID = hold.GetUserID()

	token, err := auth.CreateToken(hold.GetUserID())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Fail to generate token.")})
		log.Printf("Fail to generate token: %v\n", err)
		return
	}

	saveErr := auth.CreateAuth(hold.GetUserID(), token)
	if saveErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Fail to generate auth.")})
		log.Printf("Fail to generate auth: %v\n", err)
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully authenticated user."), "Data": resp, "Token": tokens})
	log.Printf("Successful: AuthenticateUser. user_id: %v", hold.GetUserID())
}
