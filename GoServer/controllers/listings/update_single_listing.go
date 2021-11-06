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
		if input.LItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			errormsg := "item_id is empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.LItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be uint type.")})
			errormsg := fmt.Sprintf("item_id must be uint. input: %v", input.GetLItemID())
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
		if input.ItemLocation != nil && !utils.ValidateUint(input.ItemLocation) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location must be uint type.")})
			errormsg := fmt.Sprintf("item_location must be uint. input: %v", input.GetItemLocation())
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
		if input.LSellerID != nil && !utils.ValidateUint(input.LSellerID) {
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
	if input.GetLItemID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be > 0.")})
		errormsg := fmt.Sprintf("item_id must > 0. input: %v", input.GetLItemID())
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
	if input.ItemCategory != nil && !constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_ITEM_CATEGORY, input.GetItemCategory()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category does not exist.")})
		errormsg := fmt.Sprintf("item_category does not exist. input: %v", input.GetItemCategory())
		return errors.New(errormsg)
	}

	if input.ItemLocation != nil && !constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_LOCATION, input.GetItemLocation()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location does not exist.")})
		errormsg := fmt.Sprintf("item_location does not exist. input: %v", input.GetItemLocation())
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

	//Check if record exists and is not sold / deleted
	//If yes, retrieve and store original records
	query := fmt.Sprintf("SELECT * FROM listing_tab WHERE l_item_id = %v AND item_status = %v", input.GetLItemID(), constant.ITEM_STATUS_NORMAL)
	log.Println(query)
	if err := models.DB.Raw(query).Scan(&originalListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("DB Error: %v", err.Error())
		return
	}

	if originalListing.GetLItemID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		log.Printf("item not found: item_id: %v", input.GetLItemID())
		return
	} else if originalListing.GetItemStatus() != constant.ITEM_STATUS_NORMAL {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		log.Printf("item is sold. item_id: %v", input.GetLItemID())
		return
	}

	// //If request fields are empty, we dont want to override empty fields into DB
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
	if input.ItemLocation == nil {
		input.ItemLocation = originalListing.ItemLocation
	}
	if input.ItemCategory == nil {
		input.ItemCategory = originalListing.ItemCategory
	}

	//upload image if any
	if input.ItemImage != nil {
		//upload image
		imageURL, err := utils.UploadBase64Image(originalListing.GetLItemID(), input.GetItemImage())
		input.ItemImage = &imageURL
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Failed to upload image. Listing not updated.")})
			log.Printf("Error during image upload: %v", err)
			return
		}
	}

	//If all good, proceed to update
	query = fmt.Sprintf("UPDATE listing_tab SET"+
		" item_name = \"%v\", item_price = %v, item_quantity = %v,"+
		" item_description = \"%v\", item_location = %v, item_category = %v, listing_mtime = %v,"+
		" item_status = CASE WHEN item_quantity = 0 THEN 2 ELSE item_status END,"+
		" item_status = CASE WHEN item_quantity > 0 THEN 1 ELSE item_status END WHERE l_item_id = %v",
		input.GetItemName(), input.GetItemPrice(), input.GetItemQuantity(), input.GetItemDescription(),
		input.GetItemLocation(), input.GetItemCategory(), time.Now().Unix(), input.GetLItemID())
	log.Println(query)
	if err := models.DB.Exec(query).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during DB query: %v", err.Error())
		return
	}

	//will replace existing image in S3, hence we dont remove image even if PATCH failed
	//invalidate redis
	if err := utils.InvalidateCache(utils.GetSingleListingByUserIDCacheKey, input.GetLItemID()); err != nil {
		log.Printf("Error during InvalidateCache: %v", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully updated listing details.")})
	log.Printf("Successful: UpdateSingleListing.")
}
