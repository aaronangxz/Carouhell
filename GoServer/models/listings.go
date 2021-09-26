package models

var (
	MaxStringLength         = uint(256)
	MaxListingsResponseSize = uint(50)
)

func SetMaxListingsResponseSize() *uint {
	return &MaxListingsResponseSize
}

type Listing struct {
	ItemID    *uint   `json:"item_id" gorm:"primary_key"`
	ItemName  *string `json:"item_name"`
	ItemPrice *uint   `json:"item_price"`
	ItemImg   *string `json:"item_img"`
}

type GetAllListingsResponse struct {
	ItemID    uint   `json:"item_id"`
	ItemName  string `json:"item_name"`
	ItemPrice uint   `json:"item_price"`
	ItemImg   string `json:"item_img"`
}

type CreateListingRequest struct {
	ItemName  *string `json:"item_name" binding:"required"`
	ItemPrice *uint   `json:"item_price" binding:"required"`
	ItemImg   *string `json:"item_img" binding:"required"`
}

func (r CreateListingRequest) GetItemName() string {
	return *r.ItemName
}

func (r CreateListingRequest) GetItemPrice() uint {
	return *r.ItemPrice
}

func (r CreateListingRequest) GetItemImg() string {
	return *r.ItemImg
}

type GetSingleListingRequest struct {
	ItemID *uint `json:"item_id" binding:"required"`
}

func (r GetSingleListingRequest) GetItemID() uint {
	return *r.ItemID
}

type DeleteSingleListingRequest struct {
	ItemID *uint `json:"item_id" binding:"required"`
}

func (r DeleteSingleListingRequest) GetItemID() uint {
	return *r.ItemID
}

type UpdateListingRequest struct {
	ItemID    *uint   `json:"item_id" binding:"required"`
	ItemName  *string `json:"item_name" binding:"required"`
	ItemPrice *uint   `json:"item_price" binding:"required"`
	ItemImg   *string `json:"item_img" binding:"required"`
}

func (r UpdateListingRequest) GetItemID() uint {
	return *r.ItemID
}

func (r UpdateListingRequest) GetItemName() string {
	return *r.ItemName
}

func (r UpdateListingRequest) GetItemPrice() uint {
	return *r.ItemPrice
}

func (r UpdateListingRequest) GetItemImg() string {
	return *r.ItemImg
}

type GetUserListingsRequest struct {
	UserID *uint `json:"user_id" binding:"required"`
	Limit  *uint `json:"limit"`
}

func (r GetUserListingsRequest) GetUserID() uint {
	return *r.UserID
}

func (r GetUserListingsRequest) GetLimit() uint {
	return *r.Limit
}

type GetPopularListingsRequest struct {
	ItemCategory *string
	ItemStatus   *uint
	Limit        *uint
}

type GetLatestListingsRequest struct {
	ListingTime  *uint
	ItemCategory *string
	ItemStatus   *uint
	Limit        *uint
}
