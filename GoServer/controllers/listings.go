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
	models.DB.Create(&listings)

	c.JSON(http.StatusOK, gin.H{"data": listings})
}

func GetListingByItemID(c *gin.Context) {
	var SingleListing models.Listing

	if err := models.DB.Where("item_id = ?", c.Param("item_id")).First(&SingleListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": SingleListing})
}

func UpdateListing(c *gin.Context) {
	// Get model if exist
	var UpdatedListing models.Listing
	if err := models.DB.Where("item_id = ?", c.Param("item_id")).First(&UpdatedListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var Req models.UpdateListingRequest
	if err := c.ShouldBindJSON(&Req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&UpdatedListing).Updates(Req)

	c.JSON(http.StatusOK, gin.H{"data": UpdatedListing})
}

func DeleteListing(c *gin.Context) {
	// Get model if exist
	var DeleteListing models.Listing
	if err := models.DB.Where("item_id = ?", c.Param("item_id")).First(&DeleteListing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&DeleteListing)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
