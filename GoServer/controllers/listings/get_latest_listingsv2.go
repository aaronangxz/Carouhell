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
		input          models.GetLatestListingsv2Request
		listings       []models.GetLatestListingsv2Response
		paginationResp models.Pagination
		resp           models.GetLatestListingsv2WithPaginationResponse
		totalCount     uint32
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		log.Printf("JSON error: %v\n", err.Error())
		return
	}

	//get total number of listings
	totalCountQuery := models.DB.Table("listing_tab").
		Where(" item_status = ?", constant.ITEM_STATUS_NORMAL).Count(&totalCount)
	if err := totalCountQuery.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetLatestListingsv2 - totalCount DB query: %v\n", err.Error())
		return
	}

	//default results per page
	resultsPerPage := uint32(10)

	//total pages
	totalPages := uint32(totalCount / resultsPerPage)

	paginationResp.TotalPage = totalPages
	paginationResp.CurrentPage = input.GetPagination()

	if input.GetPagination() >= totalPages {
		paginationResp.NextPage = 0
	} else {
		paginationResp.NextPage = input.GetPagination() + 1
	}

	//current offset
	offset := resultsPerPage * (input.GetPagination() - 1)

	if input.GetUserID() == -1 {
		//not logged in
		query := utils.GetListingQueryWithCustomCondition() + fmt.Sprintf(" AND l.item_status = %v GROUP BY l.l_item_id ORDER BY listing_ctime DESC LIMIT %v OFFSET %v",
			constant.ITEM_STATUS_NORMAL,
			resultsPerPage,
			offset)
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
		result := models.DB.Raw(utils.GetListingLoggedInQuery(uint64(input.GetUserID()))+" LIMIT ? OFFSET ?", resultsPerPage, offset).Scan(&listings)
		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetLatestListingsV2 - loggedIn DB query: %v\n", err.Error())
			return
		}
		log.Printf("Successful: GetLatestListingsV2 - logged in. rows: %v\n", result.RowsAffected)
	}

	resp.Listings = listings
	resp.Paginations = paginationResp

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetLatestListingsv2Result(listings), "Data": resp})
}
