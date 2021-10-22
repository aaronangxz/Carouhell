package listings

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func ValidateUpdateSingleListingRequest(c *gin.Context, input *models.UpdateListingRequest) error {
	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			errormsg := "item_id is empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.ItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be uint type.")})
			errormsg := fmt.Sprintf("item_id must be uint. input: %v", input.GetItemID())
			return errors.New(errormsg)
		}
		//other fields only check when it is not nil because we allow it to be empty
		if input.ItemName != nil && !utils.ValidateString(input.ItemName) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name must be string type.")})
			errormsg := fmt.Sprintf("item_name must be string. input: %v", input.GetItemName())
			return errors.New(errormsg)
		}
		if input.ItemPrice != nil && !utils.ValidateUint(input.ItemPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be uint type.")})
			errormsg := fmt.Sprintf("item_price must be uint. input: %v", input.GetItemPrice())
			return errors.New(errormsg)
		}
		if input.ItemQuantity != nil && !utils.ValidateUint(input.ItemQuantity) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity must be uint type.")})
			errormsg := fmt.Sprintf("item_quantity must be uint. input: %v", input.GetItemQuantity())
			return errors.New(errormsg)
		}
		if input.ItemDescription != nil && !utils.ValidateString(input.ItemDescription) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_description must be string type.")})
			errormsg := fmt.Sprintf("item_description must be string. input: %v", input.GetItemDescription())
			return errors.New(errormsg)
		}
		if input.ItemShippingInfo != nil && !utils.ValidateUint(input.ItemShippingInfo) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_shippinginfo must be uint type.")})
			errormsg := fmt.Sprintf("item_shippinginfo must be uint. input: %v", input.GetShippingInfo())
			return errors.New(errormsg)
		}
		if input.ItemPaymentInfo != nil && !utils.ValidateUint(input.ItemPaymentInfo) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_paymentinfo must be uint type.")})
			errormsg := fmt.Sprintf("item_paymentinfo must be uint. input: %v", input.GetPaymentInfo())
			return errors.New(errormsg)
		}
		if input.ItemLocation != nil && !utils.ValidateString(input.ItemLocation) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location must be string type.")})
			errormsg := fmt.Sprintf("item_location must be string. input: %v", input.GetItemLocation())
			return errors.New(errormsg)
		}
		if input.ItemCategory != nil && !utils.ValidateUint(input.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be uint type.")})
			errormsg := fmt.Sprintf("item_category must be uint. input: %v", input.GetItemCategory())
			return errors.New(errormsg)
		}
		if input.ItemImage != nil && !utils.ValidateString(input.ItemImage) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image must be string type.")})
			errormsg := fmt.Sprintf("item_image must be string. input: %v", input.GetItemImage())
			return errors.New(errormsg)
		}
		if input.SellerID != nil && !utils.ValidateUint(input.SellerID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("seller_id must be uint type.")})
			errormsg := fmt.Sprintf("seller_id must be uint. input: %v", input.GetSellerID())
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func ValidateUpdateSingleListingInput(c *gin.Context, input *models.UpdateListingRequest) error {
	//Allow:
	//quantity to be 0
	//description to be blank
	if input.GetItemID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be > 0.")})
		errormsg := fmt.Sprintf("item_id must > 0. input: %v", input.GetItemID())
		return errors.New(errormsg)
	}

	if input.ItemName != nil && input.GetItemName() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot be empty. Set to null if no changes needed.")})
		errormsg := fmt.Sprintf("item_name cannot be empty. input: %v", input.GetItemName())
		return errors.New(errormsg)
	}

	if input.ItemPrice != nil && input.GetItemPrice() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be > 0. Set to null if no changes needed.")})
		errormsg := fmt.Sprintf("item_price must > 0. input: %v", input.GetItemPrice())
		return errors.New(errormsg)
	}

	//Not in enum
	if input.ItemShippingInfo != nil && constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_SHIPPING_TYPE, input.GetShippingInfo()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_shippinginfo does not exist.")})
		errormsg := fmt.Sprintf("item_shippinginfo does not exist. input: %v", input.GetShippingInfo())
		return errors.New(errormsg)
	}

	//Not in enum
	if input.ItemPaymentInfo != nil && constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_PAYMENT_TYPE, input.GetPaymentInfo()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_paymentinfo does not exist.")})
		errormsg := fmt.Sprintf("item_paymentinfo does not exist. input: %v", input.GetPaymentInfo())
		return errors.New(errormsg)
	}

	//Not in enum
	if input.ItemCategory != nil && constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_ITEM_CATEGORY, input.GetItemCategory()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category does not exist.")})
		errormsg := fmt.Sprintf("item_category does not exist. input: %v", input.GetItemCategory())
		return errors.New(errormsg)
	}

	if input.ItemLocation != nil && input.GetItemLocation() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location cannot be empty. Set to null if no changes needed.")})
		errormsg := fmt.Sprintf("item_location cannot be empty. input: %v", input.GetItemLocation())
		return errors.New(errormsg)
	}

	if input.ItemImage != nil && input.GetItemImage() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image cannot be empty. Set to null if no changes needed.")})
		errormsg := fmt.Sprintf("item_image cannot be empty. input: %v", input.GetItemImage())
		return errors.New(errormsg)
	}
	return nil
}

func UpdateSingleListing(c *gin.Context) {
	var (
		originalListing models.Listing
		input           models.UpdateListingRequest
	)

	//Validate data type
	if err := ValidateUpdateSingleListingRequest(c, &input); err != nil {
		log.Printf("Error during ValidateUpdateSingleListingRequest: %v", err.Error())
		return
	}

	//Validate user existence
	if err := utils.ValidateUserID(c, input.GetSellerID()); err != nil {
		log.Printf("Error ValidateUserID: %v", err.Error())
		return
	}

	//Validate input values
	if err := ValidateUpdateSingleListingInput(c, &input); err != nil {
		log.Printf("Error during ValidateUpdateSingleListingInput: %v", err.Error())
		return
	}

	//Check if record exists
	//If yes, retrieve and store original records
	if err := models.DB.Raw("SELECT * FROM listing_tab WHERE item_id = ?", input.ItemID).Scan(&originalListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("DB Error: %v", err.Error())
		return
	}

	if originalListing.GetSellerID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		log.Printf("Record not found: user_id: %v", input.GetItemID())
		return
	}

	// //If request fields are empty, we dont want to override empty fields into DB
	// inputValue := reflect.ValueOf(&input)
	// originalValue := reflect.ValueOf(&originalListing)

	// inputElement := inputValue.Elem()
	// originalElement := originalValue.Elem()

	// //Fill up nil fields with original values
	// for i := 0; i < inputElement.NumField(); i++ {
	// 	if inputElement.Field(i).IsNil() {
	// 		inputElement = originalElement.Field(i)
	// 	}
	// }
	if input.ItemName == nil {
		input.ItemName = originalListing.ItemName
	}
	if input.ItemPrice == nil {
		input.ItemPrice = originalListing.ItemPrice
	}
	if input.ItemQuantity == nil {
		input.ItemQuantity = originalListing.ItemQuantity
	}
	if input.ItemDescription == nil {
		input.ItemDescription = originalListing.ItemDescription
	}
	if input.ItemShippingInfo == nil {
		input.ItemShippingInfo = originalListing.ItemShippingInfo
	}
	if input.ItemPaymentInfo == nil {
		input.ItemPaymentInfo = originalListing.ItemPaymentInfo
	}
	if input.ItemLocation == nil {
		input.ItemLocation = originalListing.ItemLocation
	}
	if input.ItemCategory == nil {
		input.ItemCategory = originalListing.ItemCategory
	}
	if input.ItemImage == nil {
		input.ItemImage = originalListing.ItemImage
	}

	//If all good, proceed to update
	if err := models.DB.Exec("UPDATE listing_tab SET "+
		"item_name = ?, item_price = ?, item_quantity = ?,"+
		"item_description = ?, item_shipping_info = ?, item_payment_info = ?,"+
		"item_location = ?, item_category = ?, item_image = ?, listing_mtime = ? WHERE item_id = ?",
		input.GetItemName(), input.GetItemPrice(), input.GetItemQuantity(), input.GetItemDescription(), input.GetShippingInfo(), input.GetPaymentInfo(),
		input.GetItemLocation(), input.GetItemCategory(), input.GetItemImage(), time.Now().Unix(), input.GetItemID()).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during DB query: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
	log.Printf("Successful: UpdateSingleListing.")
}
