package cart

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func ValidateAddItemToUserCartInput(c *gin.Context, input *models.AddItemToUserCartRequest) error {
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
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Item ID cannot be empty.")})
			errormsg := "item_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateInt64(input.ItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Item ID must be int64 type.")})
			errormsg := fmt.Sprintf("item_id must be int64 type. input: %v", input.GetUserID())
			return errors.New(errormsg)
		}
		if input.ItemQuantity == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Item Quantity cannot be empty.")})
			errormsg := "item_quantity cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.ItemQuantity) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Item Quantity must be uint type.")})
			errormsg := fmt.Sprintf("item_quantity must be uint type. input: %v", input.GetItemQuantity())
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func isItemExistsInCart(c *gin.Context, userid int64, itemid int64) (bool, error) {
	var count uint32

	//get current likes
	isItemExistsInCartQuery := models.DB.Table("user_cart_tab").Where("item_id = ? AND user_id = ?", itemid, userid).Count(&count)
	err := isItemExistsInCartQuery.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddItemToUserCart - isItemExistsInCart DB query: %v\n", err.Error())
		return false, err
	}

	return count > 0, err
}

func AddItemToUserCart(c *gin.Context) {
	var (
		input models.AddItemToUserCartRequest
	)

	if err := ValidateAddItemToUserCartInput(c, &input); err != nil {
		log.Printf("Error during ValidateAddItemToUserCartInput: %v", err.Error())
		return
	}

	//check if exists
	isExists, err := isItemExistsInCart(c, input.GetUserID(), input.GetItemID())
	if err != nil {
		return
	}

	//if item already in cart, increment quantity instead
	if isExists {
		incrementQuery := fmt.Sprintf("UPDATE user_cart_tab SET item_quantity = item_quantity + %v WHERE user_id = %v AND item_id = %v",
			input.GetItemQuantity(), input.GetUserID(), input.GetItemID())
		log.Println(incrementQuery)

		if err := models.DB.Exec(incrementQuery).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during AddItemToUserCart - IncrementQuantity DB query: %v", err.Error())
			return
		}
		log.Printf("Incremented item_quantity of existing item. item_id: %v", input.GetItemID())

	} else {
		//create new row
		cart := models.UserCart{
			UserID:       input.UserID,
			ItemID:       input.ItemID,
			ItemQuantity: input.ItemQuantity,
			Ctime:        utils.Int64(time.Now().Unix()),
		}

		if err := models.DB.Table("user_cart_tab").Create(&cart).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during AddItemToUserCart - InsertItemIntoCart DB query: %v", err.Error())
			return
		}
		log.Printf("Added item into cart. item_id: %v", input.GetItemID())
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successful: AddItemToUserCart.")})
	log.Printf("Successful: AddItemToUserCart. user_id: %v, item_id: %v", input.GetUserID(), input.GetItemID())
}
