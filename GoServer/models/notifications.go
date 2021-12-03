package models

type Notification struct {
	UserName           *string `json:"user_name"`
	ItemName           *string `json:"item_name"`
	NotificationType   *uint32 `json:"notification_type"`
	NotificationString *string `json:"notification_string"`
	Ctime              *int64  `json:"ctime"`
}

type GetUserNotificationsRequest struct {
	UserID *uint32
}

func (r *GetUserNotificationsRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type GetUserNotificationsResponse struct {
	NotificationsList []Notification
}
