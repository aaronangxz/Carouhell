package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
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
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Username cannot be empty.")})
			errormsg := "item_image cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.UserName) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Username must be string type.")})
			errormsg := fmt.Sprintf("item_image must be string type. input: %v", input.GetUserName())
			return errors.New(errormsg)
		}
		if input.UserEmail == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Email cannot be empty.")})
			errormsg := "user_email cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.UserEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Email must be string type.")})
			errormsg := fmt.Sprintf("user_email must be string type. input: %v", input.GetUserName())
			return errors.New(errormsg)
		}
		if input.UserPassword == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Password cannot be empty.")})
			errormsg := "user_password cannot be empty"
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
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Username cannot be empty.")})
		return errors.New("user_name cannot be empty")
	}
	if !utils.ValidateMaxStringLength(input.GetUserName()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Username cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		errormsg := fmt.Sprintf("user_name length cannot exceed %v. input :%v", models.MaxStringLength, len(input.GetUserName()))
		return errors.New(errormsg)
	}
	if utils.IsContainsSpecialChar(input.GetUserName()) || utils.IsContainsSpace(input.GetUserName()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Username can only contain alphanumeric characters.")})
		errormsg := fmt.Sprintf("user_name can only contains alphanumeric characters. input :%v", input.GetUserName())
		return errors.New(errormsg)
	}
	if input.GetUserEmail() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Email cannot be empty.")})
		return errors.New("user_email cannot be empty")
	}
	if !utils.ValidateMaxStringLength(input.GetUserEmail()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Email cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		errormsg := fmt.Sprintf("user_email length cannot exceed %v. input :%v", models.MaxStringLength, len(input.GetUserEmail()))
		return errors.New(errormsg)
	}
	if !utils.IsContainsAtSign(input.GetUserEmail()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Email format is invalid.")})
		errormsg := fmt.Sprintf("user_email format is invalid. input :%v", input.GetUserEmail())
		return errors.New(errormsg)
	}
	if input.GetUserPassword() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Password cannot be empty.")})
		return errors.New("user_password cannot be empty")
	}
	if len(input.GetUserPassword()) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Password cannot be shorter than 6 chars.")})
		return errors.New("user_password cannot be shorter than 6 chars")
	}
	return nil
}

func CreateAccount(c *gin.Context) {
	var (
		input  models.CreateAccountRequest
		spResp models.CreateAccountSPResponse
		hold   models.Account
	)

	if err := ValidateCreateAccountRequest(c, &input); err != nil {
		log.Printf("Error during ValidateCreateAccountRequest: %v", err.Error())
		return
	}

	if err := ValidateCreateAccountInput(c, &input); err != nil {
		log.Printf("Error during ValidateCreateAccountInput: %v", err.Error())
		return
	}

	//hash password
	hashedPassword, err := HashSecret(input.GetUserPassword())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Error during account creation. Please try again.")})
		log.Printf("Error during HashSecret - password: %v", err)
		return
	}

	//toggle to use stored procedures
	if os.Getenv("USE_STORED_PROCEDURE") == "TRUE" {
		query := fmt.Sprintf("CALL heroku_bdc39d4687a85d4.create_user('%v', '%v', '%v',@status);",
			input.GetUserName(), input.GetUserEmail(), hashedPassword)
		log.Printf("Executing SP: %v\n", query)
		result := models.DB.Exec(query)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(result.Error)})
			log.Printf("Error during CreateAccount - create_user SP DB query: %v", result.Error.Error())
			return
		}

		//check SP run status
		spResult := models.DB.Raw("SELECT @status AS status").Scan(&spResp)
		if spResult.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(spResult.Error)})
			log.Printf("Error during CreateAccount - check_sp_status DB query: %v", result.Error.Error())
			return
		}

		spData, err := json.Marshal(spResp)
		if err != nil {
			log.Printf("Failed to marshal JSON results: %v", err.Error())
		}
		log.Printf("check_sp_status. Data: %s\n", spData)

		if spResp.Status == nil {
			c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewUnknownErrorResponse()})
			log.Printf("Response from SP status is nil. Check if SP ran successfully\n")
			return
		} else {
			if spResp.GetStatus() == -1 {
				c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Username or email already exists.")})
				log.Printf("user already exists: %v / %v\n", input.GetUserName(), input.GetUserEmail())
				return
			}
			if spResp.GetStatus() == 0 {
				c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully created account and wallet.")})
				log.Printf("Successful: CreateAccount - via SP\n")
				return
			}
		}
	}

	//normal queries
	//check if user name / user email exists
	query := fmt.Sprintf("SELECT * FROM acc_tab WHERE user_name = '%v' OR user_email = '%v'", input.GetUserName(), input.GetUserEmail())
	result := models.DB.Raw(query).Scan(&hold)
	log.Println(query)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(result.Error)})
			log.Printf("Error during CreateAccount - check_user_exists DB query: %v\n", result.Error.Error())
			return
		}
		log.Printf("user_name does not exist: %v\n", input.GetUserName())
	}

	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Username or email already exists. Please log in instead.")})
		log.Printf("user already exists: %v / %v", input.GetUserEmail(), input.GetUserName())
		return
	}

	account := models.Account{
		UserName:      input.UserName,
		UserEmail:     input.UserEmail,
		UserCtime:     utils.Int64(time.Now().Unix()),
		UserStatus:    utils.Uint32(constant.ACC_STATUS_ACTIVE),
		UserLastLogin: utils.Int64(time.Now().Unix()),
	}

	//write to acc_tab
	if err := models.DB.Table("acc_tab").Create(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during CreateAccount - acc_tab DB query: %v", err.Error())
		return
	}

	credentials := models.AccountCredentials{
		CUserID:      account.AUserID,
		UserPassword: utils.String(hashedPassword),
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
		WUserID:       account.AUserID,
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

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully created account and wallet.")})

	data, err := json.Marshal(account)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v", err.Error())
	}
	log.Printf("Successful: CreateAccount. Data: %s", data)
}
