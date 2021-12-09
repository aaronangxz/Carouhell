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

func GetLatestListingsv2(c *gin.Context) {
	var (
		input    models.GetLatestListingsv2Request
		listings []models.GetLatestListingsv2Response
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		log.Printf("JSON error: %v\n", err.Error())
		return
	}

	if input.GetUserID() == -1 {
		//not logged in
		query := utils.GetListingQueryWithCustomCondition() + fmt.Sprintf(" AND l.item_status = %v GROUP BY l.l_item_id ORDER BY listing_ctime DESC", constant.ITEM_STATUS_NORMAL)
		log.Println(query)
		result := models.DB.Raw(query).Scan(&listings)
		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetLatestListingsV2 - not logged in DB query: %v\n", err.Error())
			return
		}
		log.Printf("Successful: GetLatestListingsV2 - not logged in. rows: %v\n", result.RowsAffected)
	} else {
		//logged in
		//only return available items
		log.Println(utils.GetListingLoggedInQuery(uint64(input.GetUserID())))
		result := models.DB.Raw(utils.GetListingLoggedInQuery(uint64(input.GetUserID()))).Scan(&listings)
		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetLatestListingsV2 - loggedIn DB query: %v\n", err.Error())
			return
		}
		log.Printf("Successful: GetLatestListingsV2 - logged in. rows: %v\n", result.RowsAffected)
	}
	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetLatestListingsv2Result(listings), "Data": listings})
}
