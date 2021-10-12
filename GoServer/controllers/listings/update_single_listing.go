package listings

import (
	"net/http"
	"reflect"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func ValidateUpdateSingleListingRequest(c *gin.Context, input *models.UpdateListingRequest) {
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
		//other fields only check when it is not nil because we allow it to be empty
		if input.ItemName != nil && !utils.ValidateString(input.ItemName) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name must be string type.")})
			return
		}
		if input.ItemPrice != nil && !utils.ValidateUint(input.ItemPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be uint type.")})
			return
		}
		if input.ItemQuantity != nil && !utils.ValidateUint(input.ItemQuantity) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity must be uint type.")})
			return
		}
		if input.ItemDescription != nil && !utils.ValidateString(input.ItemDescription) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_description must be string type.")})
			return
		}
		if input.ItemShippingInfo != nil && !utils.ValidateUint(input.ItemShippingInfo) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_shippinginfo must be uint type.")})
			return
		}
		if input.ItemPaymentInfo != nil && !utils.ValidateUint(input.ItemPaymentInfo) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_paymentinfo must be uint type.")})
			return
		}
		if input.ItemLocation != nil && !utils.ValidateString(input.ItemLocation) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location must be string type.")})
			return
		}
		if input.ItemCategory != nil && !utils.ValidateUint(input.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be uint type.")})
			return
		}
		if input.ItemImage != nil && !utils.ValidateString(input.ItemImage) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image must be string type.")})
			return
		}
		if input.SellerID != nil && !utils.ValidateUint(input.SellerID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("seller_id must be uint type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}
}

func ValidateUpdateSingleListingInput(c *gin.Context, input *models.UpdateListingRequest) {
	//Allow:
	//quantity to be 0
	//description to be blank
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

	//Not in enum
	if input.ItemShippingInfo != nil && input.GetShippingInfo() >= 3 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_shippinginfo does not exist.")})
		return
	}

	//Not in enum
	if input.ItemPaymentInfo != nil && input.GetPaymentInfo() >= 5 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_paymentinfo does not exist.")})
		return
	}

	//Not in enum
	if input.ItemCategory != nil && input.GetItemCategory() >= 2 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category does not exist.")})
		return
	}

	if input.ItemLocation != nil && input.GetItemLocation() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location cannot be empty. Set to null if no changes needed.")})
		return
	}

	if input.ItemImage != nil && input.GetItemImage() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image cannot be empty. Set to null if no changes needed.")})
		return
	}
}

func UpdateSingleListing(c *gin.Context) {
	var (
		originalListing models.Listing
		input           models.UpdateListingRequest
	)

	//Validate data type
	ValidateUpdateSingleListingRequest(c, &input)

	//Validate user existence
	utils.ValidateUserID(c, input.GetSellerID())

	//Validate input values
	ValidateUpdateSingleListingInput(c, &input)

	//Check if record exists
	//If yes, retrieve and store original records
	if err := models.DB.Raw("SELECT * FROM listing_tab WHERE item_id = ?", input.ItemID).Scan(&originalListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		return
	}

	//If request fields are empty, we dont want to override empty fields into DB
	inputValue := reflect.ValueOf(&input)
	originalValue := reflect.ValueOf(&originalListing)

	inputElement := inputValue.Elem()
	originalElement := originalValue.Elem()

	//Fill up nil fields with original values
	for i := 0; i < inputElement.NumField(); i++ {
		if inputElement.Field(i).IsNil() {
			inputElement = originalElement.Field(i)
		}
	}

	//If all good, proceed to update
	if err := models.DB.Exec("UPDATE listing_tab SET "+
		"item_name = ?, item_price = ?, item_quantity = ?,"+
		"item_description = ?, item_shippinginfo = ?, item_paymentinfo = ?,"+
		"item_location = ?, item_category = ?, item_image = ? WHERE item_id = ?",
		input.GetItemName(), input.GetItemPrice(), input.GetItemQuantity(), input.GetItemDescription(), input.GetShippingInfo(), input.GetPaymentInfo(),
		input.GetItemLocation(), input.GetItemCategory(), input.GetItemImage(), input.GetItemID()).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
