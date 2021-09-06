// models/book.go

package models

type Listing struct {
	ItemID    uint   `json:"item_id" gorm:"primary_key"`
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
	ItemImg   string `json:"item_img"`
}

type CreateListingRequest struct {
	ItemName  string `json:"item_name" binding:"required"`
	ItemPrice int    `json:"item_price" binding:"required"`
	ItemImg   string `json:"item_img" binding:"required"`
}

type UpdateListingRequest struct {
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
	ItemImg   string `json:"item_img"`
}
