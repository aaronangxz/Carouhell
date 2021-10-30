package models

type Wallet struct {
	WalletID      *uint32 `json:"wallet_id" gorm:"primary_key"`
	WalletBalance *uint32 `json:"wallet_balance"`
	WalletStatus  *uint32 `json:"wallet_status"`
	LastTopUp     *int64  `json:"last_top_up"`
	LastUsed      *int64  `json:"last_used"`
}

func (r *Wallet) GetWalletBalance() uint32 {
	if r != nil && r.WalletBalance != nil {
		return *r.WalletBalance
	}
	return 0
}

type WalletTransaction struct {
	WtTransactionID   *uint32 `json:"wt_transaction_id"`
	WtWalletID        *uint32 `json:"wt_wallet_id" gorm:"primary_key"`
	TransactionCtime  *int64  `json:"transaction_ctime"`
	TransactionAmount *uint32 `json:"transaction_amount"`
	TransactionType   *uint32 `json:"transaction_type"`
	TransactionRef    *uint32 `json:"transaction_ref"`
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
	WalletBalance *uint32 `json:"wallet_balance"`
	WalletStatus  *uint32 `json:"wallet_status"`
	LastTopUp     *int64  `json:"last_top_up"`
	LastUsed      *int64  `json:"last_used"`
}

type TopUpUserWalletRequest struct {
	UserID *uint32 `json:"user_id" binding:"required"`
	Amount *uint32 `json:"amount" binding:"required"`
}

func (r *TopUpUserWalletRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

func (r *TopUpUserWalletRequest) GetAmount() uint32 {
	if r != nil && r.Amount != nil {
		return *r.Amount
	}
	return 0
}

type TopUpUserWalletResponse struct {
	WalletBalance uint32 `json:"wallet_balance"`
}
