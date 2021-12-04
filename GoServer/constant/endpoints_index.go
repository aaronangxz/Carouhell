package constant

import "github.com/aaronangxz/TIC2601/models"

func GetAPIDetails() map[string]models.API {
	m := make(map[string]models.API)

	m["CreateListing"] = models.API{APIName: "create_listing", APIMethod: "POST", APIDescription: "Create a new listing."}
	m["CreateAccount"] = models.API{APIName: "create_account", APIMethod: "POST", APIDescription: "Create a new account."}
	m["AuthenticateUser"] = models.API{APIName: "authenticate_user", APIMethod: "POST", APIDescription: "Login with UserName and Password."}
	m["GetAllListings"] = models.API{APIName: "get_all_listings", APIMethod: "GET", APIDescription: "Retrieve all listings."}
	m["GetLatestListings"] = models.API{APIName: "get_latest_listings", APIMethod: "GET", APIDescription: "Retrieve all available listings sorted by time in descending order."}
	m["GetListingsUsingFilters"] = models.API{APIName: "get_listings_using_filters", APIMethod: "POST", APIDescription: "Retrieve listings based on filtering conditions."}
	m["GetListingByItemID"] = models.API{APIName: "get_single_listing_by_itemid", APIMethod: "POST", APIDescription: "Retrieve all information of an individual listing."}
	m["UpdateSingleListing"] = models.API{APIName: "update_single_listing", APIMethod: "PATCH", APIDescription: "Update information of an existing listing."}
	m["DeleteListing"] = models.API{APIName: "delete_single_listing", APIMethod: "DELETE", APIDescription: "Delete an existing listing."}
	m["GetListingReactions"] = models.API{APIName: "get_listing_reactions", APIMethod: "POST", APIDescription: "Retrieve likes and comments on a listing."}
	m["AddListingLikes"] = models.API{APIName: "add_listing_likes", APIMethod: "POST", APIDescription: "Adding / removing a like on a listing."}
	m["AddListingComments"] = models.API{APIName: "add_listing_comments", APIMethod: "POST", APIDescription: "Adding a comment on a listing."}
	m["PurchaseSingleItem"] = models.API{APIName: "purchase_single_item", APIMethod: "POST", APIDescription: "Purchase an item."}
	m["GetUserListings"] = models.API{APIName: "get_user_listings", APIMethod: "POST", APIDescription: "Retrieve listings posted by individual users."}
	m["AddUserReview"] = models.API{APIName: "add_user_review", APIMethod: "POST", APIDescription: "Adding a review of a seller."}
	m["GetUserDetails"] = models.API{APIName: "get_user_details", APIMethod: "POST", APIDescription: "Retrieve all information, reviews, listings of an individual user."}
	m["GetUserLikedListings"] = models.API{APIName: "get_user_liked_listings", APIMethod: "POST", APIDescription: "Retrieve listings liked by users."}
	m["CreateUserWallet"] = models.API{APIName: "create_user_wallet", APIMethod: "POST", APIDescription: "Create and activate user's wallet."}
	m["GetUserWalletDetails"] = models.API{APIName: "get_user_wallet_details", APIMethod: "POST", APIDescription: "Retrieve wallet balance and all transactions."}
	m["TopUpUserWallet"] = models.API{APIName: "top_up_user_wallet", APIMethod: "POST", APIDescription: "Top up the balance of user's wallet."}
	m["GetUserNotifications"] = models.API{APIName: "get_user_notifications", APIMethod: "POST", APIDescription: "Retrieve likes / comments / purchases notifications."}

	return m
}
