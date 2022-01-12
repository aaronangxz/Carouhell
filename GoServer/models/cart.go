package models

type UserCart struct {
	ID           *uint32 `json:"id" gorm:"primary_key"`
	UserID       *int64  `json:"user_id"`
	ItemID       *int64  `json:"item_id"`
	ItemQuantity *uint32 `json:"item_quantity"`
	Ctime        *int64  `json:"ctime"`
}

func (r *UserCart) GetUserID() int64 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

func (r *UserCart) GetItemID() int64 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
	}
	return 0
}

func (r *UserCart) GetItemQuantity() uint32 {
	if r != nil && r.ItemQuantity != nil {
		return *r.ItemQuantity
	}
	return 0
}

func (r *UserCart) GetCtime() int64 {
	if r != nil && r.Ctime != nil {
		return *r.Ctime
	}
	return 0
}

type AddItemToUserCartRequest struct {
	UserID       *int64  `json:"user_id" binding:"required"`
	ItemID       *int64  `json:"item_id" binding:"required"`
	ItemQuantity *uint32 `json:"item_quantity" binding:"required"`
}

func (r *AddItemToUserCartRequest) GetUserID() int64 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

func (r *AddItemToUserCartRequest) GetItemID() int64 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
	}
	return 0
}

func (r *AddItemToUserCartRequest) GetItemQuantity() uint32 {
	if r != nil && r.ItemQuantity != nil {
		return *r.ItemQuantity
	}
	return 0
}

type GetUserCartRequest struct {
	UserID *int64 `json:"user_id" binding:"required"`
}

func (r *GetUserCartRequest) GetUserID() int64 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type UserCartForFrontEnd struct {
	UserID       *int64  `json:"user_id"`
	ItemID       *int64  `json:"item_id"`
	ItemQuantity *uint32 `json:"item_quantity"`
	Ctime        *int64  `json:"ctime"`
}

type GetUserCartResponse struct {
	ValidItems   []Listing
	InvalidItems []Listing
}
