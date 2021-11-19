package models

type Notification struct {
	NotificationID    *uint32 `json:"notification_id" gorm:"primary_key"`
	NotificationText  *uint32 `json:"notification_text"`
	NotificationCTime *int64  `json:"notification_ctime"`
	NotificationType  *uint32 `json:"notification_type"`
	IsRead            *uint32
}

type GetUserNotificationsRequest struct {
	UserID *uint32
}

type GetUserNotificationsResponse struct {
	NotificationsList []Notification
}
