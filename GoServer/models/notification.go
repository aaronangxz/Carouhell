package models

type Notification struct {
	NotificationID   uint   `json:"notification_id" gorm:"primary_key"`
	UserID           uint   `json:"user_id"`
	NotificationText string `json:"notification_text"`
}

type GetNotificationsByUserIDResposne struct {
	NotificationID   uint   `json:"notification_id"`
	NotificationText string `json:"notification_text"`
}

type CreateNotificationRequest struct {
	UserID           uint   `json:"user_id" binding:"required"`
	NotificationText string `json:"notification_text" binding:"required"`
}

type ResponseMeta struct {
	DebugMsg  string
	ErrorCode int
}

const (
	CONST_NOTIFICATION_SUCCESS             = 0
	CONST_GET_NOTIFICATION_ERROR_NOT_FOUND = 1
	CONST_GET_NOTIFICATION_ERROR_UNKNOWN   = 2
)
