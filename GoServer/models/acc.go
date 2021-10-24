package models

type Account struct {
	UserID        *uint32 `json:"user_id" gorm:"primary_key"`
	UserName      *string `json:"user_name"`
	UserEmail     *string `json:"user_email"`
	UserCtime     *int64  `json:"user_ctime"`
	UserStatus    *uint32 `json:"user_status"`
	UserImage     *string `json:"user_image"`
	UserLastLogin *int64  `json:"user_last_login"`
	UserRating    *uint32 `json:"user_rating"`
}

func (r *Account) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type AccountCredentials struct {
	UserID               *uint32 `json:"user_id" gorm:"primary_key"`
	UserPassword         *string `json:"user_password"`
	UserSecurityQuestion *uint32 `json:"user_security_question"`
	UserSecurityAnswer   *string `json:"user_security_answer"`
}

func (r *AccountCredentials) GetUserPassword() string {
	if r != nil && r.UserPassword != nil {
		return *r.UserPassword
	}
	return ""
}

type CreateAccountRequest struct {
	UserName             *string `json:"user_name" binding:"required"`
	UserEmail            *string `json:"user_email" binding:"required"`
	UserPassword         *string `json:"user_password" binding:"required"`
	UserSecurityQuestion *uint32 `json:"user_security_question" binding:"required"`
	UserSecurityAnswer   *string `json:"user_security_answer" binding:"required"`
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

func (r *CreateAccountRequest) GetUserSecurityQuestion() uint32 {
	if r != nil && r.UserSecurityQuestion != nil {
		return *r.UserSecurityQuestion
	}
	return 0
}

func (r *CreateAccountRequest) GetUserSecurityAnswer() string {
	if r != nil && r.UserSecurityAnswer != nil {
		return *r.UserSecurityAnswer
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
