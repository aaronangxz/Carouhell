package controllers

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func GetAllListings(c *gin.Context) {
	var (
		allListings []models.Listing
	)

	models.DB.Find(&allListings)

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateBatchListingResult(allListings), "data": allListings})
}

func CreateListing(c *gin.Context) {
	// Validate input
	var input models.CreateListingRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listings := models.Listing{ItemName: input.ItemName, ItemPrice: input.ItemPrice, ItemImg: input.ItemImg}

	if err := models.DB.Create(&listings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewErrorResponse(err)})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("item_id = ?", input.ItemID).First(&singleListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": singleListing})
}

func UpdateSingleListing(c *gin.Context) {
	var (
		updatedListing models.Listing
		input          models.UpdateListingRequest
	)

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check if record exists
	if err := models.DB.Where("item_id = ?", input.ItemID).First(&updatedListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		return
	}

	//If all good, proceed to update
	if err := models.DB.Exec("UPDATE listings SET item_name = ?, item_price = ?, item_img = ?", input.ItemName, input.ItemPrice, input.ItemImg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewErrorResponse(err)})
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

	if err := models.DB.Where("item_id = ?", input.ItemID).First(&deleteListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		return
	}

	if err := models.DB.Delete(&deleteListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse()})
}
