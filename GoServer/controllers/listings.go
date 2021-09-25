package controllers

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	listings := models.Listing{ItemName: input.ItemName, ItemPrice: input.ItemPrice, ItemImg: input.ItemImg}

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
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
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
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
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
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
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
		userListings models.Listing
		input        models.GetUserListingsRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if utils.ValidateLimitMax(input.Limit, models.MaxListingsResponseSize) {
		c.JSON(http.StatusBadRequest, gin.H{"RespMeta": models.NewParamErrorsResponse("limit exceeds max listing response size")})
		return
	}

	if err := models.DB.Raw("SELECT * FROM listings WHERE user_id = ? ORDER BY listing_time DESC LIMIT ?", input.UserID, input.Limit).
		Scan(&userListings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse(), "Data": userListings})
}
