package models

type Account struct {
	UserID        *uint32 `json:"user_id" gorm:"primary_key"`
	UserName      *string `json:"user_name"`
	UserEmail     *string `json:"user_email"`
	UserCtime     *uint32 `json:"user_ctime"`
	UserStatus    *uint32 `json:"user_status"`
	UserType      *uint32 `json:"user_type"`
	UserImage     *string `json:"user_image"`
	UserLastLogin *uint32 `json:"user_last_login"`
	UserRating    *uint32 `json:"user_rating"`
}
