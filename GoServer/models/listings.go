package models

var (
	MaxStringLength         = uint32(256)
	MaxListingsResponseSize = uint32(50)
)

func SetMaxListingsResponseSize() *uint32 {
	return &MaxListingsResponseSize
}

type Listing struct {
	ItemID    *uint32 `json:"item_id" gorm:"primary_key"`
	ItemName  *string `json:"item_name"`
	ItemPrice *uint32 `json:"item_price"`
	ItemImg   *string `json:"item_img"`
}

type GetAllListingsResponse struct {
	ItemID    uint32 `json:"item_id"`
	ItemName  string `json:"item_name"`
	ItemPrice uint32 `json:"item_price"`
	ItemImg   string `json:"item_img"`
}

type CreateListingRequest struct {
	ItemName  *string `json:"item_name" binding:"required"`
	ItemPrice *uint32 `json:"item_price" binding:"required"`
	ItemImg   *string `json:"item_img" binding:"required"`
}

func (r CreateListingRequest) GetItemName() string {
	return *r.ItemName
}

func (r CreateListingRequest) GetItemPrice() uint32 {
	return *r.ItemPrice
}

func (r CreateListingRequest) GetItemImg() string {
	return *r.ItemImg
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
	ItemID    *uint32 `json:"item_id" binding:"required"`
	ItemName  *string `json:"item_name" binding:"required"`
	ItemPrice *uint32 `json:"item_price" binding:"required"`
	ItemImg   *string `json:"item_img" binding:"required"`
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

func (r UpdateListingRequest) GetItemImg() string {
	return *r.ItemImg
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
	ItemCategory *string
	ItemStatus   *uint32
	Limit        *uint32
}

func (r GetLatestListingsRequest) GetItemCategory() string {
	return *r.ItemCategory
}

func (r GetLatestListingsRequest) GetItemStatus() uint32 {
	return *r.ItemStatus
}
