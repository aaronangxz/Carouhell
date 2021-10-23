package models

type Wallet struct {
	WalletID      *uint32 `json:"wallet_id" gorm:"primary_key"`
	UserID        *uint32 `json:"user_id" gorm:"primary_key"`
	WalletBalance *uint32 `json:"wallet_balance"`
	WalletStatus  *uint32 `json:"wallet_status"`
	LastTopUp     *int64  `json:"last_top_up"`
	LastUsed      *int64  `json:"last_used"`
}

type WalletTransaction struct {
	TransactionID     *uint32 `json:"transaction_id"`
	WalletID          *uint32 `json:"wallet_id" gorm:"primary_key"`
	ItemID            *uint32 `json:"item_id"`
	PaymentType       *uint32 `json:"payment_type"`
	TransactionAmount *uint32 `json:"transaction_amount"`
	TransactionStatus *uint32 `json:"transaction_status"`
	TransactionCtime  *int64  `json:"transaction_ctime"`
	TransactionMtime  *int64  `json:"transaction_mtime"`
}

type CreateUserWalletRequest struct {
	UserID *uint32 `json:"user_id" binding:"required"`
}

func (r *CreateUserWalletRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type GetUserWalletDetailsRequest struct {
	UserID *uint32 `json:"user_id" binding:"required"`
}

func (r *GetUserWalletDetailsRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type GetUserWalletDetailsResponse struct {
	WalletID      *uint32 `json:"wallet_id"`
	UserID        *uint32 `json:"user_id"`
	WalletBalance *uint32 `json:"wallet_balance"`
	WalletStatus  *uint32 `json:"wallet_status"`
	LastTopUp     *int64  `json:"last_top_up"`
	LastUsed      *int64  `json:"last_used"`
}
