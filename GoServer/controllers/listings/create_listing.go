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

func ValidateCreateListingRequest(c *gin.Context, input *models.CreateListingRequest) {
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
		if input.ItemQuantity == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemQuantity) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity must be uint type.")})
			return
		}
		if input.ItemDescription == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_description cannot be empty.")})
			return
		}
		if !utils.ValidateString(input.ItemDescription) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_description must be string type.")})
			return
		}
		if input.ItemShippingInfo == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_shipping_info cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemShippingInfo) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_shipping_info must be uint type.")})
			return
		}
		if input.ItemPaymentInfo == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_payment_info cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemPaymentInfo) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_payment_info must be uint type.")})
			return
		}
		if input.ItemLocation == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location cannot be empty.")})
			return
		}
		if !utils.ValidateString(input.ItemLocation) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location must be string type.")})
			return
		}
		if input.ItemCategory == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be uint type.")})
			return
		}
		if input.SellerID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("seller_id cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.SellerID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("seller_id must be uint type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}
}

func CreateListing(c *gin.Context) {
	// Validate input
	var (
		input models.CreateListingRequest
		hold  models.Account
	)
	ValidateCreateListingRequest(c, &input)

	if input.GetItemName() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot be empty.")})
		return
	}

	if !utils.ValidateMaxStringLength(input.GetItemName()) {
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

	if input.GetItemQuantity() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity must be > 0.")})
		return
	}

	//allow blank
	if input.GetItemDescription() == "" {
		input.SetItemDescription("This item has no description.")
	}

	if input.GetItemLocation() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location cannot be empty.")})
		return
	}

	//check if seller exists
	if err := models.DB.Raw("SELECT * FROM acc_tab WHERE user_id = ?", input.SellerID).Scan(&hold).Error; err != nil {
		if hold.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("seller_id does not exist.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	listings := models.Listing{
		ItemName:              input.ItemName,
		ItemPrice:             input.ItemPrice,
		ItemQuantity:          input.ItemQuantity,
		ItemPurchasedQuantity: utils.Uint32(0),
		ItemDescription:       input.ItemDescription,
		ItemShippingInfo:      input.ItemShippingInfo,
		ItemPaymentInfo:       input.ItemPaymentInfo,
		ItemLocation:          input.ItemLocation,
		ItemStatus:            utils.Uint32(constant.ITEM_STATUS_NORMAL),
		ItemCategory:          input.ItemCategory,
		ItemImage:             input.ItemImage,
		SellerID:              input.SellerID,
		ListingCtime:          utils.Int64(time.Now().Unix()),
	}

	if err := models.DB.Table("listing_tab").Create(&listings).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
