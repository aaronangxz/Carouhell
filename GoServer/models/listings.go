package models

var (
	MaxStringLength         = uint32(256)
	MaxListingsResponseSize = uint32(50)
)

func SetMaxListingsResponseSize() *uint32 {
	return &MaxListingsResponseSize
}

type Listing struct {
	ItemID                uint32 `json:"item_id" gorm:"primary_key"`
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

func (r UpdateListingRequest) GetItemID() uint32 {
	return *r.ItemID
}

func (r UpdateListingRequest) GetItemName() string {
	return *r.ItemName
}

func (r UpdateListingRequest) GetItemPrice() uint32 {
	return *r.ItemPrice
}

func (r UpdateListingRequest) GetItemQuantity() uint32 {
	return *r.ItemQuantity
}

func (r UpdateListingRequest) GetItemDescription() string {
	return *r.ItemDescription
}

func (r UpdateListingRequest) GetShippingInfo() uint32 {
	return *r.ItemShippingInfo
}

func (r UpdateListingRequest) GetPaymentInfo() uint32 {
	return *r.ItemPaymentInfo
}

func (r UpdateListingRequest) GetItemLocation() string {
	return *r.ItemLocation
}

func (r UpdateListingRequest) GetItemCategory() uint32 {
	return *r.ItemCategory
}

func (r UpdateListingRequest) GetItemImage() string {
	return *r.ItemImage
}

func (r UpdateListingRequest) GetSellerID() uint32 {
	return *r.SellerID
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
	ItemPamentInfo        uint32
	ItemLocation          string
	ItemStatus            uint32
	ItemCategory          uint32
	ItemImage             string
	SellerID              uint32
	ListingCtime          uint32
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
