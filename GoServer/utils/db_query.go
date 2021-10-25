package utils

const (
	ListingFixedQuery = "SELECT l.item_id, l.item_name, l.item_price, l.item_quantity," +
		" l.item_purchased_quantity, l.item_description, l.item_shipping_info," +
		" l.item_payment_info,l.item_location, l.item_status, l.item_category," +
		" l.item_image, l.seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.item_id) as listing_likes" +
		" FROM listing_tab l, acc_tab a" +
		" LEFT JOIN listing_reactions_tab ON l.item_id = listing_reactions_tab.item_id AND listing_reactions_tab.reaction_type = 0" +
		" WHERE l.seller_id = a.user_id" +
		" GROUP BY l.item_id"

	ListingQueryWithCustomCondition = "SELECT l.item_id, l.item_name, l.item_price, l.item_quantity," +
		" l.item_purchased_quantity, l.item_description, l.item_shipping_info," +
		" l.item_payment_info,l.item_location, l.item_status, l.item_category," +
		" l.item_image, l.seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, COUNT(listing_reactions_tab.item_id) as listing_likes" +
		" FROM listing_tab l, acc_tab a" +
		" LEFT JOIN listing_reactions_tab ON l.item_id = listing_reactions_tab.item_id AND listing_reactions_tab.reaction_type = 0" +
		" WHERE l.seller_id = a.user_id"
)

//Fixed query, not possible to append WHERE clause
//Without ORDER clause
func GetListingFixedQuery() string {
	return ListingFixedQuery
}

//Query without GROUP BY,ORDER BY; must append it after WHERE clauses
func GetListingQueryWithCustomCondition() string {
	return ListingFixedQuery
}
