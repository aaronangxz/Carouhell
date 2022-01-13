package listings

import (
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func DeleteListing(c *gin.Context) {
	var (
		listingToBeDeleted models.Listing
		input              models.DeleteSingleListingRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetItemID() <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be > 0.")})
		return
	}

	if err := models.DB.Raw("SELECT * FROM listing_tab WHERE l_item_id = ?", input.ItemID).Scan(&listingToBeDeleted).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		return
	}

	//invalidate cache of seller profile
	//must be done first, otherwise itemid cannot be found after deletion
	if err := utils.InvalidateSellerCacheUsingItemID(utils.GetUserDetailsCacheKey, input.GetItemID()); err != nil {
		log.Printf("Error during DeleteListing InvalidateSellerCacheUsingItemID: %v", err.Error())
	}

	//archive listing
	if err := models.DB.Table("listing_archive_tab").Create(&listingToBeDeleted).Error; err != nil {
		log.Printf("Error during listing archive DB query: %v: %v", listingToBeDeleted.GetLItemID(), err.Error())
	}
	log.Printf("Successfully archived listing: %v", listingToBeDeleted.GetLItemID())

	//soft delete
	if err := models.DB.Exec("UPDATE listing_tab SET item_status = 3 WHERE l_item_id = ?", input.ItemID).Error; err != nil {
		log.Printf("Error during listing delete DB query: %v: %v", input.GetItemID(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}
	log.Printf("Successfully deleted listing: %v", input.GetItemID())

	//invalidate redis
	if err := utils.InvalidateCache(utils.GetSingleListingByUserIDCacheKey, input.GetItemID()); err != nil {
		log.Printf("Error during InvalidateCache: %v", err.Error())
	}

	log.Printf("Success: DeleteListing: %v", input.GetItemID())
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully delete listing.")})
}
