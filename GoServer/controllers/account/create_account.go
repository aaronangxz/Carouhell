package account

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
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func HashSecret(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	log.Println("Successfully hashed secret")
	return string(bytes), nil
}

func ValidateCreateAccountRequest(c *gin.Context, input *models.CreateAccountRequest) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserName == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image cannot be empty.")})
			errormsg := "item_image cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.UserName) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image must be string type.")})
			errormsg := fmt.Sprintf("item_image must be string type. input: %v", input.GetUserName())
			return errors.New(errormsg)
		}
		if input.UserEmail == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_email cannot be empty.")})
			errormsg := "user_email cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.UserEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_email must be string type.")})
			errormsg := fmt.Sprintf("user_email must be string type. input: %v", input.GetUserName())
			return errors.New(errormsg)
		}
		if input.UserPassword == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_password cannot be empty.")})
			errormsg := "user_password cannot be empty"
			return errors.New(errormsg)
		}
		if input.UserSecurityQuestion == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price cannot be empty.")})
			errormsg := "user_security_question cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.UserSecurityQuestion) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be uint type.")})
			errormsg := fmt.Sprintf("user_security_question must be uint type. input: %v", input.GetUserSecurityQuestion())
			return errors.New(errormsg)
		}
		if input.UserSecurityAnswer == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_security_answer cannot be empty.")})
			errormsg := "user_security_answer cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.UserSecurityAnswer) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_security_answer must be string type.")})
			errormsg := fmt.Sprintf("user_security_answer must be string type. input: %v", input.GetUserSecurityAnswer())
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func ValidateCreateAccountInput(c *gin.Context, input *models.CreateAccountRequest) error {
	if input.GetUserName() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_name cannot be empty.")})
		return errors.New("user_name cannot be empty")
	}
	if !utils.ValidateMaxStringLength(input.GetUserName()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		errormsg := fmt.Sprintf("item_name length cannot exceed %v. input :%v", models.MaxStringLength, len(input.GetUserName()))
		return errors.New(errormsg)
	}
	if input.GetUserEmail() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_email cannot be empty.")})
		return errors.New("user_email cannot be empty")
	}
	if !utils.ValidateMaxStringLength(input.GetUserEmail()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_email cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		errormsg := fmt.Sprintf("user_email length cannot exceed %v. input :%v", models.MaxStringLength, len(input.GetUserEmail()))
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
	//check if exists
	if input.UserSecurityQuestion != nil && !constant.CheckAccConstant(constant.ACC_CREATION_SECURITY_QUESTION, input.GetUserSecurityQuestion()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("unknown item_category.")})
		errormsg := fmt.Sprintf("unknown user_security_question. input: %v", input.GetUserSecurityQuestion())
		return errors.New(errormsg)
	}
	if input.GetUserSecurityAnswer() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_security_answer cannot be empty.")})
		return errors.New("user_email cannot be empty")
	}
	if !utils.ValidateMaxStringLength(input.GetUserSecurityAnswer()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_security_answer cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		errormsg := fmt.Sprintf("user_security_answer length cannot exceed %v. input :%v", models.MaxStringLength, len(input.GetUserEmail()))
		return errors.New(errormsg)
	}
	return nil
}

func CreateAccount(c *gin.Context) {
	var (
		input models.CreateAccountRequest
		hold  models.Account
	)

	if err := ValidateCreateAccountRequest(c, &input); err != nil {
		log.Printf("Error during ValidateCreateAccountRequest: %v", err.Error())
		return
	}

	if err := ValidateCreateAccountInput(c, &input); err != nil {
		log.Printf("Error during ValidateCreateAccountInput: %v", err.Error())
		return
	}

	//check if user name / user email exists
	result := models.DB.Raw("SELECT * FROM acc_tab WHERE user_name = ? OR user_email = ?", input.UserName, input.UserEmail).Scan(&hold)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(result.Error)})
			log.Printf("Error during DB query: %v", result.Error.Error())
			return
		}
	}

	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user already exists.")})
		log.Printf("user already exists: %v / %v", input.GetUserEmail(), input.GetUserName())
		return
	}

	account := models.Account{
		UserName:      input.UserName,
		UserEmail:     input.UserEmail,
		UserCtime:     utils.Int64(time.Now().Unix()),
		UserStatus:    utils.Uint32(constant.ACC_STATUS_ACTIVE),
		UserImage:     nil,
		UserLastLogin: utils.Int64(time.Now().Unix()),
	}

	//write to acc_tab
	if err := models.DB.Table("acc_tab").Create(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during CreateAccount - acc_tab DB query: %v", err.Error())
		return
	}

	//hash password
	hashedPassword, err := HashSecret(input.GetUserPassword())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Error during account creation. Please try again.")})
		log.Printf("Error during HashSecret - password: %v", err)
		return
	}

	//hash password
	hashedSecurityAnswer, err := HashSecret(input.GetUserSecurityAnswer())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Error during account creation. Please try again.")})
		log.Printf("Error during HashSecret - security question: %v", err)
		return
	}

	credentials := models.AccountCredentials{
		CUserID:              account.AUserID,
		UserPassword:         utils.String(hashedPassword),
		UserSecurityQuestion: input.UserSecurityQuestion,
		UserSecurityAnswer:   utils.String(hashedSecurityAnswer),
	}

	//write to acc_credentials_tab
	if err := models.DB.Table("acc_credentials_tab").Create(&credentials).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during CreateAccount - acc_credentials_tab DB query: %v", err.Error())

		//rollback acc_tab if fail
		if errRollback := models.DB.Table("acc_tab").Delete(&account).Error; errRollback != nil {
			log.Printf("Error during CreateAccount - acc_tab roll back: %v", err.Error())
		} else {
			log.Print("rollback acc_tab successful")
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	//create wallet
	wallet := models.Wallet{
		WalletID:      account.AUserID,
		WalletBalance: utils.Uint32(0),
		WalletStatus:  utils.Uint32(constant.WALLET_STATUS_ACTIVE),
		LastTopUp:     nil,
		LastUsed:      utils.Int64(time.Now().Unix()),
	}

	if err := models.DB.Table("wallet_tab").Create(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during CreateAccount - wallet_tab DB query: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully create account and wallet.")})

	data, err := json.Marshal(account)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v", err.Error())
	}
	log.Printf("Successful: CreateAccount. Data: %s", data)
}
