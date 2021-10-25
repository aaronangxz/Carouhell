package listings

import (
	"encoding/json"
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

func ValidateCreateListingRequest(c *gin.Context, input *models.CreateListingRequest) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemName == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot be empty.")})
			errormsg := "item_name cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.ItemName) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name must be string type.")})
			errormsg := fmt.Sprintf("item_name must be string type. input: %v", input.GetItemName())
			return errors.New(errormsg)
		}
		if input.ItemPrice == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price cannot be empty.")})
			errormsg := "item_price cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.ItemPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be uint type.")})
			errormsg := fmt.Sprintf("item_price must be uint type. input: %v", input.GetItemPrice())
			return errors.New(errormsg)
		}
		if input.ItemImage == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image cannot be empty.")})
			errormsg := "item_image cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.ItemImage) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image must be string type.")})
			errormsg := fmt.Sprintf("item_image must be string type. input: %v", input.GetItemImage())
			return errors.New(errormsg)
		}
		if input.ItemQuantity == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity cannot be empty.")})
			errormsg := "item_quantity cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.ItemQuantity) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity must be uint type.")})
			errormsg := fmt.Sprintf("item_quantity must be uint type. input: %v", input.GetItemQuantity())
			return errors.New(errormsg)
		}
		if input.ItemDescription == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_description cannot be empty.")})
			errormsg := "item_description cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.ItemDescription) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_description must be string type.")})
			errormsg := fmt.Sprintf("item_description must be string type. input: %v", input.GetItemDescription())
			return errors.New(errormsg)
		}
		if input.ItemShippingInfo == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_shipping_info cannot be empty.")})
			errormsg := "item_shipping_info cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.ItemShippingInfo) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_shipping_info must be uint type.")})
			errormsg := fmt.Sprintf("item_shipping_info must be uint type. input: %v", input.GetShippingInfo())
			return errors.New(errormsg)
		}
		if input.ItemPaymentInfo == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_payment_info cannot be empty.")})
			errormsg := "item_payment_info cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.ItemPaymentInfo) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_payment_info must be uint type.")})
			errormsg := fmt.Sprintf("item_payment_info must be uint type. input: %v", input.GetPaymentInfo())
			return errors.New(errormsg)
		}
		if input.ItemLocation == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location cannot be empty.")})
			errormsg := "item_location cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateString(input.ItemLocation) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location must be string type.")})
			errormsg := fmt.Sprintf("item_location must be string type. input: %v", input.GetItemLocation())
			return errors.New(errormsg)
		}
		if input.ItemCategory == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category cannot be empty.")})
			errormsg := "item_category cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be uint type.")})
			errormsg := fmt.Sprintf("item_category must be uint type. input: %v", input.GetItemCategory())
			return errors.New(errormsg)
		}
		if input.SellerID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("seller_id cannot be empty.")})
			errormsg := "seller_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.SellerID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("seller_id must be uint type.")})
			errormsg := fmt.Sprintf("seller_id must be uint type. input: %v", input.GetSellerID())
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func CreateListing(c *gin.Context) {
	// Validate input
	var (
		input models.CreateListingRequest
		hold  models.Account
	)

	if err := ValidateCreateListingRequest(c, &input); err != nil {
		log.Printf("Error during ValidateCreateListingRequest: %v", err.Error())
		return
	}

	if input.GetItemName() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot be empty.")})
		log.Print("item_name cannot be empty.")
		return
	}

	if !utils.ValidateMaxStringLength(input.GetItemName()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_name cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		log.Printf("item_name length cannot exceed %v. input :%v", models.MaxStringLength, len(input.GetItemName()))
		return
	}

	if input.GetItemPrice() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_price must be > 0.")})
		log.Print("item_price must be > 0.")
		return
	}

	if input.GetItemImage() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_image cannot be empty.")})
		log.Print("item_image cannot be empty.")
		return
	}

	if input.GetItemQuantity() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity must be > 0.")})
		log.Print("item_quantity must be > 0.")
		return
	}

	//allow blank
	if input.GetItemDescription() == "" {
		input.SetItemDescription("This item has no description.")
	}

	if input.GetItemLocation() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_location cannot be empty.")})
		log.Print("item_location cannot be empty.")
		return
	}

	//check if seller exists
	if err := models.DB.Raw("SELECT * FROM acc_tab WHERE user_id = ?", input.SellerID).Scan(&hold).Error; err != nil {
		if hold.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundMessageResponse("seller_id does not exist.")})
			log.Printf("seller not found:  %v", input.GetSellerID())
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during DB query: %v", err.Error())
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
		ItemImage:             nil,
		SellerID:              input.SellerID,
		ListingCtime:          utils.Int64(time.Now().Unix()),
		ListingMtime:          utils.Int64(time.Now().Unix()),
		ListingLikes:          utils.Uint32(0),
	}

	if err := models.DB.Table("listing_tab").Create(&listings).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during DB query: %v", err.Error())
		return
	}

	//upload image
	imageUrl, err := utils.UploadBase64Image(listings.GetItemID(), input.GetItemImage())
	if err != nil {
		//roll back listing create
		if errRollback := models.DB.Table("acc_tab").Delete(&listings).Error; errRollback != nil {
			log.Printf("Error during CreateListing - listing_tab roll back: %v", err.Error())
		} else {
			log.Print("rollback listing_tab successful")
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Failed to upload image.")})
		log.Printf("Error during image upload: %v", err)
		return
	}

	//write image URL to DB
	if err := models.DB.Exec("UPDATE listing_tab SET item_image = ?", imageUrl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during image write: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse(fmt.Sprintf("Successfully create listing. item_id: %v", listings.GetItemID()))})

	data, err := json.Marshal(listings)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v", err.Error())
	}
	log.Printf("Successful: CreateListing. Data: %s", data)
}
