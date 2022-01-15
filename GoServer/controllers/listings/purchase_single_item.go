package listings

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func isReviewed(c *gin.Context, input models.PurchaseSingleItemRequest) bool {
	var (
		count uint32
	)

	result := models.DB.Table("user_review_tab").
		Where("rv_user_id = ? AND rv_seller_id IN (SELECT l_seller_id FROM listing_tab WHERE l_item_id = ?)",
			input.GetUserID(), input.GetItemID()).Count(&count)
	log.Printf("Checking if user_id:%v reviewed seller of item_id:%v before\n", input.GetUserID(), input.GetUserID())

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during PurchaseSingleItem - isReviewed DB query: %v\n", err.Error())
		return true
	}

	if count > 0 {
		return true
	}
	return false
}

func ValidatePurchaseSingleItemRequest(c *gin.Context, input *models.PurchaseSingleItemRequest) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			errormsg := "user_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be uint type.")})
			errormsg := "user_id must be uint type"
			return errors.New(errormsg)
		}
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			errormsg := "item_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.ItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be uint type.")})
			errormsg := "item_id must be uint type"
			return errors.New(errormsg)
		}
		if input.PurchaseQuantity == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("purchase_quantity cannot be empty.")})
			errormsg := "purchase_quantity cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.PurchaseQuantity) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("purchase_quantity must be uint type.")})
			errormsg := "purchase_quantity must be uint type"
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func ValidatePurchaseSingleItemInput(c *gin.Context, input *models.PurchaseSingleItemRequest) error {
	if input.GetPurchaseQuantity() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_quantity must be > 0.")})
		errormsg := fmt.Sprintf("item_quantity must be > 0. input: %v", input.GetPurchaseQuantity())
		return errors.New(errormsg)
	}
	return nil
}

func PurchaseSingleItem(c *gin.Context) {
	var (
		input       models.PurchaseSingleItemRequest
		listingHold models.Listing
		walletHold  models.Wallet
		resp        models.PurchaseSingleItemResponse
	)

	if err := ValidatePurchaseSingleItemRequest(c, &input); err != nil {
		log.Printf("Error during ValidatePurchaseSingleItemRequest: %v", err.Error())
		return
	}

	if err := ValidatePurchaseSingleItemInput(c, &input); err != nil {
		log.Printf("Error during ValidatePurchaseSingleItemInput: %v", err.Error())
		return
	}

	//check item status (not sold, quantity not 0)
	//check price of listing
	//retrieve listing details
	listingQuery := fmt.Sprintf("SELECT * FROM listing_tab WHERE l_item_id = %v", input.GetItemID())
	log.Println(listingQuery)
	listing := models.DB.Raw(listingQuery).Scan(&listingHold)
	if err := listing.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during PurchaseSingleItem - get listing DB query: %v", err.Error())
		return
	}

	totalPrice := input.GetPurchaseQuantity() * listingHold.GetItemPrice()

	//item is not available
	if listingHold.GetItemStatus() != constant.ITEM_STATUS_NORMAL {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Item is not available for purchase.")})
		log.Printf("not available for purchase. item_status:%v", listingHold.GetItemStatus())
		return
	}

	//not enough stock to purchase
	if input.GetPurchaseQuantity() > listingHold.GetItemQuantity() {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("There is not enough quantity to purchase.")})
		log.Printf("item_quantity < purchase_quantity. purchase_quantity: %v, item_quantity: %v", input.GetPurchaseQuantity(), listingHold.GetItemQuantity())
		return
	}

	//retrieve wallet details
	walletQuery := fmt.Sprintf("SELECT * FROM wallet_tab WHERE w_user_id = %v", input.GetUserID())
	log.Println(walletQuery)
	wallet := models.DB.Raw(walletQuery).Scan(&walletHold)
	if err := wallet.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during PurchaseSingleItem - get wallet DB query: %v", err.Error())
		return
	}

	//check if wallet balance > item price * quantity
	if walletHold.GetWalletBalance() < totalPrice {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Wallet balance is insufficient.")})
		log.Printf("wallet_balance < total price. wallet_balance: %v, total price: %v", walletHold.GetWalletBalance(), totalPrice)
		return
	}

	/*
		//insert into listing_transaction
		listingTransaction := models.ListingTransaction{
			LtItemID:            input.ItemID,
			LtUserID:            input.UserID,
			TransactionCtime:    utils.Int64(time.Now().Unix()),
			TransactionQuantity: input.PurchaseQuantity,
			TransactionAmount:   utils.Uint32(totalPrice),
			TransactionStatus:   utils.Uint32(constant.LISTING_TRANSACTION_STATUS_SUCCESS),
		}

		if err := models.DB.Table("listing_transactions_tab").Create(&listingTransaction).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during PurchaseSingleItem - update listing transaction tab DB query: %v", err.Error())
			return
		}

		//insert into wallet_transaction
		walletTransaction := models.WalletTransaction{
			WtWalletID:        input.UserID,
			TransactionCtime:  utils.Int64(time.Now().Unix()),
			TransactionAmount: utils.Uint32(totalPrice),
			TransactionType:   utils.Uint32(constant.TRANSACTION_TYPE_PURCHASE),
			TransactionRef:    listingTransaction.LtTransactionID,
		}
		if err := models.DB.Table("wallet_transactions_tab").Create(&walletTransaction).Error; err != nil {
			//if fail, rollback listing_transaction
			if err := models.DB.Table("listing_transactions_tab").Delete(&listingTransaction).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
				log.Printf("Error during PurchaseSingleItem - rollback listing transaction tab DB query: %v", err.Error())
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during PurchaseSingleItem - update wallet transaction tab DB query: %v", err.Error())
			return
		}

		//update listing quantity
		updateListingQuery := fmt.Sprintf("UPDATE listing_tab SET item_quantity = item_quantity - 1, item_status = CASE WHEN item_quantity = 0 THEN 2 ELSE item_status END WHERE l_item_id = %v", input.GetItemID())
		log.Println(updateListingQuery)
		updateListing := models.DB.Exec(updateListingQuery)
		if err := updateListing.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during PurchaseSingleItem - update listing quantity DB query: %v", err.Error())
			return
		}

		//update wallet balance
		updateWalletQuery := fmt.Sprintf("UPDATE wallet_tab SET wallet_balance = wallet_balance - %v ,last_used = %v WHERE wallet_id = %v", totalPrice, time.Now().Unix(), input.GetUserID())
		log.Println(updateWalletQuery)
		updateWallet := models.DB.Exec(updateWalletQuery)
		if err := updateWallet.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during PurchaseSingleItem - update wallet balance DB query: %v", err.Error())
			return
		}
	*/

	if err := utils.StartItemPurchaseTx(input, totalPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Error during item purchase.")})
		log.Printf("Error during StartItemPurchaseTx: %v", err.Error())
		return
	}

	newBalance := walletHold.GetWalletBalance() - totalPrice
	resp.WalletBalance = newBalance

	//check if reviewed before
	resp.IsReviewed = isReviewed(c, input)

	//invalidate redis
	if err := utils.InvalidateCache(utils.GetSingleListingByUserIDCacheKey, input.GetItemID()); err != nil {
		log.Printf("Error during InvalidateCache: %v", err.Error())
	}
	if err := utils.InvalidateCache(utils.GetUserWalletDetailsCacheKey, input.GetUserID()); err != nil {
		log.Printf("Error during InvalidateCache: %v", err.Error())
	}
	if err := utils.InvalidateCache(utils.CartItemInfoCacheKey, input.GetItemID()); err != nil {
		log.Printf("Error during InvalidateCache: %v", err.Error())
	}

	//if became out of stock, need to invalidate
	if listingHold.GetItemQuantity()-input.GetPurchaseQuantity() == 0 {
		if err := utils.InvalidateCache(utils.GetUserDetailsCacheKey, input.GetUserID()); err != nil {
			log.Printf("Error during InvalidateCache: %v", err.Error())
		}
		if err := utils.InvalidateCache(utils.GetUserLikedListingsCacheKey, input.GetUserID()); err != nil {
			log.Printf("Error during InvalidateCache: %v", err.Error())
		}
	}

	//check if item stock state changes after purchase
	//FE changes display when stock is lower and equal to 25%
	//dont invalidate if quantity before purchase is already within 'selling fast' & 'low stock' stage
	if ((listingHold.GetItemQuantity())/listingHold.GetItemStock())*100 > 25 &&
		((listingHold.GetItemQuantity()-input.GetPurchaseQuantity())/listingHold.GetItemStock())*100 <= 25 {
		log.Printf("item state is selling fast: %v", (listingHold.GetItemQuantity()-input.GetPurchaseQuantity())/listingHold.GetItemStock()*100)
		if err := utils.InvalidateCache(utils.GetUserDetailsCacheKey, input.GetUserID()); err != nil {
			log.Printf("Error during InvalidateCache: %v", err.Error())
		}
		if err := utils.InvalidateCache(utils.GetUserLikedListingsCacheKey, input.GetUserID()); err != nil {
			log.Printf("Error during InvalidateCache: %v", err.Error())
		}
	} else if ((listingHold.GetItemQuantity())/listingHold.GetItemStock())*100 > 10 &&
		((listingHold.GetItemQuantity()-input.GetPurchaseQuantity())/listingHold.GetItemStock())*100 <= 10 {
		log.Printf("item state is low in stock: %v", (listingHold.GetItemQuantity()-input.GetPurchaseQuantity())/listingHold.GetItemStock()*100)
		if err := utils.InvalidateCache(utils.GetUserDetailsCacheKey, input.GetUserID()); err != nil {
			log.Printf("Error during InvalidateCache: %v", err.Error())
		}
		if err := utils.InvalidateCache(utils.GetUserLikedListingsCacheKey, input.GetUserID()); err != nil {
			log.Printf("Error during InvalidateCache: %v", err.Error())
		}
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully purchased listing."), "Data": resp})

	log.Printf("Successful: PurchaseSingleItem")
}
