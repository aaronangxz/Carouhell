package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
)

var (
	ListingFixedQuery = fmt.Sprintf("SELECT l.l_item_id, l.item_name, l.item_price, l.item_quantity,"+
		" l.item_purchased_quantity, l.item_description, l.item_location, l.item_status, l.item_category,"+
		" l.item_image, l.l_seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.rt_item_id) as listing_likes"+
		" FROM acc_tab a, listing_tab l"+
		" LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id AND listing_reactions_tab.reaction_type = %v"+
		" WHERE l.l_seller_id = a.a_user_id AND l.item_status = %v"+
		" GROUP BY l.l_item_id", constant.LISTING_REACTION_TYPE_LIKE, constant.ITEM_STATUS_NORMAL)

	ListingQueryWithCustomCondition = fmt.Sprintf("SELECT l.l_item_id, l.item_name, l.item_price, l.item_quantity,"+
		" l.item_purchased_quantity, l.item_description, l.item_location, l.item_status, l.item_category,"+
		" l.item_image, l.l_seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.rt_item_id) as listing_likes"+
		" FROM acc_tab a, listing_tab l"+
		" LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id AND listing_reactions_tab.reaction_type = %v"+
		" WHERE l.l_seller_id = a.a_user_id", constant.LISTING_REACTION_TYPE_LIKE)

	WalletTransactionQuery = fmt.Sprintf("SELECT transaction_history.lt_item_id AS item_id,"+
		" transaction_history.transaction_amount, "+
		" transaction_history.transaction_type, "+
		" transaction_history.transaction_ctime, "+
		" item_info.item_name,item_info.item_image FROM("+
		" SELECT lt_item_id , transaction_amount, transaction_type, transaction_ctime FROM("+
		" SELECT lt_item_id, transaction_amount, 2 AS transaction_type, transaction_ctime FROM listing_transactions_tab "+
		" WHERE lt_item_id IN"+
		" (SELECT l_item_id FROM listing_tab "+
		" WHERE l_seller_id = ?)"+
		" UNION ALL"+
		" SELECT lt.lt_item_id, wt.transaction_amount, wt.transaction_type, wt.transaction_ctime FROM wallet_transactions_tab wt, listing_transactions_tab lt"+
		" WHERE wt.wt_wallet_id = ? AND wt.transaction_ref = lt.lt_transaction_id"+
		" UNION ALL"+
		" SELECT NULL AS lt_item_id, transaction_amount, transaction_type, transaction_ctime FROM wallet_transactions_tab"+
		" WHERE wt_wallet_id = ? AND transaction_type = %v) AS transactions) AS transaction_history"+
		" LEFT JOIN"+
		" (SELECT l_item_id, item_name, item_image FROM listing_tab) "+
		" AS item_info ON transaction_history.lt_item_id = item_info.l_item_id"+
		" ORDER BY transaction_ctime DESC", constant.TRANSACTION_TYPE_TOPUP)
)

//Fixed query, not possible to append WHERE clause
//Without ORDER clause
func GetListingFixedQuery() string {
	return ListingFixedQuery
}

func GetListingLoggedInQuery(user_id uint64) string {
	return fmt.Sprintf("SELECT l.l_item_id, l.item_name, l.item_price, l.item_quantity,"+
		" l.item_purchased_quantity, l.item_description, l.item_location, l.item_status, l.item_category,"+
		" l.item_image, l.l_seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.rt_item_id) as listing_likes, (CASE WHEN listing_reactions_tab.rt_user_id = %v THEN TRUE ELSE FALSE END) AS is_liked"+
		" FROM acc_tab a, listing_tab l"+
		" LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id AND listing_reactions_tab.reaction_type = %v"+
		" WHERE l.l_seller_id = a.a_user_id AND l.item_status = %v"+
		" GROUP BY l.l_item_id ORDER BY listing_ctime DESC", user_id, constant.LISTING_REACTION_TYPE_LIKE, constant.ITEM_STATUS_NORMAL)
}

//Query without GROUP BY,ORDER BY; must append it after WHERE clauses
func GetListingQueryWithCustomCondition() string {
	return ListingQueryWithCustomCondition
}

func GetListingLoggedInQueryWithCustomCondition(user_id uint64) string {
	return fmt.Sprintf("SELECT l.l_item_id, l.item_name, l.item_price, l.item_quantity,"+
		" l.item_purchased_quantity, l.item_description, l.item_location, l.item_status, l.item_category,"+
		" l.item_image, l.l_seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.rt_item_id) as listing_likes, (CASE WHEN listing_reactions_tab.rt_user_id = %v THEN TRUE ELSE FALSE END) AS is_liked"+
		" FROM acc_tab a, listing_tab l"+
		" LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id AND listing_reactions_tab.reaction_type = %v"+
		" WHERE l.l_seller_id = a.a_user_id", user_id, constant.LISTING_REACTION_TYPE_LIKE)
}

func GetWalletTransactionQuery() string {
	return WalletTransactionQuery
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

//Transaction to purchase item and update corresponding tables
func StartItemPurchaseTx(input models.PurchaseSingleItemRequest, totalPrice uint32) error {

	log.Println("beginning StartItemPurchaseTx")
	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("failed to recover: StartItemPurchaseTx")
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	//insert into listing_transaction
	listingTransaction := models.ListingTransaction{
		LtItemID:            input.ItemID,
		LtUserID:            input.UserID,
		TransactionCtime:    Int64(time.Now().Unix()),
		TransactionQuantity: input.PurchaseQuantity,
		TransactionAmount:   Uint32(totalPrice),
		TransactionStatus:   Uint32(constant.LISTING_TRANSACTION_STATUS_SUCCESS),
	}
	if err := tx.Table("listing_transactions_tab").Create(&listingTransaction).Error; err != nil {
		log.Printf("Error during StartItemPurchaseTx:insertListingTransaction: %v", err.Error())
		log.Println("rolling back insertListingTransaction")
		tx.Rollback()
		return err
	}

	walletTransaction := models.WalletTransaction{
		WtWalletID:        input.UserID,
		TransactionCtime:  Int64(time.Now().Unix()),
		TransactionAmount: Uint32(totalPrice),
		TransactionType:   Uint32(constant.TRANSACTION_TYPE_PURCHASE),
		TransactionRef:    listingTransaction.LtTransactionID,
	}
	if err := tx.Table("wallet_transactions_tab").Create(walletTransaction).Error; err != nil {
		log.Printf("Error during StartItemPurchaseTx:updateWalletTransaction: %v", err.Error())
		log.Println("rolling back updateWalletTransaction")
		tx.Rollback()
		return err
	}

	updateListingQuery := fmt.Sprintf("UPDATE listing_tab SET item_quantity = item_quantity - 1, item_status = CASE WHEN item_quantity = 0 THEN 2 ELSE item_status END WHERE l_item_id = %v", input.GetItemID())
	log.Println(updateListingQuery)
	if err := tx.Exec(updateListingQuery).Error; err != nil {
		log.Printf("Error during StartItemPurchaseTx:updateListing: %v", err.Error())
		log.Println("rolling back updateListing")
		tx.Rollback()
		return err
	}

	updateWalletQuery := fmt.Sprintf("UPDATE wallet_tab SET wallet_balance = wallet_balance - %v ,last_used = %v WHERE wallet_id = %v", totalPrice, time.Now().Unix(), input.GetUserID())
	log.Println(updateWalletQuery)
	if err := tx.Exec(updateWalletQuery).Error; err != nil {
		log.Printf("Error during StartItemPurchaseTx:updateWalletBalance: %v", err.Error())
		log.Println("rolling back updateWalletBalance")
		tx.Rollback()
		return err
	}

	log.Println("committing StartItemPurchaseTx")
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func GetFullTextSearchQuery(keyword string) string {
	return fmt.Sprintf("SELECT l.l_item_id, l.item_name, l.item_price, l.item_quantity,"+
		" l.item_purchased_quantity, l.item_description, l.item_location, l.item_status, l.item_category,"+
		" l.item_image, l.l_seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.rt_item_id) AS listing_likes,"+
		" ((1.5 * (MATCH(l.item_name) AGAINST ('%v*' IN BOOLEAN MODE))) + (0.5 * (MATCH(l.item_description) AGAINST ('%v*' IN BOOLEAN MODE)))) AS relevance"+
		" FROM acc_tab a, listing_tab l"+
		" LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id AND listing_reactions_tab.reaction_type = %v"+
		" WHERE MATCH(l.item_name,l.item_description) AGAINST ('%v*' IN BOOLEAN MODE)"+
		" AND l.l_seller_id = a.a_user_id  AND l.item_status = %v", keyword, keyword, constant.LISTING_REACTION_TYPE_LIKE, keyword, constant.ITEM_STATUS_NORMAL)
}

func GetFullTextSearchLoggedInQuery(keyword string, user_id uint64) string {
	return fmt.Sprintf("SELECT l.l_item_id, l.item_name, l.item_price, l.item_quantity,"+
		" l.item_purchased_quantity, l.item_description, l.item_location, l.item_status, l.item_category,"+
		" l.item_image, l.l_seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.rt_item_id) AS listing_likes,"+
		"(CASE WHEN listing_reactions_tab.rt_user_id = %v THEN TRUE ELSE FALSE END) AS is_liked,"+
		" ((1.5 * (MATCH(l.item_name) AGAINST ('%v*' IN BOOLEAN MODE))) + (0.5 * (MATCH(l.item_description) AGAINST ('%v*' IN BOOLEAN MODE)))) AS relevance"+
		" FROM acc_tab a, listing_tab l"+
		" LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id AND listing_reactions_tab.reaction_type = %v"+
		" WHERE MATCH(l.item_name,l.item_description) AGAINST ('%v*' IN BOOLEAN MODE)"+
		" AND l.l_seller_id = a.a_user_id  AND l.item_status = %v", user_id, keyword, keyword, constant.LISTING_REACTION_TYPE_LIKE, keyword, constant.ITEM_STATUS_NORMAL)
}
