package listings

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func ValidateGetListingsUsingFiltersRequest(c *gin.Context, input *models.GetListingsUsingFiltersRequest) error {
	//allow nil and empty
	if err := c.ShouldBindJSON(&input); err != nil {

		//not checking search term since we accept any string
		if input.CategoryFilter.ItemCategory != nil && !utils.ValidateUint(input.CategoryFilter.ItemCategory) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_category must be uint type.")})
			errormsg := fmt.Sprintf("item_category must be uint. input: %v", input.CategoryFilter.GetItemCategory())
			return errors.New(errormsg)
		}

		if input.LocationFilter.Location != nil && !utils.ValidateUint(input.LocationFilter.Location) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("location must be uint type.")})
			errormsg := fmt.Sprintf("location must be uint. input: %v", input.LocationFilter.GetLocation())
			return errors.New(errormsg)
		}

		if input.PriceFilter.MinPrice != nil && !utils.ValidateUint(input.PriceFilter.MinPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("min_price must be uint type.")})
			errormsg := fmt.Sprintf("min_price must be uint. input: %v", input.PriceFilter.GetMinPrice())
			return errors.New(errormsg)
		}

		if input.PriceFilter.MaxPrice != nil && !utils.ValidateUint(input.PriceFilter.MaxPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("max_price must be uint type.")})
			errormsg := fmt.Sprintf("max_price must be uint. input: %v", input.PriceFilter.GetMaxPrice())
			return errors.New(errormsg)
		}

		if input.SortFlag != nil && !utils.ValidateUint(input.SortFlag) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("sort_flag must be uint type.")})
			errormsg := fmt.Sprintf("sort_flag must be uint. input: %v", input.PriceFilter.GetMinPrice())
			return errors.New(errormsg)
		}

		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprintf("JSON error: %v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func ValidateGetListingsUsingFiltersInput(c *gin.Context, input *models.GetListingsUsingFiltersRequest) error {

	if input.SearchKeyword != nil && !utils.ValidateMaxStringLength(input.GetSearchKeyword()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("search_keyword cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		errormsg := fmt.Sprintf("search_keyword exceeded max length. input: %v chars", len(input.GetSearchKeyword()))
		return errors.New(errormsg)
	}

	//check if exists
	if input.CategoryFilter.ItemCategory != nil && !constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_ITEM_CATEGORY, input.GetItemCategory()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("unknown item_category.")})
		errormsg := fmt.Sprintf("unknown item_category. input: %v", input.CategoryFilter.GetItemCategory())
		return errors.New(errormsg)
	}

	if input.SortFlag != nil && !constant.CheckSearchAndFiltersConstant(constant.SEARCH_AND_FILTERS_CONSTANT_SORTFLAG, input.GetSortFlag()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("unknown sort_flag.")})
		errormsg := fmt.Sprintf("unknown sort_flag. input: %v", input.GetSortFlag())
		return errors.New(errormsg)
	}
	//check location
	if input.LocationFilter.Location != nil && !(constant.CheckListingConstant(constant.LISTING_CONSTANT_TYPE_LOCATION_DIRECTION, input.GetLocation())) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("unknown item_location")})
		errormsg := fmt.Sprintf("unknown item_location. input: %v", input.GetLocation())
		return errors.New(errormsg)
	}
	if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice != nil {
		if input.PriceFilter.GetMaxPrice() < input.PriceFilter.GetMinPrice() {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("min_price cannot > max_price.")})
			errormsg := fmt.Sprintf("min_price cannot > max_price. input: min_price: %v, max_price: %v", input.PriceFilter.GetMinPrice(), input.PriceFilter.GetMaxPrice())
			return errors.New(errormsg)
		}
	}
	return nil
}

func GetListingsUsingFilters(c *gin.Context) {
	var (
		listings          []models.GetListingsUsingFiltersResponse
		input             models.GetListingsUsingFiltersRequest
		query             = ""
		categoryCondition = ""
		locationCondition = ""
		priceCondition    = ""
		groupCondition    = " GROUP BY l.l_item_id"
		orderCondition    = " ORDER BY"
	)

	if err := ValidateGetListingsUsingFiltersRequest(c, &input); err != nil {
		log.Printf("Error during ValidateGetListingsUsingFiltersRequest: %v", err.Error())
		return
	}

	if err := ValidateGetListingsUsingFiltersInput(c, &input); err != nil {
		log.Printf("Error during ValidateGetListingsUsingFiltersInput: %v", err.Error())
		return
	}

	//category filter
	if input.CategoryFilter.ItemCategory != nil {
		//else concat into query
		// categoryCondition += " item_category = " + fmt.Sprint(input.CategoryFilter.GetItemCategory())
		categoryCondition += fmt.Sprintf(" AND item_category = %v", input.CategoryFilter.GetItemCategory())

	}

	//build location SQL query
	if input.Location != nil {
		locationCondition = " AND FIND_IN_SET(item_location,'"
		locationList := constant.GetLocationFromDirection(input.GetLocation())

		for i, element := range locationList {
			if i == len(locationList)-1 {
				locationCondition += fmt.Sprintf("%v')", element)
				break
			}
			locationCondition += fmt.Sprintf("%v,", element)
		}
	}

	if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice == nil {
		// priceCondition += " item_price >= " + fmt.Sprint(input.PriceFilter.GetMinPrice())
		priceCondition += fmt.Sprintf(" AND item_price >= %v", input.PriceFilter.GetMinPrice())
	} else if input.PriceFilter.MinPrice == nil && input.PriceFilter.MaxPrice != nil {
		// priceCondition += " item_price <= " + fmt.Sprint(input.PriceFilter.GetMaxPrice())
		priceCondition += fmt.Sprintf("  AND item_price <= %v", input.PriceFilter.GetMaxPrice())
	} else if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice != nil {
		// priceCondition += " item_price >= " + fmt.Sprint(input.PriceFilter.GetMinPrice()) + " AND item_price <= " + fmt.Sprint(input.PriceFilter.GetMaxPrice())
		priceCondition += fmt.Sprintf(" AND item_price >= %v AND item_price <= %v", input.PriceFilter.GetMinPrice(), input.PriceFilter.GetMaxPrice())
	}

	if input.SortFlag == nil {
		input.SortFlag = utils.Uint32(constant.SEARCH_RESULT_SORTFLAG_DEFAULT)
	}

	switch input.GetSortFlag() {
	case constant.SEARCH_RESULT_SORTFLAG_DEFAULT, constant.SEARCH_RESULT_SORTFLAG_RECENT:
		orderCondition += " listing_ctime DESC"
	case constant.SEARCH_RESULT_SORTFLAG_PRICE_HIGH:
		orderCondition += " item_price DESC,listing_ctime DESC "
	case constant.SEARCH_RESULT_SORTFLAG_PRICE_LOW:
		orderCondition += " item_price ASC, listing_ctime DESC"
	case constant.SEARCH_RESULT_SORTFLAG_POPULAR:
		orderCondition += " listing_likes DESC, listing_ctime DESC"
	}

	if input.SearchKeyword != nil {
		query = utils.GetFullTextSearchQuery(input.GetSearchKeyword()) + categoryCondition + locationCondition + priceCondition + groupCondition +
			" ORDER BY relevance DESC" + strings.ReplaceAll(orderCondition, " ORDER BY", " ,")
	} else {
		query = utils.GetListingQueryWithCustomCondition() + fmt.Sprintf(" AND l.item_status = %v", constant.ITEM_STATUS_NORMAL) +
			categoryCondition + locationCondition + priceCondition + groupCondition + orderCondition
	}

	log.Printf("Executing DB query: %v\n", query)

	result := models.DB.Raw(query).Scan(&listings)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetListingsUsingFilters DB query: %v\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingsUsingFiltersResult(listings), "Data": listings})
	log.Printf("Successful: GetListingsUsingFilters. rows: %v\n", result.RowsAffected)
}
