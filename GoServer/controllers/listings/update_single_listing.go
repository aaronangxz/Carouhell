package listings

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func UpdateSingleListing(c *gin.Context) {
	var (
		originalListing models.Listing
		input           models.UpdateListingRequest
	)

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be uint type.")})
			return
		}
		//other 3 fields only check when it is not nil because we allow it to be empty
		if input.ItemName != nil && !utils.ValidateString(input.ItemName) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name must be string type.")})
			return
		}
		if input.ItemPrice != nil && !utils.ValidateUint(input.ItemPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be uint type.")})
			return
		}
		if input.ItemImg != nil && !utils.ValidateString(input.ItemImg) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_img must be string type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetItemID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be > 0.")})
		return
	}

	if input.ItemName != nil && input.GetItemName() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot be empty. Set to null if no changes needed.")})
		return
	}

	if input.ItemPrice != nil && input.GetItemPrice() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be > 0. Set to null if no changes needed.")})
		return
	}

	if input.ItemImg != nil && input.GetItemImg() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_img cannot be empty. Set to null if no changes needed.")})
		return
	}

	//Check if record exists
	if err := models.DB.Raw("SELECT * FROM listing_tab WHERE item_id = ?", input.ItemID).Scan(&originalListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		return
	}

	//If request fields are empty, we dont want to override empty fields into DB
	if input.ItemName == nil {
		input.ItemName = &originalListing.ItemName
	}
	if input.ItemPrice == nil {
		input.ItemPrice = &originalListing.ItemPrice
	}
	if input.ItemImg == nil {
		input.ItemImg = &originalListing.ItemImage
	}

	//If all good, proceed to update
	if err := models.DB.Exec("UPDATE listing_tab SET item_name = ?, item_price = ?, item_image = ? WHERE item_id = ?", input.ItemName, input.ItemPrice, input.ItemImg, input.ItemID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
