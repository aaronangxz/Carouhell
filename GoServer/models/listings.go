package models

const (
	MaxListingsResponseSize = 50
)

type Listing struct {
	ItemID    uint   `json:"item_id" gorm:"primary_key"`
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
	ItemImg   string `json:"item_img"`
}

type GetAllListingsResponse struct {
	ItemID    uint   `json:"item_id"`
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
	ItemImg   string `json:"item_img"`
}

type CreateListingRequest struct {
	ItemName  string `json:"item_name" binding:"required"`
	ItemPrice int    `json:"item_price" binding:"required"`
	ItemImg   string `json:"item_img" binding:"required"`
}

type GetSingleListingRequest struct {
	ItemID uint `json:"item_id" binding:"required"`
}

type DeleteSingleListingRequest struct {
	ItemID uint `json:"item_id" binding:"required"`
}

type UpdateListingRequest struct {
	ItemID    uint   `json:"item_id" binding:"required"`
	ItemName  string `json:"item_name" binding:"required"`
	ItemPrice int    `json:"item_price" binding:"required"`
	ItemImg   string `json:"item_img" binding:"required"`
}

type GetUserListingsRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	Limit  int
}

type GetPopularListingsRequest struct {
	ItemCategory string
	ItemStatus   int
	Limit        int
}

type GetLatestListingsRequest struct {
	ListingTime  int
	ItemCategory string
	ItemStatus   int
	Limit        int
}
