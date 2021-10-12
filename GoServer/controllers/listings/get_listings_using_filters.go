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

func ValidateGetListingsUsingFiltersRequest(c *gin.Context, input *models.GetListingsUsingFiltersRequest) {
	//allow nil and empty
	if err := c.ShouldBindJSON(&input); err != nil {

		if input.CategoryFilter.ItemCategory != nil && !utils.ValidateUint(input.CategoryFilter.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be uint type.")})
			return
		}

		if input.LocationFilter.Location != nil && !utils.ValidateString(input.LocationFilter.Location) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("location must be string type.")})
			return
		}

		if input.PriceFilter.MinPrice != nil && !utils.ValidateUint(input.PriceFilter.MinPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("min_price must be uint type.")})
			return
		}

		if input.PriceFilter.MaxPrice != nil && !utils.ValidateUint(input.PriceFilter.MaxPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("max_price must be uint type.")})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}
}

func ValidateGetListingsUsingFiltersInput(c *gin.Context, input *models.GetListingsUsingFiltersRequest) {
	//check if exists
	if input.CategoryFilter.ItemCategory != nil && !constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_ITEM_CATEGORY, input.GetItemCategory()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("unknown item_category.")})
		return
	}

	if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice != nil {
		if input.PriceFilter.GetMaxPrice() < input.PriceFilter.GetMinPrice() {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("min_price cannot > max_price.")})
			return
		}
	}
}

func GetListingsUsingFilters(c *gin.Context) {
	var (
		listings          []models.Listing
		input             models.GetListingsUsingFiltersRequest
		categoryCondition = ""
		locationCondition = ""
		priceCondition    = ""
	)

	if err := c.ShouldBindJSON(&input); err != nil {

		if input.CategoryFilter.ItemCategory != nil && !utils.ValidateUint(input.CategoryFilter.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be uint type.")})
			return
		}

		if input.LocationFilter.Location != nil && !utils.ValidateString(input.LocationFilter.Location) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("location must be string type.")})
			return
		}

		if input.PriceFilter.MinPrice != nil && !utils.ValidateUint(input.PriceFilter.MinPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("min_price must be uint type.")})
			return
		}

		if input.PriceFilter.MaxPrice != nil && !utils.ValidateUint(input.PriceFilter.MaxPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("max_price must be uint type.")})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	//check if exists
	if input.CategoryFilter.ItemCategory != nil && !constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_ITEM_CATEGORY, input.GetItemCategory()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("unknown item_category.")})
		return
	}

	if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice != nil {
		if input.PriceFilter.GetMaxPrice() < input.PriceFilter.GetMinPrice() {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("min_price cannot > max_price.")})
			return
		}
	}

	//build SQL queries

	//category filter
	if input.CategoryFilter.ItemCategory != nil {
		//else concat into query
		categoryCondition += " WHERE item_category = " + fmt.Sprint(input.CategoryFilter.GetItemCategory())
	}

	if input.LocationFilter.Location != nil {
		locationCondition += " item_location = '" + fmt.Sprint(input.LocationFilter.GetLocation()) + "'"
	}

	if categoryCondition != "" && locationCondition != "" {
		categoryCondition += " AND"
	} else if categoryCondition == "" {
		categoryCondition = " WHERE"
	}

	if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice == nil {
		priceCondition += " item_price >= " + fmt.Sprint(input.PriceFilter.GetMinPrice())
	} else if input.PriceFilter.MinPrice == nil && input.PriceFilter.MaxPrice != nil {
		priceCondition += " item_price <= " + fmt.Sprint(input.PriceFilter.GetMaxPrice())
	} else if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice != nil {
		priceCondition += " item_price >= " + fmt.Sprint(input.PriceFilter.GetMinPrice()) + " AND item_price <= " + fmt.Sprint(input.PriceFilter.GetMaxPrice())
	}

	if categoryCondition == " WHERE" && locationCondition == "" && priceCondition == "" {
		categoryCondition = ""
	}

	if priceCondition != "" && locationCondition != "" {
		locationCondition += " AND"
	}

	orderCondition := " ORDER BY listing_ctime DESC"
	query := "SELECT * FROM listing_tab" + categoryCondition + locationCondition + priceCondition + orderCondition
	log.Println(query)

	if err := models.DB.Raw(query).Scan(&listings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingsUsingFiltersResult(listings), "Data": listings})
}
