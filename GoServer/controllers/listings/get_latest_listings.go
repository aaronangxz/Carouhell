package listings

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetLatestListings(c *gin.Context) {
	var (
		listings []models.GetLatestListingsResponse
	)

	//only return available items
	query := utils.GetListingQueryWithCustomCondition() + fmt.Sprintf(" AND l.item_status = %v GROUP BY l.l_item_id ORDER BY listing_ctime DESC", constant.ITEM_STATUS_NORMAL)
	log.Println(query)

	result := models.DB.Raw(query).Scan(&listings)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetLatestListings DB query: %v\n", err.Error())
		return
	}

	log.Printf("Successful: GetLatestListings. rows: %v\n", result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetLatestListingsResult(listings), "Data": listings})
}
