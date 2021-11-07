package models

var (
	MaxStringLength                = uint32(256)
	MaxItemNameStringLength        = uint32(25)
	MaxItemDescriptionStringLength = uint32(256)
	MaxListingsResponseSize        = uint32(50)
	DefaultImagePrefix             = "https://tic2601test.s3.ap-southeast-1.amazonaws.com/"
	DefaultImageExtension          = ".jpg"
)

func SetMaxListingsResponseSize() *uint32 {
	return &MaxListingsResponseSize
}

type Listing struct {
	LItemID               *uint32 `json:"item_id" gorm:"primary_key"`
	ItemName              *string `json:"item_name"`
	ItemPrice             *uint32 `json:"item_price"`
	ItemQuantity          *uint32 `json:"item_quantity"`
	ItemPurchasedQuantity *uint32 `json:"item_purchasedquantity"`
	ItemDescription       *string `json:"item_description"`
	ItemLocation          *uint32 `json:"item_location"`
	ItemStatus            *uint32 `json:"item_status"`
	ItemCategory          *uint32 `json:"item_category"`
	LSellerID             *uint32 `json:"seller_id"`
	ListingCtime          *int64  `json:"listing_ctime"`
	ListingMtime          *int64  `json:"listing_mtime"`
}

func (r *Listing) GetLItemID() uint32 {
	if r != nil && r.LItemID != nil {
		return *r.LItemID
	}
	return 0
}

func (r *Listing) GetLSellerID() uint32 {
	if r != nil && r.LSellerID != nil {
		return *r.LSellerID
	}
	return 0
}

func (r *Listing) GetItemQuantity() uint32 {
	if r != nil && r.ItemQuantity != nil {
		return *r.ItemQuantity
	}
	return 0
}

func (r *Listing) GetItemStatus() uint32 {
	if r != nil && r.ItemStatus != nil {
		return *r.ItemStatus
	}
	return 0
}

func (r *Listing) GetItemPrice() uint32 {
	if r != nil && r.ItemPrice != nil {
		return *r.ItemPrice
	}
	return 0
}

type GetAllListingsResponse struct {
	LItemID               uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemLocation          uint32 `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	LSellerID             uint32 `json:"seller_id"`
	SellerName            string `json:"seller_name"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
}

type CreateListingRequest struct {
	ItemName        *string `json:"item_name" binding:"required"`
	ItemPrice       *uint32 `json:"item_price" binding:"required"`
	ItemImage       *string `json:"item_image" binding:"required"`
	ItemQuantity    *uint32 `json:"item_quantity" binding:"required"`
	ItemDescription *string `json:"item_description" binding:"required"`
	ItemLocation    *uint32 `json:"item_location" binding:"required"`
	ItemCategory    *uint32 `json:"item_category" binding:"required"`
	SellerID        *uint32 `json:"seller_id" binding:"required"`
}

func (r CreateListingRequest) GetItemName() string {
	return *r.ItemName
}

func (r CreateListingRequest) GetItemPrice() uint32 {
	return *r.ItemPrice
}

func (r CreateListingRequest) GetItemImage() string {
	return *r.ItemImage
}

func (r CreateListingRequest) GetItemQuantity() uint32 {
	return *r.ItemQuantity
}

func (r CreateListingRequest) GetItemDescription() string {
	return *r.ItemDescription
}

func (r CreateListingRequest) SetItemDescription(s string) {
	*r.ItemDescription = s
}

func (r CreateListingRequest) GetItemLocation() uint32 {
	return *r.ItemLocation
}

func (r CreateListingRequest) GetItemCategory() uint32 {
	return *r.ItemCategory
}

func (r CreateListingRequest) GetSellerID() uint32 {
	return *r.SellerID
}

type CreateListingResponse struct {
	LItemID uint32 `json:"item_id"`
}

type GetSingleListingRequest struct {
	ItemID *uint32 `json:"item_id" binding:"required"`
}

func (r GetSingleListingRequest) GetItemID() uint32 {
	return *r.ItemID
}

type GetSingleListingResponse struct {
	LItemID               uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemLocation          uint32 `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	LSellerID             uint32 `json:"seller_id"`
	SellerName            string `json:"seller_name"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
	// IsLiked               bool                       `json:"is_liked"`
	CommentCount uint32                     `json:"listing_comment_count"`
	Comments     []ListingReactionsComments `json:"listing_comments"`
}

type GetSingleListingLoggedInResponse struct {
	LItemID               uint32                     `json:"item_id"`
	ItemName              string                     `json:"item_name"`
	ItemPrice             uint32                     `json:"item_price"`
	ItemQuantity          uint32                     `json:"item_quantity"`
	ItemPurchasedQuantity uint32                     `json:"item_purchasedquantity"`
	ItemDescription       string                     `json:"item_description"`
	ItemLocation          uint32                     `json:"item_location"`
	ItemStatus            uint32                     `json:"item_status"`
	ItemCategory          uint32                     `json:"item_category"`
	LSellerID             uint32                     `json:"seller_id"`
	SellerName            string                     `json:"seller_name"`
	ListingCtime          int64                      `json:"listing_ctime"`
	ListingMtime          int64                      `json:"listing_mtime"`
	ListingLikes          uint32                     `json:"listing_likes"`
	IsLiked               bool                       `json:"is_liked"`
	CommentCount          uint32                     `json:"listing_comment_count"`
	Comments              []ListingReactionsComments `json:"listing_comments"`
}

type DeleteSingleListingRequest struct {
	ItemID *uint32 `json:"item_id" binding:"required"`
}

func (r DeleteSingleListingRequest) GetItemID() uint32 {
	return *r.ItemID
}

type UpdateListingRequest struct {
	LItemID         *uint32 `json:"item_id"`
	ItemName        *string `json:"item_name"`
	ItemPrice       *uint32 `json:"item_price"`
	ItemQuantity    *uint32 `json:"item_quantity"`
	ItemDescription *string `json:"item_description"`
	ItemLocation    *uint32 `json:"item_location"`
	ItemCategory    *uint32 `json:"item_category"`
	ItemImage       *string `json:"item_image"`
	LSellerID       *uint32 `json:"seller_id"`
}

func (r *UpdateListingRequest) GetLItemID() uint32 {
	if r != nil && r.LItemID != nil {
		return *r.LItemID
	}
	return 0
}

func (r *UpdateListingRequest) GetItemName() string {
	if r != nil && r.ItemName != nil {
		return *r.ItemName
	}
	return ""
}

func (r *UpdateListingRequest) GetItemPrice() uint32 {
	if r != nil && r.ItemPrice != nil {
		return *r.ItemPrice
	}
	return 0
}

func (r *UpdateListingRequest) GetItemQuantity() uint32 {
	if r != nil && r.ItemQuantity != nil {
		return *r.ItemQuantity
	}
	return 0
}

func (r *UpdateListingRequest) GetItemDescription() string {
	if r != nil && r.ItemDescription != nil {
		return *r.ItemDescription
	}
	return ""
}

func (r *UpdateListingRequest) GetItemLocation() uint32 {
	if r != nil && r.ItemName != nil {
		return *r.ItemLocation
	}
	return 0
}

func (r *UpdateListingRequest) GetItemCategory() uint32 {
	if r != nil && r.ItemCategory != nil {
		return *r.ItemCategory
	}
	return 0
}

func (r *UpdateListingRequest) GetItemImage() string {
	if r != nil && r.ItemImage != nil {
		return *r.ItemImage
	}
	return ""
}

func (r *UpdateListingRequest) GetSellerID() uint32 {
	if r != nil && r.LSellerID != nil {
		return *r.LSellerID
	}
	return 0
}

type GetUserListingsRequest struct {
	UserID *uint32 `json:"user_id" binding:"required"`
	Limit  *uint32 `json:"limit"`
}

func (r GetUserListingsRequest) GetUserID() uint32 {
	return *r.UserID
}

func (r GetUserListingsRequest) GetLimit() uint32 {
	return *r.Limit
}

type GetUserListingsResponse struct {
	LItemID               uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemLocation          uint32 `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	LSellerID             uint32 `json:"seller_id"`
	SellerName            string `json:"seller_name"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
}

type GetPopularListingsRequest struct {
	ItemCategory *string
	ItemStatus   *uint32
	Limit        *uint32
}

type GetLatestListingsRequest struct {
	ItemCategory *uint32 `json:"item_category"`
	ItemStatus   *uint32 `json:"item_status"`
	Limit        *uint32 `json:"limit"`
}

type GetLatestListingsResponse struct {
	LItemID               uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemLocation          uint32 `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	LSellerID             uint32 `json:"seller_id"`
	SellerName            string `json:"seller_name"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
}

func (r GetLatestListingsRequest) GetItemCategory() uint32 {
	return *r.ItemCategory
}

func (r GetLatestListingsRequest) GetItemStatus() uint32 {
	return *r.ItemStatus
}

func (r GetLatestListingsRequest) GetLimit() uint32 {
	return *r.Limit
}

type GetLatestListingsLoggedInResponse struct {
	LItemID               uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemLocation          uint32 `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	LSellerID             uint32 `json:"seller_id"`
	SellerName            string `json:"seller_name"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
	IsLiked               bool   `json:"is_liked"`
}

type CategoryFilter struct {
	ItemCategory *uint32 `json:"item_category"`
}

type LocationFilter struct {
	Location *uint32 `json:"location"`
}

type PriceFilter struct {
	MinPrice *uint32 `json:"min_price"`
	MaxPrice *uint32 `json:"max_price"`
}

type GetListingsUsingFiltersRequest struct {
	SearchKeyword  *string `json:"search_keyword"`
	CategoryFilter `json:"category_filter"`
	LocationFilter `json:"location_filter"`
	PriceFilter    `json:"price_filter"`
	SortFlag       *uint32 `json:"sort_flag"`
}

func (r GetListingsUsingFiltersRequest) GetSearchKeyword() string {
	return *r.SearchKeyword
}

func (r GetListingsUsingFiltersRequest) GetSortFlag() uint32 {
	return *r.SortFlag
}

func (r GetListingsUsingFiltersRequest) SetSortFlag(flag uint32) {
	*r.SortFlag = flag
}

func (r CategoryFilter) GetItemCategory() uint32 {
	return *r.ItemCategory
}

func (r LocationFilter) GetLocation() uint32 {
	return *r.Location
}

func (r PriceFilter) GetMinPrice() uint32 {
	return *r.MinPrice
}

func (r PriceFilter) GetMaxPrice() uint32 {
	return *r.MaxPrice
}

type GetListingsUsingFiltersResponse struct {
	LItemID               uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemLocation          uint32 `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	LSellerID             uint32 `json:"seller_id"`
	SellerName            string `json:"seller_name"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
}

type GetListingsUsingFiltersLoggedInResponse struct {
	LItemID               uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemLocation          uint32 `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	LSellerID             uint32 `json:"seller_id"`
	SellerName            string `json:"seller_name"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
	IsLiked               bool   `json:"is_liked"`
}

type ListingReactions struct {
	ReactionsID  *uint32 `json:"id"`
	RTUserID     *uint32 `json:"user_id"`
	RTItemID     *uint32 `json:"item_id"`
	ReactionType *uint32 `json:"reaction_type"`
	Comment      *string `json:"comment"`
	Ctime        *uint32 `json:"ctime"`
}

type GetListingReactionsRequest struct {
	ItemID *uint32 `json:"item_id" binding:"required"`
}

func (r *GetListingReactionsRequest) GetItemID() uint32 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
	}
	return 0
}

type ListingReactionsComments struct {
	UserName string `json:"user_name"`
	Comment  string `json:"comment"`
	Ctime    int64  `json:"ctime"`
}
type GetListingReactionsResponse struct {
	LikesCount uint32                     `json:"likes_count"`
	Comments   []ListingReactionsComments `json:"comments"`
}

type AddListingLikeRequest struct {
	ItemID *uint32 `json:"item_id" binding:"required"`
	UserID *uint32 `json:"user_id" binding:"required"`
}

func (r *AddListingLikeRequest) GetItemID() uint32 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
	}
	return 0
}

func (r *AddListingLikeRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type AddListingLikeResponse struct {
	IsLiked    bool   `json:"is_liked"`
	LikesCount uint32 `json:"latest_likes_count"`
}

type GetUserLikedListingsRequest struct {
	UserID *uint32 `json:"user_id" binding:"required"`
}

func (r *GetUserLikedListingsRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type GetUserLikedListingsResponse struct {
	LItemID               uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemLocation          uint32 `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	LSellerID             uint32 `json:"seller_id"`
	SellerName            string `json:"seller_name"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
	IsLiked               bool   `json:"is_liked"`
}

type AddListingCommentsRequest struct {
	ItemID  *uint32 `json:"item_id" binding:"required"`
	UserID  *uint32 `json:"user_id" binding:"required"`
	Comment *string `json:"comment" binding:"required"`
}

func (r *AddListingCommentsRequest) GetItemID() uint32 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
	}
	return 0
}

func (r *AddListingCommentsRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

func (r *AddListingCommentsRequest) GetComment() string {
	if r != nil && r.Comment != nil {
		return *r.Comment
	}
	return ""
}

type PurchaseSingleItemRequest struct {
	ItemID           *uint32 `json:"item_id" binding:"required"`
	UserID           *uint32 `json:"user_id" binding:"required"`
	PurchaseQuantity *uint32 `json:"purchase_quantity" binding:"required"`
}

func (r *PurchaseSingleItemRequest) GetItemID() uint32 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
	}
	return 0
}

func (r *PurchaseSingleItemRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

func (r *PurchaseSingleItemRequest) GetPurchaseQuantity() uint32 {
	if r != nil && r.PurchaseQuantity != nil {
		return *r.PurchaseQuantity
	}
	return 0
}

type PurchaseSingleItemResponse struct {
	WalletBalance *uint32 `json:"wallet_balance"`
}

type ListingTransaction struct {
	LtTransactionID     *uint32 `json:"lt_transaction_id" gorm:"primary_key"`
	LtItemID            *uint32 `json:"lt_item_id"`
	LtUserID            *uint32 `json:"lt_user_id"`
	TransactionCtime    *int64  `json:"transaction_ctime"`
	TransactionQuantity *uint32 `json:"transaction_quantity"`
	TransactionAmount   *uint32 `json:"transaction_amount"`
}
