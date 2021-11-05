package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	//Redis key
	key := fmt.Sprint("get_user_wallet_details:", input.GetUserID())
	//check redis
	val, err := models.RedisClient.Get(models.Ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("No result of %v in Redis, reading from DB", key)
		} else {
			log.Printf("Error while reading from redis: %v", err.Error())
		}
	} else {
		redisResp := models.GetUserWalletDetailsResponse{}
		err := json.Unmarshal([]byte(val), &redisResp)
		if err != nil {
			log.Printf("Fail to unmarshal Redis value of key %v : %v, reading from DB", key, err)
		}
		c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieved user wallet details and transactions."), "Data": redisResp})
		log.Printf("Successful: GetUserWalletDetails: %v - Cached", input.GetUserID())
		return
	}

	if err := models.DB.Raw("SELECT * FROM wallet_tab WHERE w_user_id = ?", input.UserID).Scan(&walletDetails).Error; err != nil {
		//check if user exists
		if walletDetails.WUserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("Wallet is not activated.")})
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

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v\n", err.Error())
	}

	//Retrieved new data, set to Redis
	expiry := 60 * time.Second
	if err := models.RedisClient.Set(models.Ctx, key, data, expiry).Err(); err != nil {
		log.Printf("Error while writing to redis: %v", err.Error())
	} else {
		log.Printf("Written %v to redis", key)
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieved user wallet details and transactions."), "Data": resp})
	log.Printf("Successful: GetUserWalletDetails: %v - DB", input.GetUserID())
}
