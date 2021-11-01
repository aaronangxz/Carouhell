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
)

func InvalidateCache(service string, ID uint32) error {
	key := fmt.Sprintf("%v%v", service, ID)
	if err := models.RedisClient.Del(models.Ctx, key); err != nil {
		log.Printf("Failed to invalidate cache: %v", err)
		errormsg := fmt.Sprintf("Failed to invalidate cache: %v. %v", key, err)
		return errors.New(errormsg)
	}
	log.Printf("Successfully invalidated cache: %v", key)
	return nil
}
