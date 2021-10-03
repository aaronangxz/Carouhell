package listings

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

	if err := models.DB.Raw("SELECT * FROM listing_tab").Scan(&listings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetAllListingsResult(listings), "Data": listings})
}
