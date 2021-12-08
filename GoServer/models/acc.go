package models

type Account struct {
	AUserID       *uint32 `json:"user_id" gorm:"primary_key"`
	UserName      *string `json:"user_name"`
	UserEmail     *string `json:"user_email"`
	UserCtime     *int64  `json:"user_ctime"`
	UserStatus    *uint32 `json:"user_status"`
	UserLastLogin *int64  `json:"user_last_login"`
	SellerType    *uint32 `json:"seller_type"`
}

func (r *Account) GetUserID() uint32 {
	if r != nil && r.AUserID != nil {
		return *r.AUserID
	}
	return 0
}

type AccountCredentials struct {
	CUserID      *uint32 `json:"user_id" gorm:"primary_key"`
	UserPassword *string `json:"user_password"`
}

func (r *AccountCredentials) GetUserPassword() string {
	if r != nil && r.UserPassword != nil {
		return *r.UserPassword
	}
	return ""
}

type UserReview struct {
	RVUserID   *uint32 `json:"user_id"`
	RVSellerID *uint32 `json:"seller_id"`
	Ratings    *uint32 `json:"ratings"`
	ReviewText *string `json:"review_text"`
	Ctime      *int64  `json:"ctime"`
}

type UserReviewWithNames struct {
	RVUserID   *uint32 `json:"user_id"`
	UserName   *string `json:"user_name"`
	RVSellerID *uint32 `json:"seller_id"`
	Ratings    *uint32 `json:"ratings"`
	ReviewText *string `json:"review_text"`
	Ctime      *int64  `json:"ctime"`
}
type CreateAccountRequest struct {
	UserName     *string `json:"user_name" binding:"required"`
	UserEmail    *string `json:"user_email" binding:"required"`
	UserPassword *string `json:"user_password" binding:"required"`
}

func (r *CreateAccountRequest) GetUserName() string {
	if r != nil && r.UserName != nil {
		return *r.UserName
	}
	return ""
}

func (r *CreateAccountRequest) GetUserEmail() string {
	if r != nil && r.UserEmail != nil {
		return *r.UserEmail
	}
	return ""
}

func (r *CreateAccountRequest) GetUserPassword() string {
	if r != nil && r.UserPassword != nil {
		return *r.UserPassword
	}
	return ""
}

type GetUserDetailsRequest struct {
	UserID *uint32 `json:"user_id" binding:"required"`
}

func (r *GetUserDetailsRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type UserRatings struct {
	Ratings float32 `json:"user_ratings"`
}
type GetUserDetailsResponse struct {
	AccountInfo  Account                   `json:"account_info"`
	Ratings      UserRatings               `json:"ratings"`
	ReviewCount  uint32                    `json:"review_count"`
	UserReviews  []UserReviewWithNames     `json:"user_reviews"`
	ListingCount uint32                    `json:"listing_count"`
	UserListings []GetUserListingsResponse `json:"user_listings"`
}

type AuthenticateUser struct {
	UserName     *string `json:"user_name" binding:"required"`
	UserPassword *string `json:"user_password" binding:"required"`
}

func (r *AuthenticateUser) GetUserName() string {
	if r != nil && r.UserName != nil {
		return *r.UserName
	}
	return ""
}

func (r *AuthenticateUser) GetUserPassword() string {
	if r != nil && r.UserPassword != nil {
		return *r.UserPassword
	}
	return ""
}

type AuthenticateUserResponse struct {
	UserID uint32 `json:"user_id"`
}
type AddUserReviewRequest struct {
	RVUserID   *uint32 `json:"user_id" binding:"required"`
	RVSellerID *uint32 `json:"seller_id" binding:"required"`
	Ratings    *uint32 `json:"ratings" binding:"required"`
	ReviewText *string `json:"review_text" binding:"required"`
}

func (r *AddUserReviewRequest) GetUserID() uint32 {
	if r != nil && r.RVUserID != nil {
		return *r.RVUserID
	}
	return 0
}

func (r *AddUserReviewRequest) GetSellerID() uint32 {
	if r != nil && r.RVSellerID != nil {
		return *r.RVSellerID
	}
	return 0
}

func (r *AddUserReviewRequest) GetRatings() uint32 {
	if r != nil && r.Ratings != nil {
		return *r.Ratings
	}
	return 0
}

func (r *AddUserReviewRequest) GetReviewText() string {
	if r != nil && r.ReviewText != nil {
		return *r.ReviewText
	}
	return ""
}

type AddUserReviewResponse struct {
	Ratings float32 `json:"latest_ratings"`
}

type CreateAccountSPResponse struct {
	Status *int
}

func (r *CreateAccountSPResponse) GetStatus() int {
	if r != nil && r.Status != nil {
		return *r.Status
	}
	return -1
}
