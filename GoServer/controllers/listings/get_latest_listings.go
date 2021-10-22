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
		input             models.GetLatestListingsRequest
		listings          []models.GetLatestListingsResponse
		categoryCondition = ""
		statusCondition   = ""
		limitCondition    = ""
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemCategory != nil && !utils.ValidateUint(input.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be uint type.")})
			return
		}
		if input.ItemStatus != nil && !utils.ValidateUint(input.ItemStatus) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_status must be uint type.")})
			return
		}
		if input.Limit != nil && !utils.ValidateUint(input.Limit) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("limit must be uint type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	//process item_category params
	// if nil, we ignore
	if input.ItemCategory != nil {
		//else concat into query
		categoryCondition += " AND item_category = " + fmt.Sprint(input.GetItemCategory())
	}

	//process item_status
	if input.ItemStatus != nil {
		switch input.GetItemStatus() {
		case constant.ITEM_STATUS_ALL:
			break
		case constant.ITEM_STATUS_NORMAL:
			statusCondition += " item_status = " + fmt.Sprint(constant.ITEM_STATUS_NORMAL)
		case constant.ITEM_STATUS_SOLDOUT:
			statusCondition += " item_status = " + fmt.Sprint(constant.ITEM_STATUS_SOLDOUT)
		case constant.ITEM_STATUS_DELETED:
			statusCondition += " item_status = " + fmt.Sprint(constant.ITEM_STATUS_DELETED)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Unknown item_status.")})
			return
		}
	}

	if categoryCondition != "" && statusCondition != "" {
		categoryCondition += " AND"
	} else if categoryCondition == "" {
		categoryCondition = ""
	}

	//process limit
	//if nil, set to default value
	if input.Limit == nil {
		input.Limit = models.SetDefaultNotificationResponseLimit()
	}

	if utils.ValidateLimitMax(input.GetLimit(), models.MaxListingsResponseSize) {
		c.JSON(http.StatusBadRequest, gin.H{"RespMeta": models.NewParamErrorsResponse("limit exceeds max listing response size.")})
		return
	}

	switch input.GetLimit() {
	case 0:
		limitCondition += ""
	default:
		limitCondition += " LIMIT " + fmt.Sprint(input.GetLimit())
	}

	if input.ItemCategory == nil && input.ItemStatus == nil {
		categoryCondition = ""
		statusCondition = ""
	}

	orderCondition := " ORDER BY listing_ctime DESC"
	groupCondition := " GROUP BY item_id"
	joinCondition := " WHERE l.seller_id = a.user_id"

	query := "SELECT l.item_id, l.item_name, l.item_price, l.item_quantity," +
		" l.item_purchased_quantity, l.item_description, l.item_shipping_info," +
		" l.item_payment_info,l.item_location, l.item_status, l.item_category," +
		" l.item_image, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, l.listing_likes" +
		" FROM listing_tab l, acc_tab a " + joinCondition +
		categoryCondition + statusCondition + groupCondition + orderCondition + limitCondition

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
