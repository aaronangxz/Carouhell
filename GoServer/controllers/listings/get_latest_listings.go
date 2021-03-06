package listings

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/auth"
	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetLatestListings(c *gin.Context) {
	var (
		listings          []models.GetLatestListingsResponse
		listingsWithLikes []models.GetLatestListingsLoggedInResponse
		userId            uint64
	)

	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		log.Printf("Error during ExtractTokenMetadata: %v", err)
		//not logged in
		//only return available items
		query := utils.GetListingQueryWithCustomCondition() + fmt.Sprintf(" AND l.item_status = %v GROUP BY l.l_item_id ORDER BY listing_ctime DESC", constant.ITEM_STATUS_NORMAL)
		log.Println(query)

		result := models.DB.Raw(query).Scan(&listings)
		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetLatestListings DB query: %v\n", err.Error())
			return
		}

		log.Printf("Successful: GetLatestListings. rows: %v\n", result.RowsAffected)
		c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetLatestListingsResult(listings), "Data": listings})
		return
	}

	userId, err = auth.FetchAuth(tokenAuth)
	if err != nil {
		log.Printf("Error during FetchAuth: %v, user is not logged in.\n", err)
	}
	log.Printf("user_id requesting is %v\n", userId)

	//logged in
	//only return available items
	log.Println(utils.GetListingLoggedInQuery(userId))
	result := models.DB.Raw(utils.GetListingLoggedInQuery(userId)).Scan(&listingsWithLikes)
	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetLatestListings - LoggedIn DB query: %v\n", err.Error())
		return
	}

	log.Printf("Successful: GetLatestListings. rows: %v\n", result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetLatestListingsLoggedInResult(listingsWithLikes), "Data": listingsWithLikes})
}
