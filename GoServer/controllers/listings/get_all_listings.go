package listings

import (
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetAllListings(c *gin.Context) {
	var (
		listings []models.GetAllListingsResponse
	)

	query := utils.GetListingFixedQuery()
	log.Println(query)

	result := models.DB.Raw(query).Scan(&listings)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetAllListings DB query: %v\n", err.Error())
		return
	}

	log.Printf("Successful: GetAllListings. rows: %v\n", result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetAllListingsResult(listings), "Data": listings})
}
