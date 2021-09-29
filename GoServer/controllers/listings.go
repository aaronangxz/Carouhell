package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetAllListings(c *gin.Context) {
	var (
		listings []models.GetAllListingsResponse
	)

	if err := models.DB.Raw("SELECT * FROM listings").Scan(&listings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetAllListingsResult(listings), "Data": listings})
}

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
		if input.ItemImg == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_img cannot be empty.")})
			return
		}
		if !utils.ValidateString(input.ItemImg) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_img must be string type.")})
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

	if input.GetItemImg() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_img cannot be empty.")})
		return
	}

	listings := models.Listing{
		ItemName:         input.ItemName,
		ItemPrice:        input.ItemPrice,
		ItemImg:          input.ItemImg,
		ItemCategory:     input.ItemCategory,
		ItemStatus:       input.ItemStatus,
		ItemCreationTime: time.Now().Unix(),
	}

	if err := models.DB.Create(&listings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}

func GetListingByItemID(c *gin.Context) {
	var (
		singleListing models.Listing
		input         models.GetSingleListingRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be uint type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetItemID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be > 0.")})
		return
	}

	if err := models.DB.Where("item_id = ?", input.ItemID).First(&singleListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": singleListing})
}

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
	if err := models.DB.Where("item_id = ?", input.ItemID).First(&originalListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		return
	}

	//If request fields are empty, we dont want to override empty fields into DB
	if input.ItemName == nil {
		input.ItemName = originalListing.ItemName
	}
	if input.ItemPrice == nil {
		input.ItemPrice = originalListing.ItemPrice
	}
	if input.ItemImg == nil {
		input.ItemImg = originalListing.ItemImg
	}

	//If all good, proceed to update
	if err := models.DB.Exec("UPDATE listings SET item_name = ?, item_price = ?, item_img = ?", input.ItemName, input.ItemPrice, input.ItemImg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}

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

func GetUserListings(c *gin.Context) {
	var (
		userListings   []models.Listing
		input          models.GetUserListingsRequest
		extraCondition = ""
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetUserID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be > 0.")})
		return
	}

	if input.Limit == nil {
		input.Limit = models.SetDefaultNotificationResponseLimit()
	}

	if utils.ValidateLimitMax(input.GetLimit(), models.MaxListingsResponseSize) {
		c.JSON(http.StatusBadRequest, gin.H{"RespMeta": models.NewParamErrorsResponse("limit exceeds max listing response size.")})
		return
	}

	switch input.GetLimit() {
	case 0:
		extraCondition = ""
	default:
		extraCondition = " LIMIT " + fmt.Sprint(input.GetLimit())
	}

	query := "SELECT * FROM listings WHERE user_id = ? ORDER BY listing_time DESC" + extraCondition

	if err := models.DB.Raw(query, input.GetUserID()).
		Scan(&userListings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	if len(userListings) == 0 {
		c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewNotFoundResponse(), "Data": userListings})
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse(), "Data": userListings})
}

func GetLatestListings(c *gin.Context) {
	var (
		input             models.GetLatestListingsRequest
		listings          []models.GetLatestListingsResponse
		categoryCondition = ""
		statusCondition   = ""
		limitCondition    = ""
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemCategory != nil && !utils.ValidateString(input.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be string type.")})
			return
		}
		if input.ItemStatus != nil && !utils.ValidateUint(input.ItemStatus) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_status must be uint type.")})
			return
		}
		if input.Limit != nil && !utils.ValidateUint(input.Limit) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("limit must be uint type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	//process item_category params
	// if nil or empty, we ignore
	if input.ItemCategory != nil && input.GetItemCategory() == "" {
		if len(input.GetItemCategory()) > int(models.MaxStringLength) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
			return
		}
		//else concat into query
		categoryCondition += " WHERE item_category = " + input.GetItemCategory()
	}

	//process item_status
	if input.ItemStatus != nil {
		switch input.GetItemStatus() {
		case constant.LISTING_STATUS_ALL:
			break
		case constant.LISTING_STATUS_NORMAL:
			statusCondition += " listing_status = " + fmt.Sprint(constant.LISTING_STATUS_NORMAL)
		case constant.LISTING_STATUS_SOLDOUT:
			statusCondition += " listing_status = " + fmt.Sprint(constant.LISTING_STATUS_SOLDOUT)
		case constant.LISTING_STATUS_DELETED:
			statusCondition += " listing_status = " + fmt.Sprint(constant.LISTING_STATUS_DELETED)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Unknown listing status.")})
			return
		}
	}

	if categoryCondition != "" && statusCondition != "" {
		categoryCondition += " AND"
	}

	//process limit
	//if nil, set to default value
	if input.Limit == nil {
		input.Limit = models.SetDefaultNotificationResponseLimit()
	}

	if utils.ValidateLimitMax(input.GetLimit(), models.MaxListingsResponseSize) {
		c.JSON(http.StatusBadRequest, gin.H{"RespMeta": models.NewParamErrorsResponse("limit exceeds max listing response size.")})
		return
	}

	switch input.GetLimit() {
	case 0:
		limitCondition += ""
	default:
		limitCondition += " LIMIT " + fmt.Sprint(input.GetLimit())
	}

	if statusCondition != "" && limitCondition != "" {
		statusCondition += " AND"
	}

	query := "SELECT * FROM listings" + categoryCondition + statusCondition + limitCondition + " ORDER BY item_creation_time DESC"

	if err := models.DB.Raw(query).Scan(&listings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetLatestListingsResult(listings), "Data": listings})
}
