package listings

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/models"

	"github.com/gin-gonic/gin"
)

func DeleteListing(c *gin.Context) {
	// Get model if exist
	var (
		deleteListing models.Listing
		input         models.DeleteSingleListingRequest
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

	if err := models.DB.Where("item_id = ?", input.ItemID).First(&deleteListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		return
	}

	if err := models.DB.Delete(&deleteListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
