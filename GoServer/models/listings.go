// models/book.go

package models

type Listing struct {
	ItemID    uint   `json:"item_id" gorm:"primary_key"`
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
}

type CreateListingRequest struct {
	ItemName  string `json:"item_name" binding:"required"`
	ItemPrice int    `json:"item_price" binding:"required"`
}

type UpdateListingRequest struct {
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
}
