package listings

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetUserListings(c *gin.Context) {
	var (
		userListings   []models.GetUserListingsResponse
		input          models.GetUserListingsRequest
		extraCondition = ""
		orderCondition = " ORDER BY listing_ctime DESC"
		whereCondition = " AND l.l_seller_id = ?"
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			log.Println("user_id cannot be empty.")
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		log.Printf("JSON error: %v\n", err.Error())
		return
	}

	if input.GetUserID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be > 0.")})
		log.Println("user_id must be > 0.")
		return
	}

	if input.Limit == nil {
		input.Limit = models.SetDefaultNotificationResponseLimit()
		log.Println("applied DefaultNotificationResponseLimit.")
	}

	if utils.ValidateLimitMax(input.GetLimit(), models.MaxListingsResponseSize) {
		c.JSON(http.StatusBadRequest, gin.H{"RespMeta": models.NewParamErrorsResponse("limit exceeds max listing response size.")})
		log.Printf("limit exceeds max listing response size: %v\n", input.GetLimit())
		return
	}

	switch input.GetLimit() {
	case 0:
		extraCondition = ""
	default:
		extraCondition = " LIMIT " + fmt.Sprint(input.GetLimit())
	}

	//also return deleted and sold items
	query := utils.GetListingQueryWithCustomCondition() + whereCondition + orderCondition + extraCondition
	log.Println(query)

	result := models.DB.Raw(query, input.GetUserID()).Scan(&userListings)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserListings DB query: %v\n", err.Error())
		return
	}

	data, err := json.Marshal(userListings)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v\n", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetUserListingsResult(userListings), "Data": userListings})
	log.Printf("Successful: GetListingsUsingFilters. rows: %v\n", result.RowsAffected)
	log.Printf("Result: %s\n", data)
}
