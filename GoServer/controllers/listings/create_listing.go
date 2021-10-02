package listings

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func CreateListing(c *gin.Context) {
	// Validate input
	var input models.CreateListingRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemName == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot be empty.")})
			return
		}
		if !utils.ValidateString(input.ItemName) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name must be string type.")})
			return
		}
		if input.ItemPrice == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be uint type.")})
			return
		}
		if input.ItemImage == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_img cannot be empty.")})
			return
		}
		if !utils.ValidateString(input.ItemImage) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image must be string type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetItemName() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot be empty.")})
		return
	}

	if len(input.GetItemName()) > int(models.MaxStringLength) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		return
	}

	if input.GetItemPrice() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be > 0.")})
		return
	}

	if input.GetItemImage() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_img cannot be empty.")})
		return
	}

	listings := models.Listing{
		ItemName:              input.GetItemName(),
		ItemPrice:             input.GetItemPrice(),
		ItemImage:             input.GetItemImage(),
		ItemCreationTime:      time.Now().Unix(),
		ItemQuantity:          input.GetItemQuantity(),
		ItemPurchasedQuantity: 0,
		ItemDescription:       input.GetItemDescription(),
		ItemShippingInfo:      input.GetShippingInfo(),
		ItemPaymentInfo:       input.GetPaymentInfo(),
		ItemLocation:          input.GetItemLocation(),
		ItemCategory:          input.GetItemCategory(),
		ItemStatus:            constant.ITEM_STATUS_NORMAL,
	}

	if err := models.DB.Create(&listings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
