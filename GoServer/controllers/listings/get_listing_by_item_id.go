package listings

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetListingByItemID(c *gin.Context) {
	var (
		singleListing models.GetSingleListingResponse
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

	query := utils.GetListingFixedQuery()
	log.Println(query)

	result := models.DB.Raw(query, input.GetItemID()).Scan(&singleListing)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetListingByItemID DB query: %v\n", err.Error())
		return
	}

	data, err := json.Marshal(singleListing)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v\n", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("GetListingByItemID success."), "Data": singleListing})
	log.Printf("Successful: GetListingByItemID. rows: %v\n", result.RowsAffected)
	log.Printf("Result: %s\n", data)
}
