package utils

import (
	"errors"
	"fmt"
	"log"

	"github.com/aaronangxz/TIC2601/models"
)

var (
	GetSingleListingByUserIDCacheKey = "get_single_listing_by_itemid:"
	GetUserWalletDetailsCacheKey     = "get_user_wallet_details:"
	GetUserLikedListingsCacheKey     = "get_user_liked_listings:"
	GetUserDetailsCacheKey           = "get_user_details:"
)

func InvalidateCache(service string, ID uint32) error {
	key := fmt.Sprintf("%v%v", service, ID)
	if _, err := models.RedisClient.Del(models.Ctx, key).Result(); err != nil {
		log.Printf("Failed to invalidate cache: %v", err)
		errormsg := fmt.Sprintf("Failed to invalidate cache: %v. %v", key, err)
		return errors.New(errormsg)
	}
	log.Printf("Successfully invalidated cache: %v", key)
	return nil
}

//Used to invalidate seller cache when we dont have direct access to userID
func InvalidateSellerCacheUsingItemID(service string, itemID uint32) error {
	var (
		hold models.Listing
	)

	//check sellerid
	checkSellerIDQuery := fmt.Sprintf("SELECT * FROM listing_tab WHERE l_item_id = %v", itemID)
	if err := models.DB.Raw(checkSellerIDQuery).Scan(&hold).Error; err != nil {
		errormsg := fmt.Sprintf("Error during invalidateSellerCacheUsingItemID query %v.", err.Error())
		return errors.New(errormsg)
	}

	//invalid cache
	if err := InvalidateCache(GetUserDetailsCacheKey, hold.GetLSellerID()); err != nil {
		log.Printf("Error during invalidateSellerCacheUsingItemID: %v", err.Error())
	}
	log.Println("Successfull: InvalidateSellerCacheUsingItemID")
	return nil
}
