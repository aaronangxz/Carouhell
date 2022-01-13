package cart

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ValidateGetUserCartInput(c *gin.Context, input *models.GetUserCartRequest) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("User ID cannot be empty.")})
			errormsg := "user_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateInt64(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("User ID must be int64 type.")})
			errormsg := fmt.Sprintf("user_id must be int64 type. input: %v", input.GetUserID())
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func GetUserCart(c *gin.Context) {
	var (
		input          models.GetUserCartRequest
		currentCart    []models.UserCart
		itemInfo       models.UserCartItem
		itemInfoSorted []models.UserCartItem
		validItems     []models.UserCartItem
		invalidItems   []models.UserCartItem
		resp           models.GetUserCartResponse
	)

	if err := ValidateGetUserCartInput(c, &input); err != nil {
		log.Printf("Error during ValidateGetUserCartInput: %v", err.Error())
		return
	}

	//get cart items
	currentCartQuery := fmt.Sprintf("SELECT * FROM user_cart_tab WHERE user_id = %v ORDER BY ctime DESC", input.GetUserID())
	log.Println(currentCartQuery)
	currentCartResult := models.DB.Raw(currentCartQuery).Scan(currentCart)

	if err := currentCartResult.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No items in cart. user_id: %v", input.GetUserID())
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Cart is empty.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserCart - GetCurrentCart DB query: %v", err.Error())
		return
	}

	//get item info for each item in cart
	for i, cartItemsId := range currentCart {
		//get item info
		getItemInfoQuery := fmt.Sprintf("SELECT * FROM listing_tab WHERE l_item_id = %v", cartItemsId)
		log.Println(getItemInfoQuery)
		getItemInfoResult := models.DB.Raw(getItemInfoQuery).Scan(itemInfo)
		if err := getItemInfoResult.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetUserCart - GetCurrentCart DB query: %v", err.Error())
			return
		}
		//fill in cart ctime into object
		itemInfo.CartCtime = currentCart[i].GetCtime()
		itemInfoSorted = append(itemInfoSorted, itemInfo)
	}

	//check each items
	for i, items := range itemInfoSorted {

		//status != deleted, soldout
		if items.ItemStatus == constant.ITEM_STATUS_DELETED {
			items.ItemInfo.InvalidMessage = "Item is deleted."
			items.ItemInfo.InvalidErrorCode = uint32(constant.INVALID_CART_ITEM_DELETED)
			invalidItems = append(invalidItems, items)
			continue
		} else if items.ItemStatus == constant.ITEM_STATUS_SOLDOUT {
			items.ItemInfo.InvalidMessage = "Item is out of stock."
			items.ItemInfo.InvalidErrorCode = uint32(constant.INVALID_CART_ITEM_SOLDOUT)
			invalidItems = append(invalidItems, items)
			continue
		}

		//item quantity >= cart quantity
		if items.ItemQuantity < currentCart[i].GetItemQuantity() {
			items.ItemInfo.InvalidMessage = "Item has not enough stock."
			items.ItemInfo.InvalidErrorCode = uint32(constant.INVALID_CART_ITEM_NOT_ENOUGH_STOCK)
			invalidItems = append(invalidItems, items)
			continue
		}

		//if valid, add into valid slice
		validItems = append(validItems, items)
	}

	resp.ValidCount = len(validItems)
	resp.InvalidCount = len(invalidItems)

	sort.Slice(validItems, func(i, j int) bool {
		return validItems[i].CartCtime < validItems[j].CartCtime
	})

	sort.Slice(invalidItems, func(i, j int) bool {
		return invalidItems[i].CartCtime < invalidItems[j].CartCtime
	})

	resp.ValidItems = validItems
	resp.InvalidItems = invalidItems

}
