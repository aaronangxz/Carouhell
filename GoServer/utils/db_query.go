package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
)

const (
	ListingFixedQuery = "SELECT l.l_item_id, l.item_name, l.item_price, l.item_quantity," +
		" l.item_purchased_quantity, l.item_description, l.item_location, l.item_status, l.item_category," +
		" l.item_image, l.l_seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.rt_item_id) as listing_likes" +
		" FROM acc_tab a, listing_tab l" +
		" LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id AND listing_reactions_tab.reaction_type = 0" +
		" WHERE l.l_seller_id = a.a_user_id" +
		" GROUP BY l.l_item_id"

	ListingQueryWithCustomCondition = "SELECT l.l_item_id, l.item_name, l.item_price, l.item_quantity," +
		" l.item_purchased_quantity, l.item_description, l.item_location, l.item_status, l.item_category," +
		" l.item_image, l.l_seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.rt_item_id) as listing_likes" +
		" FROM acc_tab a, listing_tab l" +
		" LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id AND listing_reactions_tab.reaction_type = 0" +
		" WHERE l.l_seller_id = a.a_user_id"
)

//Fixed query, not possible to append WHERE clause
//Without ORDER clause
func GetListingFixedQuery() string {
	return ListingFixedQuery
}

//Query without GROUP BY,ORDER BY; must append it after WHERE clauses
func GetListingQueryWithCustomCondition() string {
	return ListingQueryWithCustomCondition
}

//Transaction to top up wallet and update corresponding tables
func StartWalletTopUpTx(input models.TopUpUserWalletRequest) (uint32, error) {
	var (
		resp models.TopUpUserWalletResponse
	)
	log.Println("beginning StartWalletTopUpTx")
	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("failed to recover: StartWalletTopUpTx")
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}
	updateWalletTransaction := fmt.Sprintf("INSERT INTO wallet_transactions_tab (wt_wallet_id,transaction_ctime,transaction_amount,transaction_type)"+
		"VALUES (%v,%v,%v,%v)", input.GetUserID(), time.Now().Unix(), input.GetAmount(), constant.TRANSACTION_TYPE_TOPUP)
	if err := tx.Exec(updateWalletTransaction).Error; err != nil {
		log.Printf("Error during StartWalletTopUpTx:updateWalletTransaction: %v", err.Error())
		log.Println("rolling back updateWalletTransaction")
		tx.Rollback()
		return 0, err
	}

	updateWalletBalance := fmt.Sprintf("UPDATE wallet_tab SET wallet_balance = wallet_balance + %v, last_top_up = %v WHERE wallet_id = %v",
		input.GetAmount(), time.Now().Unix(), input.GetUserID())
	if err := tx.Exec(updateWalletBalance).Error; err != nil {
		log.Printf("Error during StartWalletTopUpTx:updateWalletBalance: %v", err.Error())
		log.Println("rolling back updateWalletBalance")
		tx.Rollback()
		return 0, err
	}

	returnWalletBalance := fmt.Sprintf("SELECT wallet_balance FROM wallet_tab WHERE wallet_id = %v", input.GetUserID())
	if err := tx.Raw(returnWalletBalance).Scan(&resp).Error; err != nil {
		log.Printf("Error during StartWalletTopUpTx:returnWalletBalance: %v", err.Error())
		log.Println("rolling back returnWalletBalance")
		tx.Rollback()
		return 0, err
	}

	log.Println("committing StartWalletTopUpTx")
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}
	return resp.WalletBalance, nil
}
