package models

type Notification struct {
	NotificationID   uint   `json:"notification_id" gorm:"primary_key"`
	UserID           uint   `json:"user_id"`
	NotificationText string `json:"notification_text"`
}

type GetNotificationsByUserIDRequest struct {
	UserID uint `json:"user_id"`
	Limit  int
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
