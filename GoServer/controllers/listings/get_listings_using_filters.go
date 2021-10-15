package listings

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

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

		if input.LocationFilter.Location != nil && !utils.ValidateString(input.LocationFilter.Location) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("location must be string type.")})
			errormsg := fmt.Sprintf("location must be string. input: %v", input.LocationFilter.GetLocation())
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
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
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
		listings          []models.Listing
		input             models.GetListingsUsingFiltersRequest
		nameCondition     = ""
		categoryCondition = ""
		locationCondition = ""
		priceCondition    = ""
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

	//build SQL queries
	if input.SearchKeyword != nil {
		key := "'%" + input.GetSearchKeyword() + "%'"
		nameCondition += fmt.Sprintf(" WHERE item_name LIKE %v", key)
	}

	//category filter
	if input.CategoryFilter.ItemCategory != nil {
		//else concat into query
		// categoryCondition += " item_category = " + fmt.Sprint(input.CategoryFilter.GetItemCategory())
		categoryCondition += fmt.Sprintf(" item_category = %v", input.CategoryFilter.GetItemCategory())

	}

	if nameCondition != "" && categoryCondition != "" {
		nameCondition += " AND"
	} else if nameCondition == "" {
		nameCondition = " WHERE"
	}

	if input.LocationFilter.Location != nil {
		// locationCondition += " item_location = '" + fmt.Sprint(input.LocationFilter.GetLocation()) + "'"
		locationCondition += fmt.Sprintf(" item_location = '%v'", input.LocationFilter.GetLocation())
	}

	if categoryCondition != "" && locationCondition != "" {
		categoryCondition += " AND"
	}

	if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice == nil {
		// priceCondition += " item_price >= " + fmt.Sprint(input.PriceFilter.GetMinPrice())
		priceCondition += fmt.Sprintf(" item_price >= %v", input.PriceFilter.GetMinPrice())
	} else if input.PriceFilter.MinPrice == nil && input.PriceFilter.MaxPrice != nil {
		// priceCondition += " item_price <= " + fmt.Sprint(input.PriceFilter.GetMaxPrice())
		priceCondition += fmt.Sprintf(" item_price <= %v", input.PriceFilter.GetMaxPrice())
	} else if input.PriceFilter.MinPrice != nil && input.PriceFilter.MaxPrice != nil {
		// priceCondition += " item_price >= " + fmt.Sprint(input.PriceFilter.GetMinPrice()) + " AND item_price <= " + fmt.Sprint(input.PriceFilter.GetMaxPrice())
		priceCondition += fmt.Sprintf(" item_price >= %v AND item_price <= %v", input.PriceFilter.GetMinPrice(), input.PriceFilter.GetMaxPrice())
	}

	if nameCondition == " WHERE" && categoryCondition == "" && locationCondition == "" && priceCondition == "" {
		nameCondition = ""
	}

	if priceCondition != "" && locationCondition != "" {
		locationCondition += " AND"
	}

	if input.SortFlag == nil {
		input.SortFlag = utils.Uint32(constant.SEARCH_RESULT_SORTFLAG_DEFAULT)
	}

	switch input.GetSortFlag() {
	case constant.SEARCH_RESULT_SORTFLAG_DEFAULT, constant.SEARCH_RESULT_SORTFLAG_RECENT:
		orderCondition += " listing_ctime DESC"
	case constant.SEARCH_RESULT_SORTFLAG_PRICE_HIGH:
		orderCondition += " item_price DESC"
	case constant.SEARCH_RESULT_SORTFLAG_PRICE_LOW:
		orderCondition += " item_price ASC"
	case constant.SEARCH_RESULT_SORTFLAG_RATINGS: //pending case 4: sort by ratings
		orderCondition += " listing_ctime DESC"
	}

	query := "SELECT * FROM listing_tab" + nameCondition + categoryCondition + locationCondition + priceCondition + orderCondition
	log.Printf("Executing DB query: %v", query)

	if err := models.DB.Raw(query).Scan(&listings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during DB query: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingsUsingFiltersResult(listings), "Data": listings})
	data, err := json.Marshal(listings)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v", err.Error())
	}
	log.Printf("Successful: GetListingsUsingFilters. Returned: %s", data)
}
