package listings

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetUserLikedListings(c *gin.Context) {
	var (
		userLikedListings []models.GetUserLikedListingsResponse
		input             models.GetUserLikedListingsRequest
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

	query := fmt.Sprintf(utils.GetListingQueryWithCustomCondition()+" AND l.item_id IN"+
		" (SELECT r.item_id FROM listing_reactions_tab r"+
		" WHERE r.user_id = %v AND r.reaction_type = %v)"+
		" GROUP BY l.item_id ORDER BY listing_reactions_tab.ctime DESC", input.GetUserID(), constant.LISTING_REACTION_TYPE_LIKE)
	log.Println(query)

	result := models.DB.Raw(query).Scan(&userLikedListings)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserLikedListings DB query: %v\n", err.Error())
		return
	}

	data, err := json.Marshal(userLikedListings)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v\n", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetUserLikedListingsResult(userLikedListings), "Data": userLikedListings})
	log.Printf("Successful: GetUserLikedListings. rows: %v\n", result.RowsAffected)
	log.Printf("Result: %s\n", data)
}
