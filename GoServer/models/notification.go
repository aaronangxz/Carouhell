package models

var (
	MaxNotificationTextLength        = uint(256)
	MaxNotificationResponseSize      = uint(50)
	DefaultNotificationResponseLimit = uint(0)
)

func SetMaxNotificationTextLength() *uint {
	return &MaxNotificationTextLength
}

func SetDefaultNotificationResponseLimit() *uint {
	return &DefaultNotificationResponseLimit
}

func SetMaxNotificationResponseSize() *uint {
	return &MaxNotificationResponseSize
}

type Notification struct {
	NotificationID   *uint   `json:"notification_id" gorm:"primary_key"`
	UserID           *uint   `json:"user_id"`
	NotificationText *string `json:"notification_text"`
}

type GetNotificationsByUserIDRequest struct {
	UserID *uint `json:"user_id" binding:"required"`
	Limit  *uint `json:"limit"`
}

func (r GetNotificationsByUserIDRequest) GetUserID() uint {
	return *r.UserID
}

func (r GetNotificationsByUserIDRequest) GetLimit() uint {
	return *r.Limit
}

type GetNotificationsByUserIDResponse struct {
	NotificationID   uint   `json:"notification_id"`
	NotificationText string `json:"notification_text"`
}

type CreateNotificationRequest struct {
	UserID           *uint   `json:"user_id" binding:"required"`
	NotificationText *string `json:"notification_text" binding:"required"`
}

func (r CreateNotificationRequest) GetUserID() uint {
	return *r.UserID
}

func (r CreateNotificationRequest) GetNotificationText() string {
	return *r.NotificationText
}

type ResponseMeta struct {
	DebugMsg  string
	ErrorCode int
}
