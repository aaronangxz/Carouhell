package models

var (
	MaxStringLength         = uint32(256)
	MaxListingsResponseSize = uint32(50)
)

func SetMaxListingsResponseSize() *uint32 {
	return &MaxListingsResponseSize
}

type Listing struct {
	ItemID                *uint32 `json:"item_id" gorm:"primary_key"`
	ItemName              *string `json:"item_name"`
	ItemPrice             *uint32 `json:"item_price"`
	ItemQuantity          *uint32 `json:"item_quantity"`
	ItemPurchasedQuantity *uint32 `json:"item_purchasedquantity"`
	ItemDescription       *string `json:"item_description"`
	ItemShippingInfo      *uint32 `json:"item_shippinginfo"`
	ItemPaymentInfo       *uint32 `json:"item_paymentinfo"`
	ItemLocation          *string `json:"item_location"`
	ItemStatus            *uint32 `json:"item_status"`
	ItemCategory          *uint32 `json:"item_category"`
	ItemImage             *string `json:"item_image"`
	SellerID              *uint32 `json:"seller_id"`
	ListingCtime          *int64  `json:"listing_ctime"`
	ListingMtime          *int64  `json:"listing_mtime"`
	ListingLikes          *uint32 `json:"listing_likes"`
}

func (r *Listing) GetItemID() uint32 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
	}
	return 0
}

func (r *Listing) GetSellerID() uint32 {
	if r != nil && r.SellerID != nil {
		return *r.SellerID
	}
	return 0
}

type GetAllListingsResponse struct {
	ItemID                uint32 `json:"item_id"`
	ItemName              string `json:"item_name"`
	ItemPrice             uint32 `json:"item_price"`
	ItemQuantity          uint32 `json:"item_quantity"`
	ItemPurchasedQuantity uint32 `json:"item_purchasedquantity"`
	ItemDescription       string `json:"item_description"`
	ItemShippingInfo      uint32 `json:"item_shippinginfo"`
	ItemPaymentInfo       uint32 `json:"item_paymentinfo"`
	ItemLocation          string `json:"item_location"`
	ItemStatus            uint32 `json:"item_status"`
	ItemCategory          uint32 `json:"item_category"`
	ItemImage             string `json:"item_image"`
	SellerID              uint32 `json:"seller_id"`
	ListingCtime          int64  `json:"listing_ctime"`
	ListingMtime          int64  `json:"listing_mtime"`
	ListingLikes          uint32 `json:"listing_likes"`
}

type CreateListingRequest struct {
	ItemName         *string `json:"item_name" binding:"required"`
	ItemPrice        *uint32 `json:"item_price" binding:"required"`
	ItemImage        *string `json:"item_image" binding:"required"`
	ItemQuantity     *uint32 `json:"item_quantity" binding:"required"`
	ItemDescription  *string `json:"item_description" binding:"required"`
	ItemShippingInfo *uint32 `json:"item_shipping_info" binding:"required"`
	ItemPaymentInfo  *uint32 `json:"item_payment_info" binding:"required"`
	ItemLocation     *string `json:"item_location" binding:"required"`
	ItemCategory     *uint32 `json:"item_category" binding:"required"`
	SellerID         *uint32 `json:"seller_id" binding:"required"`
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

func (r CreateListingRequest) GetShippingInfo() uint32 {
	return *r.ItemShippingInfo
}

func (r CreateListingRequest) GetPaymentInfo() uint32 {
	return *r.ItemPaymentInfo
}

func (r CreateListingRequest) GetItemLocation() string {
	return *r.ItemLocation
}

func (r CreateListingRequest) GetItemCategory() uint32 {
	return *r.ItemCategory
}

func (r CreateListingRequest) GetSellerID() uint32 {
	return *r.SellerID
}

type GetSingleListingRequest struct {
	ItemID *uint32 `json:"item_id" binding:"required"`
}

func (r GetSingleListingRequest) GetItemID() uint32 {
	return *r.ItemID
}

type DeleteSingleListingRequest struct {
	ItemID *uint32 `json:"item_id" binding:"required"`
}

func (r DeleteSingleListingRequest) GetItemID() uint32 {
	return *r.ItemID
}

type UpdateListingRequest struct {
	ItemID           *uint32 `json:"item_id"`
	ItemName         *string `json:"item_name"`
	ItemPrice        *uint32 `json:"item_price"`
	ItemQuantity     *uint32 `json:"item_quantity"`
	ItemDescription  *string `json:"item_description"`
	ItemShippingInfo *uint32 `json:"item_shippinginfo"`
	ItemPaymentInfo  *uint32 `json:"item_paymentinfo"`
	ItemLocation     *string `json:"item_location"`
	ItemCategory     *uint32 `json:"item_category"`
	ItemImage        *string `json:"item_image"`
	SellerID         *uint32 `json:"seller_id"`
}

func (r *UpdateListingRequest) GetItemID() uint32 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
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

func (r *UpdateListingRequest) GetShippingInfo() uint32 {
	if r != nil && r.ItemShippingInfo != nil {
		return *r.ItemShippingInfo
	}
	return 0
}

func (r *UpdateListingRequest) GetPaymentInfo() uint32 {
	if r != nil && r.ItemPaymentInfo != nil {
		return *r.ItemPaymentInfo
	}
	return 0
}

func (r *UpdateListingRequest) GetItemLocation() string {
	if r != nil && r.ItemName != nil {
		return *r.ItemLocation
	}
	return ""
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
	if r != nil && r.SellerID != nil {
		return *r.SellerID
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
	ItemID                uint32
	ItemName              string
	ItemPrice             uint32
	ItemQuantity          uint32
	ItemPurchasedQuantity uint32
	ItemDescription       string
	ItemShippingInfo      uint32
	ItemPaymentInfo       uint32
	ItemLocation          string
	ItemStatus            uint32
	ItemCategory          uint32
	ItemImage             string
	SellerID              uint32
	ListingCtime          uint32
	ListingMtime          int64
	ListingLikes          uint32
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

type CategoryFilter struct {
	ItemCategory *uint32 `json:"item_category"`
}

type LocationFilter struct {
	Location *string `json:"location"`
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

func (r LocationFilter) GetLocation() string {
	return *r.Location
}

func (r PriceFilter) GetMinPrice() uint32 {
	return *r.MinPrice
}

func (r PriceFilter) GetMaxPrice() uint32 {
	return *r.MaxPrice
}
