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
	"github.com/jinzhu/gorm"
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
		if input.GetItemQuantity() == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Item Quantity cannot be 0.")})
			errormsg := "item_quantity cannot be 0"
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

func checkUserCart(c *gin.Context, userid int64, itemid int64, quantity uint32, currentStock uint32) (bool, error) {
	var (
		userCart models.UserCart
	)

	//retrieve cart
	//check if item is already in cart
	userCartQuery := models.DB.Raw("SELECT * FROM user_cart_tab WHERE user_id = ? AND item_id = ?", userid, itemid).Scan(&userCart)
	err := userCartQuery.Error
	if err != nil {
		//not in cart
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("item_id: %v is not in user_id: %v cart, skipping quantity checks", itemid, userid)
			return false, nil
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddItemToUserCart - checkUserCart DB query: %v\n", err.Error())
		return false, err
	}

	//in cart
	//check if after adding item it will exceed item stock
	if userCart.GetItemQuantity()+quantity > currentStock {
		errormsg := fmt.Sprintf("item quantity will exceed total item stock if added into cart. item_id: %v", itemid)
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("The item quantity you want to add exceeds the total stock.")})
		return false, errors.New(errormsg)
	}
	return true, nil
}

func isItemHasEnoughStock(c *gin.Context, itemid int64, quantity uint32) (bool, uint32, error) {
	var (
		stockCount models.ListingQuantity
	)
	checkStockQuery := fmt.Sprintf("SELECT item_quantity FROM listing_tab WHERE l_item_id = %v", itemid)
	log.Println(checkStockQuery)
	result := models.DB.Raw(checkStockQuery).Scan(&stockCount)

	err := result.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Item does not exist.")})
			log.Printf("item_id: %v does not exist", err)
			return false, 0, err
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddItemToUserCart - isItemHasEnoughStock DB query: %v\n", err.Error())
		return false, 0, err
	}
	return quantity <= stockCount.GetItemQuantity(), stockCount.GetItemQuantity(), err
}

func AddItemToUserCart(c *gin.Context) {
	var (
		input models.AddItemToUserCartRequest
	)

	if err := ValidateAddItemToUserCartInput(c, &input); err != nil {
		log.Printf("Error during ValidateAddItemToUserCartInput: %v", err.Error())
		return
	}

	if input.GetUserID() < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Please log in to add items to cart.")})
		log.Printf("Anonymous user cannot add to cart: %v", input.GetUserID())
		return
	}

	//check if item has enough stock
	isHasStock, currentStock, err := isItemHasEnoughStock(c, input.GetItemID(), input.GetItemQuantity())
	if err != nil {
		return
	}

	if !isHasStock {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("There is not enough stock to add to cart.")})
		log.Printf("item: %v stock < quantity to add to cart. quantity to add: %v", input.GetItemID(), input.GetItemQuantity())
		return
	}
	log.Printf("item_id: %v has enough stock", input.GetItemID())

	//check if exists in cart
	//check if after adding item it will exceed item stock
	isCartOk, err := checkUserCart(c, input.GetUserID(), input.GetItemID(), input.GetItemQuantity(), currentStock)
	if err != nil {
		return
	}

	//if item already in cart, increment quantity instead
	if isCartOk {
		log.Printf("item_id: %v already exists in cart of user_id: %v", input.GetItemID(), input.GetUserID())

		//check if after adding item, quantity in cart > stock

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
		log.Printf("item_id: %v does not exist in cart of user_id: %v", input.GetItemID(), input.GetUserID())
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
