package models

var (
	MaxNotificationTextLength        = uint32(256)
	MaxNotificationResponseSize      = uint32(50)
	DefaultNotificationResponseLimit = uint32(0)
)

func SetMaxNotificationTextLength() *uint32 {
	return &MaxNotificationTextLength
}

func SetDefaultNotificationResponseLimit() *uint32 {
	return &DefaultNotificationResponseLimit
}

func SetMaxNotificationResponseSize() *uint32 {
	return &MaxNotificationResponseSize
}

type Notification struct {
	NotificationID   *uint32 `json:"notification_id" gorm:"primary_key"`
	UserID           *uint32 `json:"user_id"`
	NotificationText *string `json:"notification_text"`
}

type GetNotificationsByUserIDRequest struct {
	UserID *uint32 `json:"user_id" binding:"required"`
	Limit  *uint32 `json:"limit"`
}

func (r GetNotificationsByUserIDRequest) GetUserID() uint32 {
	return *r.UserID
}

func (r GetNotificationsByUserIDRequest) GetLimit() uint32 {
	return *r.Limit
}

type GetNotificationsByUserIDResponse struct {
	NotificationID   uint   `json:"notification_id"`
	NotificationText string `json:"notification_text"`
}

type CreateNotificationRequest struct {
	UserID           *uint32 `json:"user_id" binding:"required"`
	NotificationText *string `json:"notification_text" binding:"required"`
}

func (r CreateNotificationRequest) GetUserID() uint32 {
	return *r.UserID
}

func (r CreateNotificationRequest) GetNotificationText() string {
	return *r.NotificationText
}

type ResponseMeta struct {
	DebugMsg  string
	ErrorCode int32
}
