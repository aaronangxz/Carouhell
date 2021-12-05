package listings

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func GetRecommendedListingsByItemId(c *gin.Context) {
	var (
		input models.GetRecommendedListingsByItemIdRequest
		hold  models.Listing
		resp  []models.GetRecommendedListingsByItemIdResponse
	)
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("UserId cannot be empty.")})
			log.Println("user_id cannot be empty.")
			return
		}
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("ItemId cannot be empty.")})
			log.Println("item_id cannot be empty.")
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		log.Printf("JSON error: %v\n", err.Error())
		return
	}

	if input.GetUserID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("UserId must be > 0.")})
		log.Println("user_id must be > 0.")
		return
	}

	if input.GetItemID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("ItemId must be > 0.")})
		log.Println("user_id must be > 0.")
		return
	}

	//get item name
	itemNameQuery := fmt.Sprintf("SELECT * FROM listing_tab WHERE l_item_id = %v", input.GetItemID())
	log.Println(itemNameQuery)

	itemNameResult := models.DB.Raw(itemNameQuery).Scan(&hold)

	if err := itemNameResult.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetRecommendedListingsByItemId DB query: %v\n", err.Error())
		return
	}

	//sanitize item name
	sanitizedItemName := utils.RemoveSpecialChar(hold.GetItemName())

	//get recommendations
	recommendationQuery := fmt.Sprintln(utils.GetRecommendedListingsByItemIdQuery(input.GetItemID(), input.GetUserID(), sanitizedItemName, hold.GetItemCategory()))
	log.Println(recommendationQuery)

	recommendationResult := models.DB.Raw(recommendationQuery).Scan(&resp)

	if err := recommendationResult.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetRecommendedListingsByItemId DB query: %v\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse(), "Data": resp})
	log.Println("Successful: GetRecommendedListingsByItemId - DB")
	log.Printf("Result: %v\n", resp)
}
